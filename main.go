package main

import "net/http"

func main() {

	static := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.Handle("/", static)

	// Starting the server
	srvErr := http.ListenAndServe(":3333", mux)
	if srvErr != nil {
		panic(srvErr)
	}

}
