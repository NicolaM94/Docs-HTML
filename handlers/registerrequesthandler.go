package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"html/template"
	"log"
	"net/http"
)

func RegReqHandler(w http.ResponseWriter, r *http.Request) {

	// Verifying the mail equality
	mailOne := r.FormValue("emailfield")
	mailTwo := r.FormValue("emailagainfield")
	if !secmanagers.EqualString(mailOne, mailTwo) {
		log.Default().Println("Mails are different. Should print error to frontend")
		http.Redirect(w, r, "/", http.StatusAccepted)
	}

	// Verifying the password equality
	passOne := r.FormValue("passfield")
	passTwo := r.FormValue("passagainfield")
	if !secmanagers.EqualString(passOne, passTwo) {
		log.Default().Println("Passwords are different. Should print error to frontend")
		http.Redirect(w, r, "/", http.StatusAccepted)
	}

	// Setting cookies to pass to the check-code section
	nameCookie := secmanagers.CreateSecCk("name", r.FormValue("name"))
	http.SetCookie(w, &nameCookie)
	surnameCookie := secmanagers.CreateSecCk("surname", r.FormValue("surname"))
	http.SetCookie(w, &surnameCookie)
	email := secmanagers.CreateSecCk("email", mailOne)
	http.SetCookie(w, &email)
	password := secmanagers.CreateSecCk("password", secmanagers.HashNSault(passOne))
	http.SetCookie(w, &password)

	// Sending mail with code and setting cookie to responsewriter
	code := managers.Codegen()
	codecookie := secmanagers.CreateSecCk("Code", code)
	http.SetCookie(w, &codecookie)
	err := managers.SendCodeMail(mailOne, code)
	if err != nil {
		log.Fatalln(err)
	}

	// Passing to code-validation
	t, _ := template.ParseFiles("./static/confirmcode.html")
	t.Execute(w, nil)

}
