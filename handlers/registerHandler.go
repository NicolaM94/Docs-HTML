package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
	"text/template"
)

func RegistrationRequest(w http.ResponseWriter, r *http.Request) {

	var name string = r.FormValue("name")
	var surname string = r.FormValue("surname")
	var email string = r.FormValue("email")
	var password string = r.FormValue("password")
	var repassword string = r.FormValue("repassword")

	// Checks if passwords are the same. If not, redirect to pwmismatch.html
	if password != repassword {
		http.Redirect(w, r, "./register-pwmismatch.html", http.StatusFound)
	}

	var code string = utilities.GenerateCode()

	http.SetCookie(w, &http.Cookie{Name: "name", Value: name})
	http.SetCookie(w, &http.Cookie{Name: "surname", Value: surname})
	http.SetCookie(w, &http.Cookie{Name: "email", Value: email})
	http.SetCookie(w, &http.Cookie{Name: "password", Value: password})

	fmt.Println("SendCodeMail complains about: ", utilities.SendCodeMail(email, code))

	t, _ := template.ParseFiles("./static/confirmcode.html")
	t.Execute(w, nil)
}
