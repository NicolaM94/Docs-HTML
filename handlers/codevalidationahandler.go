package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"html/template"
	"log"
	"net/http"
	"os"
)

// TODO: Debug
// Responds to /
func CodeValidationHandler(w http.ResponseWriter, r *http.Request) {

	// Verifies if inserted code is the same as the given one
	var inputCode string = r.FormValue("inputcode")
	ck, err := r.Cookie("Code")
	if err != nil {
		log.Fatal(err)
	}
	var cookieCode string = secmanagers.DecodeSecCk(*ck)
	if !(inputCode == cookieCode) {
		log.Default().Printf("InputCode: %v   Cookie Code: %v\n", inputCode, cookieCode)
		log.Default().Println(r.RemoteAddr, " has got the code wrong. Redirecting to home page...")
		http.Redirect(w, r, "/", http.StatusAccepted)
		return
	}

	// If codes are equal...
	// ... tries to create a folder with the name and surname
	ck, err = r.Cookie("name")
	if err != nil {
		log.Fatal(err)
	}
	name := secmanagers.DecodeSecCk(*ck)
	ck, err = r.Cookie("surname")
	if err != nil {
		log.Fatal(err)
	}
	surname := secmanagers.DecodeSecCk(*ck)
	var foldername string = name + "_" + surname

	var settings managers.Settings = managers.Settings{}.Populate()
	if settings.DocBasePath[len(settings.DocBasePath)-1] == '/' {
		foldername = settings.DocBasePath + foldername + "/"
	} else {
		foldername = settings.DocBasePath + "/" + foldername + "/"
	}
	err = os.Mkdir(foldername, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	ck, err = r.Cookie("email")
	if err != nil {
		log.Fatal(err)
	}
	email := secmanagers.DecodeSecCk(*ck)

	ck, err = r.Cookie("password")
	if err != nil {
		log.Fatal(err)
	}
	password := secmanagers.DecodeSecCk(*ck)

	// ... tries to register the user in the udb
	err = managers.RegisterUserUDB(email, password, name, surname)
	if err != nil {
		log.Fatal(err)
	}

	// ...route to confirm registration
	t, _ := template.ParseFiles("./static/registration-confirm.html")
	t.Execute(w, nil)
}
