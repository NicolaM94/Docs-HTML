package main

import (
	"docs/handlers"
	"docs/utilities"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("InsertRows complains about:", utilities.InsertRow("nicola", "djbsgbdsglagfh"))

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/loginrequest", handlers.LoginHandler)
	mux.HandleFunc("/registrationrequest", handlers.RegistrationRequest)

	e := http.ListenAndServe(":3333", mux)
	if e != nil {
		panic(e)
	}
}
