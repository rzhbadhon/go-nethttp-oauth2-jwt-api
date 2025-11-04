package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-auth-manual/auth"
	"go-auth-manual/models"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type GoogleUserInfo struct {
	Email     string `json:"email"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
}

func (h *Handler) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {

	// 1. security check: the state verfication (for csrf attack)
	oauthstate, err := r.Cookie("oauthState")
	if err != nil {
		log.Println("Failed to get oauthstate cookie:", err)
		http.Error(w, "State cookie not found or expired", http.StatusUnauthorized)
		return
	}


	if r.FormValue("state") != oauthstate.Value {
		log.Println("Invalid oauth google state")
		http.Error(w, "Invalid state", http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "oauthstate",
		Value:    "",
		Expires:  time.Unix(0, 0), // set to expire cookie 
		HttpOnly: true,
		Path:     "/",
	})

	// 2. access the code  sent by google
	code := r.FormValue("code")
	if code == "" {
		log.Println("Code not found")
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	// 3. exchange the code into google token
	token, err := auth.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Code exchange failed", err)
		http.Error(w, "Code exchange failed", http.StatusInternalServerError)
		return
	}

	client := auth.GoogleOAuthConfig.Client(context.Background(), token)

	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Println("Failed getting user info:", err)
		http.Error(w, "Failed getting user info", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// reading user info content that google sent into response body
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed reading user info response: ", err)
		http.Error(w, "Failed reading use info response", http.StatusInternalServerError)
		return
	}

	var userInfo GoogleUserInfo
	json.Unmarshal(contents, &userInfo)

	// 5. database logic: finding user or create new
	var user models.User
	query := `SELECT * FROM users WHERE email=$1`
	err = h.DB.Get(&user, query, userInfo.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			// no user so create new
			log.Println("New user detected via google: ", user.Email)

			// create rand pass as we dont ned them
			randomPass := uuid.New().String()
			hashedPass, _ := bcrypt.GenerateFromPassword([]byte(randomPass), bcrypt.DefaultCost)

			user = models.User{
				ID:        uuid.New(),
				FirstName: userInfo.FirstName,
				LastName:  userInfo.LastName,
				Email:     userInfo.Email,
				Password:  string(hashedPass),
				Role:      "user",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			// insert new user to the db
			insertQuery := `INSERT INTO users (id, first_name, last_name, email, password, role, created_at, updated_at) VALUES (:id, :first_name, :last_name, :email, :password, :role, :created_at, :updated_at)`

			_, err = h.DB.NamedExec(insertQuery, &user)
			if err != nil {
				log.Println("Failed to create new OAuth user:", err)
				http.Error(w, "Failed to register user", http.StatusInternalServerError)
				return
			} else {
				log.Println("Database error:", err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}

		} else {
			log.Println("Existing user logged in via google:", user.Email)
		}

	}

	// create our own jwt... same as login func
	tokenString, err := auth.GenerateJWT(user.ID, user.Role)
	if err != nil {
		log.Println("Error generating our JWT:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// sending token to user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful via Google",
		"token":   tokenString,
	})

}
