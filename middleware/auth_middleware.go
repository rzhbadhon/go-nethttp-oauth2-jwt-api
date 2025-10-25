package middleware

import (
	"context"
	"go-auth-manual/auth"
	"log"
	"net/http"
)

type contextKey string

const userClaimsKey contextKey = "userClaims"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc{

	return func (w http.ResponseWriter, r *http.Request)  {
		// step 1: extract token from the header
		tokenString, err := auth.ExtractTokenFromHeader(r)
		if err != nil{
			log.Println("Auth error: ", err)
			http.Error(w, "Unauthorized"+ err.Error(), http.StatusUnauthorized)
			return 
		}

		// step 2: validate the freakin token
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil{
			log.Println("Auth error: Invalid token", err)
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// step 3: pass the data to the next handler with context

		ctx := context.WithValue(r.Context(), userClaimsKey, claims)

		// step 4: call the next handler... I mean middleware
		log.Printf("Autheticated user %s (Role: %s) accessing %s", claims.UserID, claims.Role, r.URL.Path)

		next.ServeHTTP(w, r. WithContext(ctx))


	}
}
