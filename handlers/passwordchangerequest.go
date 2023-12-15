package handlers

import (
	"docs/utilities"
	"net/http"
	"text/template"
)

func PwdChangeRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve mail from field
	mail := r.FormValue("mailfield")

	// Verify if mail in database
	rows, err := utilities.QueryRow("select * from users")
	if err != nil {
		panic(err)
	}

	presence := utilities.SearchInRows(mail, rows)

	if !presence {
		t, _ := template.ParseFiles("./static/not-registered.html")
		t.Execute(w, nil)
		return
	}

	ck := utilities.GenerateSecureCookie("mail", mail)
	http.SetCookie(w, ck)

	code := utilities.GenerateCode()
	ck = utilities.GenerateSecureCookie("code", code)
	http.SetCookie(w, ck)

	err = utilities.SendCodeMail(mail, code)
	if err != nil {
		panic(err)
	}

	t, _ := template.ParseFiles("./static/confirmcodepassreset.html")
	t.Execute(w, nil)

}
