package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"log"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) error {
	log.Default().Printf("%v - Login started...", r.RemoteAddr)
	var authCode string = secmanagers.AuthCookieGen()
	var cookie http.Cookie = secmanagers.CreateSecCk("authToken", authCode)
	cookie.Expires = time.Now().Add(30 * time.Minute)
	http.SetCookie(w, &cookie)

	// Writes auth cookie to db with a ttl of 30 mins
	err := managers.RegisterToken(authCode, time.Now().Add(30*time.Minute))
	if err != nil {
		log.Fatal("Error while loggin in: ", err)
	}
	http.Redirect(w, r, "/datadelivery", http.StatusFound)
	return nil
}
