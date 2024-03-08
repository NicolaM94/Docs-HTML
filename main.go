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
	log.Default().Println("Check for settings...")
	err := managers.CheckSettings()
	if err != nil {
		log.Fatalln(err)
	}

	// Checks for data file system existance.
	// If false, fatal.
	log.Default().Printf("Check for docbase in %v...\n", managers.Settings{}.Populate().DocBasePath)
	err = managers.CheckDocBase()
	if err != nil {
		log.Fatalln(err)
	}

	//Check for users database existance
	log.Default().Printf("Checkign for udb in %v...\n", managers.Settings{}.Populate().UDBLocation)
	err = managers.InitUserDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	// Defining routes
	log.Default().Println("Defining routes...")
	static := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.Handle("/", static)
	mux.HandleFunc("/searchdata", handlers.SearchHandler)
	mux.HandleFunc("/datadelivery", handlers.DataDeliveryHandler)
	mux.HandleFunc("/loginrequest", handlers.LoginReqHanlder)
	mux.HandleFunc("/codevalidation", handlers.CodeValidationHandler)
	mux.HandleFunc("/registerrequest", handlers.RegReqHandler)
	mux.HandleFunc("/index", handlers.IndexHandler)

	// Starting the server
	log.Default().Println("Server started. Awaiting connections...")
	srvErr := http.ListenAndServe(":"+managers.Settings{}.Populate().ServerPort, mux)
	if srvErr != nil {
		log.Fatalln(err)
	}

}
