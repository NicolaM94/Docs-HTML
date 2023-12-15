package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
	"text/template"
)

// This function only verifies the code and moves on to the new password fillout process.
// If the validation fails, it returns to login page.
func PassCodeValidationHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieve insertedCode from form and decodes code cookie
	var insertedCode string = r.FormValue("inserted-code")
	cookieCode, err := utilities.DecodeSecureCookie("code", r)
	if err != nil {
		fmt.Println("Cannot decode code cookie")
	}

	// Checks if the two codes are equal. If not, redirect to login page
	if insertedCode != cookieCode["code"] {
		http.Redirect(w, r, "./static/login.html", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("./static/newpwsubmit.html")
	t.Execute(w, nil)
}
