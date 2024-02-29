package handlers

import (
	"fmt"
	"net/http"
)

// Catches and responds to login.html
func LoginReqHanlder(w http.ResponseWriter, r *http.Request) {

	// Retrieves cookies for authentication
	ck := r.Cookies()
	for c := range ck {
		fmt.Println(ck[c].Value)
	}
}
