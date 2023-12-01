package main

import (
	"docs/handlers"
	"docs/utilities"
	"net/http"
)

func main() {

	println(utilities.InString("", "hsydonadsgyabuz"))

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/loginrequest", handlers.LoginHandler)
	mux.HandleFunc("/registrationrequest", handlers.RegistrationRequest)
	mux.HandleFunc("/codevalidator", handlers.CodeValidatorHandler)
	mux.HandleFunc("/datacollection", handlers.DataCollection)
	mux.HandleFunc("/searchdocs", handlers.SearchDocs)

	e := http.ListenAndServe(":3333", mux)
	if e != nil {
		panic(e)
	}
}
