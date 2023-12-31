package handlers

import (
	"docs/utilities"
	"net/http"
	"text/template"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Catches email and password from the previous form and calculates the hash
	email := r.FormValue("emailfield")
	password := r.FormValue("passwordfield")
	tempHash := utilities.HashNSault(email + password)
	print(tempHash)

	ck := utilities.GenerateSecureCookie("email", email)
	http.SetCookie(w, ck)

	// Queries the database to get the rows, panic if resul is empty or the lenght is 0
	rows, err := utilities.QueryRow("SELECT * FROM USERS")
	if err != nil {
		panic(err)
	}
	if len(rows) == 0 {
		panic("empty query to db")
	}

	// Looks for a result with the same email and stores the password in dbPass
	dbPass := ""
	for r := range rows {
		if rows[r].Email == email {
			dbPass = rows[r].Password
			break
		}
	}
	// Verifies that the password is the same
	if tempHash == dbPass {

		secCookie := utilities.GenerateSecureCookie("authToken", email+time.Now().String())
		secCookie.Expires = time.Now().Add(15 + time.Minute)
		http.SetCookie(w, secCookie)
		http.Redirect(w, r, "/datacollection", http.StatusFound)
	}

	// If password is not the same, return the login page with a warning message.
	t, _ := template.ParseFiles("./static/login-error.html")
	t.Execute(w, nil)
}
