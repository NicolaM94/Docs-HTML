package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Parses the request uri to get the requested file name
// Breaks the requestUri given as parameter after finding a '?'
// Returns the remaining string after it.
// If no '?' is found, returns an error and an empty string.
//
//	ALWAYS CHECK THE ERROR TO AVOID EMPTY FILE NAMES
func parseDownloadPath(requestUri string) (error, string) {
	for i, j := range requestUri {
		if j == '?' {
			return nil, requestUri[i+1:]
		}
	}
	return errors.New("malformed request uri, no breakpoint found"), ""
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	log.Default().Println(">>>>>>>>>>>>> Download handler called")
	authCk, err := r.Cookie("authToken")
	if err != nil {
		t, _ := template.ParseFiles("./static/loggedout.html")
		t.Execute(w, nil)
	}
	authCookie := secmanagers.DecodeSecCk(*authCk)
	tokenPresence, err := managers.IsTokenPresent(authCookie)
	if err != nil || !tokenPresence {
		fmt.Println("Token error: ", err)
		t, _ := template.ParseFiles("./static/loggedout.html")
		t.Execute(w, nil)
	}

	err, request := parseDownloadPath(r.RequestURI)
	if err != nil {
		log.Default().Printf("%v download func - %v", r.RemoteAddr, err)
		http.Redirect(w, r, "/data", http.StatusFound)
	}

	//Build user folder
	ck, err := r.Cookie("email")
	if err != nil {
		panic(err)
	}
	email := secmanagers.DecodeSecCk(*ck)
	usrs, err := managers.QueryByMail(email)
	if err != nil {
		log.Default().Printf("%v download func - %v", r.RemoteAddr, err)
		http.Redirect(w, r, "/data", http.StatusFound)
	}

	downloadPath := managers.Settings{}.Populate().DocBasePath
	if downloadPath[len(downloadPath)-1] != '/' {
		downloadPath = downloadPath + "/" + usrs[0].Name + "_" + usrs[0].Surname + "/" + request
	} else {
		downloadPath = downloadPath + usrs[0].Name + "_" + usrs[0].Surname + "/" + request
	}

	// Reads the file and adds it to the response header
	file, err := os.ReadFile(downloadPath)
	if err != nil {
		log.Default().Printf("%v download func - %v", r.RemoteAddr, err)
		http.Redirect(w, r, "/data", http.StatusFound)
	}
	filename := "attachment;filename=" + request
	w.Header().Set("Content-Disposition", filename)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Write(file)
	http.Redirect(w, r, "/data", http.StatusFound)
}
