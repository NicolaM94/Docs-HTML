package handlers

import (
	"docshelf/secmanagers"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// TODO: Debug
// Responds to /
func CodeValidationHandler(w http.ResponseWriter, r *http.Request) {

	// Checks if inserted code is the same as the given one
	// Retrieves form value and cookie "code"
	var inputCode string = r.FormValue("inputcode")
	ck, err := r.Cookie("code")
	if err != nil {
		log.Fatal(err)
	}
	// Decodes the secure cookie instance to get its value
	var cookieCode string = secmanagers.DecodeSecCk(*ck)
	// Checks for equality
	if !(inputCode == cookieCode) {
		// If the codes do not match logs it to the server and redirects to index
		// TODO: The log here is intended for debugging purposes, to be removed later.
		log.Default().Printf("InputCode: %v   Cookie Code: %v\n", inputCode, cookieCode)
		log.Default().Println(r.RemoteAddr, " has got the code wrong. Redirecting to home page...")
		http.Redirect(w, r, "/", http.StatusAccepted)
		return
	}

	// Check if user reqeusted a login or a registration.
	// If cookie is not present or different from login and register.
	ck, err = r.Cookie("reqtype")
	if err != nil {
		fmt.Println("Error here")
		log.Fatal(err)
	}

	// Based on the value of the cookie regtype switches between the login or registration
	switch ck.Value {

	case "login":
		err = Login(w, r)
		if err != nil {
			log.Fatal(err)
		}
		// If no error is raised by Login function, reroute to /datadelivery
		http.Redirect(w, r, "/datadelivery", http.StatusFound)

	case "register":
		err = Register(w, r)
		if err != nil {
			log.Fatal("Register function : ", err)
		}
		// If no error is raised by Register function, proceed to http parsing.
		t, _ := template.ParseFiles("./static/registration-confirm.html")
		t.Execute(w, nil)

	// If the cookie does not hold "login" or "register" as value that means something went wrong while assigning its value.
	// Breaks the verification, logs the error and redirects to index.
	default:
		log.Fatal(errors.New(">> error: malformed reqtype cookie. Redirecting request to index"))
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
