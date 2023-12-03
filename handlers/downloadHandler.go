package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

// function to handle the download requests of the documents

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	// To avoid showing the path of the document, we use the hash id and compare it with the one of the docuements collected
	// by CollectDocuments function.

	// Retrieve the requested ID from the URI
	requestedFile := r.RequestURI[10:]
	fmt.Println("Requested file: ", requestedFile)

	// Decodes the mail cookie to look for the user folder
	email, err := utilities.DecodeSecureCookie("email", r)
	if err != nil {
		panic(err)
	}
	// Collects all the users and store the fiscal code of the correct one to build the collection path
	rows, err := utilities.QueryRow("select * from users")
	if err != nil {
		panic(err)
	}
	fiscalcode := ""
	for r := range rows {
		if rows[r].Email == email["email"] {
			fiscalcode = rows[r].FiscalCode
			break
		}
	}

	// Builds the download folder and collect documents
	downloadFolder := utilities.GetSettings().ContentPath + string(os.PathSeparator) + fiscalcode
	docs, err := utilities.CollectDocuments(downloadFolder)
	if err != nil {
		panic(err)
	}

	// Compares the requested ID and id of each document to find a possible path to it.
	downloadPath := ""
	for d := range docs {
		if docs[d].Id == requestedFile {
			downloadPath = docs[d].Path
			break
		}
	}
	// If download path is still empty after the search, return 404
	if downloadPath == "" {
		t, _ := template.ParseFiles("./static/404.html")
		t.Execute(w, nil)
	}

	// TODO: finish downloader

}
