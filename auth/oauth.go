package auth

import (
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// declaring variable globally to share
var GoogleOAuthConfig *oauth2.Config

func InitOAuthConfig() {

	// init caling this func to load env file
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("FATAL: GOOGLE CLIENT ID or GOOGLE CLIENT SECRET is not set")
	}

	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		// redirectUrl, redirets google user to the go backend
		RedirectURL: "http://localhost:9000/auth/google/callback",
		// scopes is what data of user we want from google
		Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",    
		"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	log.Println("OAuth Config loaded successfully")

}
