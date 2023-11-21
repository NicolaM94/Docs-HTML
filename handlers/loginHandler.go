package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("usernamefield")
	password := r.FormValue("passwordfield")

	fmt.Println(utilities.Hasher(username + password))
	fmt.Println(utilities.HashNSault(username + password))

	w.Write([]byte("LoginHandler"))
}
