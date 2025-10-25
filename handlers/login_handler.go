package handlers

import (
	"database/sql"
	"encoding/json"
	"go-auth-manual/auth"
	"go-auth-manual/models"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// decode from json body that comes in request body
	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// find that req data into db
	var user models.User // this will help us to get user hashed pass
	query := `SELECT FROM users WHERE email=$1`

	/* db get search email across database and if found it puts all the data on the user struct */
	err = h.DB.Get(&user, query, req.Email)
	if err != nil{
		if err == sql.ErrNoRows{
			log.Println("User not found: ", err)
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		// this error means its a db internal error
		log.Println("Database query error: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// compare password with the hash pass stored in db
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil{
		log.Println("Password mismatch error")
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)// 401
		return
	}

	// calling jwt generator to get a access token
	tokenString, err := auth.GenerateJWT(user.ID, user.Role)
	if err != nil{
		log.Println("Error generating JWT: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// login successful yeeeyyy
	log.Println("User logged in successfully", user.Email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//  sending response with access token to frontend via map
	json.NewEncoder(w).Encode(map[string]string{
		"token" : tokenString,
	})
}