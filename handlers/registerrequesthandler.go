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
	log.Default().Printf(">> %v - Starting registration process\n", r.RemoteAddr)
	log.Default().Printf(">> %v - Verifying mails differences...\n", r.RemoteAddr)
	mailOne := r.FormValue("emailfield")
	mailTwo := r.FormValue("emailagainfield")
	if !secmanagers.EqualString(mailOne, mailTwo) {
		log.Default().Printf(">>> %v - Mails are different\n", r.RemoteAddr) //TODO :Should print error to frontend
		http.Redirect(w, r, "/", http.StatusAccepted)
	}

	// Verifying the password equality
	log.Default().Printf(">> %v - Verifying passwords differences...\n", r.RemoteAddr)
	passOne := r.FormValue("passfield")
	passTwo := r.FormValue("passagainfield")
	if !secmanagers.EqualString(passOne, passTwo) {
		log.Default().Printf(">>> %v - Passwords are different\n", r.RemoteAddr) //TODO :Should print error to frontend
		http.Redirect(w, r, "/", http.StatusAccepted)
	}

	// Setting cookies to pass to the check-code section
	log.Default().Printf(">> %v - Setting up the cookies used in code validation\n", r.RemoteAddr)
	nameCookie := secmanagers.CreateSecCk("name", r.FormValue("namefield"))
	http.SetCookie(w, &nameCookie)
	surnameCookie := secmanagers.CreateSecCk("surname", r.FormValue("surnamefield"))
	http.SetCookie(w, &surnameCookie)
	email := secmanagers.CreateSecCk("email", mailOne)
	http.SetCookie(w, &email)
	password := secmanagers.CreateSecCk("password", secmanagers.HashNSault(passOne))
	http.SetCookie(w, &password)
	reqtype := http.Cookie{Name: "reqtype", Value: "register"}
	http.SetCookie(w, &reqtype)

	// Sending mail with code and setting cookie to responsewriter
	log.Default().Printf(">> %v - Trying to send code mail...\n", r.RemoteAddr)
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
