package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
	"text/template"
)

// Comes from registration.html, handles password match verification, sets cookies for temporary storage and parse confirmcode.html
func RegistrationRequest(w http.ResponseWriter, r *http.Request) {

	var name string = r.FormValue("name")
	var surname string = r.FormValue("surname")
	var fiscalcode string = r.FormValue("fiscalcode")
	var email string = r.FormValue("email")
	var password string = r.FormValue("password")
	var repassword string = r.FormValue("repassword")

	// Checks if passwords are the same. If not, redirect to pwmismatch.html
	if password != repassword {
		http.Redirect(w, r, "./register-pwmismatch.html", http.StatusFound)
	}

	var code string = utilities.GenerateCode()
	var pwHash string = utilities.HashNSault(name + surname + email + password)

	http.SetCookie(w, utilities.GenerateSecureCookie("name", name))
	http.SetCookie(w, utilities.GenerateSecureCookie("surname", surname))
	http.SetCookie(w, utilities.GenerateSecureCookie("fiscalcode", fiscalcode))
	http.SetCookie(w, utilities.GenerateSecureCookie("email", email))
	http.SetCookie(w, utilities.GenerateSecureCookie("password", pwHash))
	http.SetCookie(w, utilities.GenerateSecureCookie("code", code))
	http.SetCookie(w, utilities.GenerateSecureCookie("type", "registration"))

	fmt.Println("SendCodeMail complains about: ", utilities.SendCodeMail(email, code))

	t, _ := template.ParseFiles("./static/confirmcode.html")
	t.Execute(w, nil)
}
