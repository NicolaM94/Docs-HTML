package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
	"text/template"
)

func CodeValidatorHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieve insertedCode from form and decodes code cookie
	var insertedCode string = r.FormValue("inserted-code")
	cookieCode, err := utilities.DecodeSecureCookie("code", r)
	if err != nil {
		fmt.Println("Cannot decode code cookie")
	}

	// Checks if the two codes are equal. If not, redirect to register page
	if insertedCode != cookieCode["code"] {
		http.Redirect(w, r, "/register.html", http.StatusFound)
		return
	}

	// Checks request type, either register or login
	reqType, err := utilities.DecodeSecureCookie("type", r)
	if err != nil {
		fmt.Println("Cannot decode reqType cookie")
	}

	if reqType["type"] == "registration" {
		goto REGISTRATION
	} else {
		goto LOGIN
	}

REGISTRATION:
	{
		cookieName, err := utilities.DecodeSecureCookie("name", r)
		if err != nil {
			panic(err)
		}
		cookieSurname, err := utilities.DecodeSecureCookie("surname", r)
		if err != nil {
			panic(err)
		}
		cookieFiscalCode, err := utilities.DecodeSecureCookie("fiscalcode", r)
		if err != nil {
			panic(err)
		}
		cookieMail, err := utilities.DecodeSecureCookie("email", r)
		if err != nil {
			panic(err)
		}
		cookiePassword, err := utilities.DecodeSecureCookie("password", r)
		if err != nil {
			panic(err)
		}

		// Checks if user is already registered
		rows, err := utilities.QueryRow("select * from users")
		if err != nil {
			panic(err)
		}
		if utilities.SearchInRows(cookieMail["email"], rows) {
			t, _ := template.ParseFiles("./static/already-registered.html")
			fmt.Println("Already registered")
			t.Execute(w, nil)
		}

		err = utilities.InsertRow(cookieName["name"], cookieSurname["surname"], cookieFiscalCode["fiscalcode"], cookieMail["email"], cookiePassword["password"])
		if err != nil {
			panic(err)
		}
		t, _ := template.ParseFiles("./static/registration-confirm.html")
		t.Execute(w, nil)
	}

LOGIN:
	{

	}

	// if so, register the new user and parses registration confirm.

}
