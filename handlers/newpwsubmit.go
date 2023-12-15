package handlers

import (
	"docs/utilities"
	"net/http"
	"text/template"
)

func NewPwSubmitHandler(w http.ResponseWriter, r *http.Request) {

	//TODO: Serve un modo per autenticare la richiesta, altrimenti uno arriva da plain html e cambia Forse un cookie di auth per il cambio password?
	// Retrieves mail for latter search of the user
	_, err := utilities.DecodeSecureCookie("mail", r)
	if err != nil {
		panic(err)
	}

	// Check for password match. In case of failure, return the pass submission again with an error
	passwordOne := r.FormValue("password-one")
	passwordTwo := r.FormValue("password-two")
	if passwordOne != passwordTwo {
		t, _ := template.ParseFiles("./static/newpwsubmiterror.html")
		t.Execute(w, nil)
		return
	}

	//TODO: Scrivi il cambio passowrd qui

}
