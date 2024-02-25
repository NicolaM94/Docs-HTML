package main

import (
	"docshelf/handlers"
	"docshelf/managers"
	"log"
	"net/http"
)

func main() {

	// Checks for settings file existance.
	// If false, fatal.
	// Then instantiate settings.
	err := managers.CheckSettings()
	if err != nil {
		log.Fatalln(err)
	}

	// Checks for data file system existance.
	// If false, fatal.

	static := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.Handle("/", static)
	mux.HandleFunc("/index", handlers.IndexHandler)

	// Starting the server
	srvErr := http.ListenAndServe(":3333", mux)
	if srvErr != nil {
		panic(srvErr)
	}

}
