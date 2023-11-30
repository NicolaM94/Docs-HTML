package main

import (
	"docs/handlers"
	"docs/utilities"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Test settings")
	settings := utilities.GetSettings()
	fmt.Println("Mail: ", settings.Mail)
	fmt.Println("Path: ", settings.ContentPath)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/loginrequest", handlers.LoginHandler)
	mux.HandleFunc("/registrationrequest", handlers.RegistrationRequest)
	mux.HandleFunc("/codevalidator", handlers.CodeValidatorHandler)
	mux.HandleFunc("/datacollection", handlers.DataCollection)

	e := http.ListenAndServe(":3333", mux)
	if e != nil {
		panic(e)
	}
}
