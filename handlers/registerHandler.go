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

	// Checks if user is already registered
	rows, err := utilities.QueryRow("select * from users")
	if err != nil {
		panic(err)
	}
	if utilities.SearchInRows(email, rows) {
		t, _ := template.ParseFiles("./static/already-registered.html")
		fmt.Println("Already registered")
		t.Execute(w, nil)
		return
	}

	// Generate the code and the password hash
	var code string = utilities.GenerateCode()
	var pwHash string = utilities.HashNSault(email + password)

	// Set cookie for reception in the next page
	http.SetCookie(w, utilities.GenerateSecureCookie("name", name))
	http.SetCookie(w, utilities.GenerateSecureCookie("surname", surname))
	http.SetCookie(w, utilities.GenerateSecureCookie("fiscalcode", fiscalcode))
	http.SetCookie(w, utilities.GenerateSecureCookie("email", email))
	http.SetCookie(w, utilities.GenerateSecureCookie("password", pwHash))
	http.SetCookie(w, utilities.GenerateSecureCookie("code", code))
	http.SetCookie(w, utilities.GenerateSecureCookie("type", "registration"))

	// Sends mail with the code
	fmt.Println("SendCodeMail complains about: ", utilities.SendCodeMail(email, code))

	// Parses the confirmcode page and directs there
	t, _ := template.ParseFiles("./static/confirmcode.html")
	t.Execute(w, nil)
}
