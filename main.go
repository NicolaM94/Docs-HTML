package main

import (
	"docs/handlers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/loginrequest", handlers.LoginHandler)
	mux.HandleFunc("/registrationrequest", handlers.RegistrationRequest)
	mux.HandleFunc("/codevalidator", handlers.CodeValidatorHandler)
	mux.HandleFunc("/datacollection", handlers.DataCollection)
	mux.HandleFunc("/searchdocs", handlers.SearchDocs)
	mux.HandleFunc("/orderdocs", handlers.OrderDocs)
	mux.HandleFunc("/download", handlers.DownloadHandler)

	e := http.ListenAndServe(":3333", mux)
	if e != nil {
		panic(e)
	}
}
