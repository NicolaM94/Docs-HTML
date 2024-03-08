package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"html/template"
	"net/http"
)

// Utility to filter names
func filterDocsByName(documents []managers.Document, tester string) []managers.Document {
	if len(tester) == 0 {
		return documents
	}
	out := []managers.Document{}
	for d := range documents {
		current := documents[d].Name
		for n := 0; n <= len(current)-len(tester); n++ {
			if current[n:n+len(tester)] == tester {
				out = append(out, documents[d])
			}
		}
	}
	return out
}

// Function used by datadelivery html to filter documents by specific strings in their name
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	searchstring := r.FormValue("searchinput")
	emailck, err := r.Cookie("email")
	if err != nil {
		panic(err)
	}
	email := secmanagers.DecodeSecCk(*emailck)

	user, err := managers.QueryByMail(email)
	if err != nil {
		panic(err)
	}
	foldername := user[0].Name + "_" + user[0].Surname

	docs, err := managers.CollectDocuments(foldername)
	if err != nil {
		panic(err)
	}
	filtered := filterDocsByName(docs, searchstring)
	t, _ := template.ParseFiles("./static/data.html")
	data := Data{Username: email, Documents: filtered}
	t.Execute(w, data)
}
