package main

import (
	"docs/handlers"
	"docs/utilities"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println(utilities.HashNSault("mamma" + "papà"))
	fmt.Println(utilities.HashNSault("nonno" + "nonna"))

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/loginrequest", handlers.LoginHandler)
	mux.HandleFunc("/registrationrequest", handlers.RegistrationRequest)
	mux.HandleFunc("/codevalidator", handlers.CodeValidatorHandler)

	e := http.ListenAndServe(":3333", mux)
	if e != nil {
		panic(e)
	}
}
