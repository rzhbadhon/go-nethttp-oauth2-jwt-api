package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
//  init func calls first before the main func
//  init is using to load env first before the main func
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
	log.Println("JWT Secret loaded successfully")

	InitOAuthConfig()
}
