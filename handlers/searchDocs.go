package handlers

import (
	"docs/utilities"
	"net/http"
	"os"
	"text/template"
)

// Function used to filter the documents and parse them back to data.html
func SearchDocs(w http.ResponseWriter, r *http.Request) {
	// Retrieves the search pattern wanted from the user
	searchPattern := r.FormValue("searchpattern")
	// Retrieves email from the cookies
	email, err := utilities.DecodeSecureCookie("mail", r)
	if err != nil {
		panic(err)
	}
	rows, err := utilities.QueryRow("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	// Retrieves fiscal code from mail
	fiscalCode := ""
	for r := range rows {
		if rows[r].Email == email["email"] {
			fiscalCode = rows[r].FiscalCode
		}
	}
	// Collects documents from the folder
	settings := utilities.GetSettings()
	docs, err := utilities.CollectDocuments(settings.ContentPath + string(os.PathSeparator) + fiscalCode)
	if err != nil {
		panic(err)
	}
	// Filters docs based on "searchpattern"
	var filtered []utilities.Document
	for d := range docs {
		if utilities.InString(searchPattern, docs[d].Name) {
			filtered = append(filtered, docs[d])
		}
	}
	// Return
	t, _ := template.ParseFiles("./static/data.html")
	t.Execute(w, utilities.DataToPass{Email: email["email"], Data: filtered})
}
