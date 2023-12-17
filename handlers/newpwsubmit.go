package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"
)

func NewPwSubmitHandler(w http.ResponseWriter, r *http.Request) {

	//TODO: Serve un modo per autenticare la richiesta, altrimenti uno arriva da plain html e cambia Forse un cookie di auth per il cambio password?
	// Checks auth cookie for the password change
	authck, err := utilities.DecodeSecureCookie("pwAthTkn", r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	codecookie, err := utilities.DecodeSecureCookie("code", r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	if authck["pwAthTkn"] != utilities.HashNSault(codecookie["code"]+time.DateOnly) {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Retrieves mail for latter search of the user
	mail, err := utilities.DecodeSecureCookie("mail", r)
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

	afct, err := utilities.UpdateRow(mail["mail"], utilities.HashNSault(mail["mail"]+passwordOne))
	if err != nil {
		panic(err)
	}
	if afct != 1 {
		fmt.Println(r.Host)
		fmt.Println("Email caught: ", mail["mail"])
		fmt.Println("Users affected: ", afct, "err: ", err)
		fmt.Println("Extremely severe exception caught - Multiple rows affected by password change request. Shutting down for security reason. Investigation neeeded: please contact the support.")
		os.Exit(2)
	}

	utilities.SendPassChangeMail(mail["mail"])
	t, _ := template.ParseFiles("./static/registration-confirm.html")
	t.Execute(w, nil)
}
