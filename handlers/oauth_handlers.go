package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"go-auth-manual/auth"
	"net/http"
	"time"
)

// this func rediects user to the google login page
func (h *Handler) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {

	// 1. create random state token to protect from csrf attack
	//its a random string we will save it on cookie and  send this to google
	state := generateStateOauthCookie(w)

	// 2. redirect user to the google Authcode url
	// already its declared in auth pkg
	url := auth.GoogleOAuthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

// this creates a random state and set it on the cookie
func generateStateOauthCookie(w http.ResponseWriter) string {
	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	// set cookie with time limit of 10min
	cookie := http.Cookie{
		Name:     "oauthState",
		Value:    state,
		Expires:  time.Now().Add(10 * time.Minute),
		HttpOnly: true, // cant access cookie from javascript (no xss attack)
		Path:     "/",  // cookie will availble whole site
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie) // setting cookie to the response header

	return state

}
