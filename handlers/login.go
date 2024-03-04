package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Function to manage the login internally
func Login(w http.ResponseWriter, r *http.Request) error {

	log.Default().Printf("%v - Login started...", r.RemoteAddr)

	// Generate an auth code to set the cookie
	var authCode string = secmanagers.AuthCookieGen()

	// Create the authtoken cookie
	// Set expiration time in 30 minutes and writes it to responsewriter
	var cookie http.Cookie = secmanagers.CreateSecCk("authToken", authCode)
	cookie.Expires = time.Now().Add(30 * time.Minute)
	http.SetCookie(w, &cookie)

	// Writes the same auth cookie to db with a ttl of 30 mins
	err := managers.RegisterToken(authCode, time.Now().Add(30*time.Minute))
	if err != nil {
		return fmt.Errorf("Login function: %v", err)
	}
	return nil
}
