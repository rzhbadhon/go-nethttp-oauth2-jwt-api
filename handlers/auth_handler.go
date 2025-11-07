package handlers

import (
	"encoding/json"
	"go-auth-manual/auth"
	"go-auth-manual/models"
	"log"
	"net/http"
	"go-auth-manual/middleware"
)

// this only for admin.. so middleware applied here
func (h *Handler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {

	// read context what middleware inserted in the request
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*auth.AppClaims)
	if !ok {
		http.Error(w, "Could not retrieve user claims", http.StatusInternalServerError)
		return
	}

	// only admin can access
	if claims.Role != "admin" {
		log.Printf("Forbidden: User %s (Role: %s) tried to access /users", claims.UserID, claims.Role)
		http.Error(w, "Forbidden: Only admins can view all users", http.StatusForbidden) // 403 Forbidden
		return
	}

	var users []models.User
	// we aint selecting the passwords of the users
	query := `SELECT id, first_name, last_name, email, role, created_at FROM users`
	err := h.DB.Select(&users, query)
	if err != nil {
		log.Println("Database error getting users:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// send response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
