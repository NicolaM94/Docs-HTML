package handlers

import (
	"docs/utilities"
	"net/http"
)

func DataCollection(w http.ResponseWriter, r *http.Request) {

	// Retrieve email cookie
	ck, err := utilities.DecodeSecureCookie("email", r)
	if err != nil {
		panic(err)
	}
	ckValue := ck["mail"]
	settings := utilities.GetSettings()

	// TODO: Qua

}
