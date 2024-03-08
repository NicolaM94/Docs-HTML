package handlers

import (
	"docshelf/managers"
	"docshelf/secmanagers"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Username  string
	Documents []managers.Document
}

func DataDeliveryHandler(w http.ResponseWriter, r *http.Request) {

	log.Default().Printf("%v - Delivering fresh data", r.RemoteAddr)
	email, err := r.Cookie("email")
	if err != nil {
		log.Fatal(err)
	}
	plainmail := secmanagers.DecodeSecCk(*email)

	user, err := managers.QueryByMail(plainmail)
	if err != nil {
		log.Fatal(err)
	}

	if len(user) != 1 {
		log.Fatal("Multiple users found")
	}

	collector, err := managers.CollectDocuments(user[0].Name + "_" + user[0].Surname)
	if err != nil {
		log.Fatal(err)
	}
	for c := range collector {
		fmt.Println(collector[c])
	}

	t, _ := template.ParseFiles("./static/data.html")
	t.Execute(w, Data{Username: plainmail, Documents: collector})

}
