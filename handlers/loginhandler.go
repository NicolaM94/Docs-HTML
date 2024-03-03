package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Catches and responds to login.html
func LoginReqHanlder(w http.ResponseWriter, r *http.Request) {

	// Retrieves forms for authentication
	mailinput := r.FormValue("mailinput")
	passinput := r.FormValue("passwordinput")

	// Retrieves the user requesting the login from the database
	usr, err := managers.QueryByMail(mailinput)
	if err != nil {
		log.Default().Println("Cannot find the requested user") //TODO: Should print error to front end
		http.Redirect(w, r, "/", http.StatusFound)
	}

	// Checks if the password is correct
	if secmanagers.HashNSault(passinput) != usr[0].Password {
		log.Default().Println("Password mismatch.") // TODO: Should print error to frontend
		http.Redirect(w, r, "/", http.StatusFound)
	}

	// If the password match, send authcode
	// Set code as cookie for further verification
	var code string = managers.Codegen()
	err = managers.SendCodeMail(mailinput, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating secure code cookie")
	var ck http.Cookie = secmanagers.CreateSecCk("code", code)
	http.SetCookie(w, &ck)
	// Set request type as login
	fmt.Println("Create cookie with req type")
	ck = http.Cookie{Name: "reqtype", Value: "login"}
	http.SetCookie(w, &ck)

	// Parse codevalidation template
	fmt.Println("Parsing validation code")
	t, _ := template.ParseFiles("./static/confirmcode.html")
	t.Execute(w, nil)
}
