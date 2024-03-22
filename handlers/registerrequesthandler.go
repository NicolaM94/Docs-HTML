package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"html/template"
	"log"
	"net/http"
)

func RegReqHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieving form infos
	mailOne := r.FormValue("emailfield")
	passOne := r.FormValue("passfield")
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
	codecookie := secmanagers.CreateSecCk("code", code)
	http.SetCookie(w, &codecookie)
	err := managers.SendCodeMail(mailOne, code)
	if err != nil {
		log.Fatalln(err)
	}

	// Passing to code-validation
	t, _ := template.ParseFiles("./static/confirmcode.html")
	t.Execute(w, nil)

}
