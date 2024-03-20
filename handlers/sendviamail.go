package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"net/http"
)

func SendDocumentMailHandler(w http.ResponseWriter, r *http.Request) {

	// First check if auth token is present and not depleeted
	ck, err := r.Cookie("authToken")
	if err != nil {
		http.Redirect(w, r, "/loggedout.html", http.StatusFound)
	}
	authTokenCookie := secmanagers.DecodeSecCk(*ck)
	authtoken, err := managers.IsTokenPresent(authTokenCookie)
	if err != nil || !authtoken {
		http.Redirect(w, r, "./static/loggedout.html", http.StatusFound)
	}

	// Collect

}
