package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString string) (*AppClaims, error) {
	claims := &AppClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error){
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok{
		return nil, errors.New("unexpected signing method")
	}
	return jwtSecretKey, nil
	})

	if err != nil{
		log.Println("Error parsing token", err)
		return  nil, err
	}

	if !token.Valid{
		return  nil, errors.New("invalid token")
	}

	return claims, nil


}


// extract header brearer token
func ExtractTokenFromHeader(r *http.Request) (string, error){
	authHeader := r.Header.Get("Authorization")
	if authHeader == ""{
		return "", errors.New("authorization header required")
	}

	// token must be in "Bearer <token>" format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer"{
		return "", errors.New("invalid auth header format")
	}
	return parts[1], nil
}