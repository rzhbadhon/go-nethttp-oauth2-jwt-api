package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var jwtSecretKey []byte

// this is payload you can call hehe
type AppClaims struct {
	UserID               uuid.UUID `json:"user_id"`
	Role                 string    `json:"role"`
	jwt.RegisteredClaims           // like expiry
}

func GenerateJWT(userID uuid.UUID, role string) (string, error) {
	// setting token expiry time 24 hour only
	expirationTime := time.Now().Add(24 * time.Hour)

	// setting claims
	claims := &AppClaims{ // this is payload actually
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// cooking header + payload with (hs256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// adding lil bit signature with secret key claims
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	// finsl token cooking done it must taste sooooo goood for sake
	return tokenString, nil
}

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, checking system env")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET env not set")
	}
	jwtSecretKey = []byte(secret)
}
