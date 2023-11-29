package handlers

import (
	"docs/utilities"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Catches
	email := r.FormValue("emailfield")
	password := r.FormValue("passwordfield")
	tempHash := utilities.HashNSault(email + password)

	rows, err := utilities.QueryRow("SELECT * FROM USERS")
	if err != nil {
		panic(err)
	}
	if len(rows) == 0 {
		panic("empty query to db")
	}

	dbPass := ""
	for r := range rows {
		if rows[r].Email == email {
			dbPass = rows[r].Password
			break
		}
	}
	if tempHash == dbPass {
		w.Write([]byte("LoggedIn"))
		return
	}
}
