package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func DataCollection(w http.ResponseWriter, r *http.Request) {

	// Retrieve email cookie
	ck, err := utilities.DecodeSecureCookie("email", r)
	if err != nil {
		panic(err)
	}
	ckValue := ck["email"]

	rows, err := utilities.QueryRow("select * from users")
	if err != nil {
		panic(err)
	}
	fiscalCode := ""
	for r := range rows {
		if rows[r].Email == ckValue {
			fiscalCode = rows[r].FiscalCode
			break
		}
	}

	docs, err := utilities.CollectDocuments(utilities.GetSettings().ContentPath + string(os.PathSeparator) + fiscalCode)
	if err != nil {
		panic(err)
	}
	fmt.Println("Collected documents: ", len(docs))

	dataToPass := utilities.DataToPass{
		Email: ckValue,
		Data:  docs,
	}

	t, _ := template.ParseFiles("./static/data.html")
	t.Execute(w, dataToPass)

}
