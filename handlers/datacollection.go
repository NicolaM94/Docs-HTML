package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

// Base data collection in data.html
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

// Function used to filter the documents and parse them back to data.html
func SearchDocs(w http.ResponseWriter, r *http.Request) {
	// Retrieves the search pattern wanted from the user
	searchPattern := r.FormValue("searchpattern")
	// Retrieves email from the cookies
	email, err := utilities.DecodeSecureCookie("email", r)
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
	fmt.Println(fiscalCode)
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

// Base selection sort algorithm used to order documents
func selectionSort(array []utilities.Document, sortType string) []utilities.Document {
	for i := 0; i <= len(array)-2; i++ {
		if sortType == "name" {
			minValue := i + 1
			for j := i + 2; j < len(array); j++ {
				if array[j].Name < array[minValue].Name {
					minValue = j
				}
			}
			if array[i].Name > array[minValue].Name {
				temp := array[i]
				array[i] = array[minValue]
				array[minValue] = temp
			}
		} else if sortType == "size" {
			minValue := i + 1
			for j := i + 2; j < len(array); j++ {
				if array[j].Size < array[minValue].Size {
					minValue = j
				}
			}
			if array[i].Size > array[minValue].Size {
				temp := array[i]
				array[i] = array[minValue]
				array[minValue] = temp
			}
		} else if sortType == "type" {
			minValue := i + 1
			for j := i + 2; j < len(array); j++ {
				if array[j].Type < array[minValue].Type {
					minValue = j
				}
			}
			if array[i].Type > array[minValue].Type {
				temp := array[i]
				array[i] = array[minValue]
				array[minValue] = temp
			}
		} else if sortType == "date" {
			minValue := i + 1
			for j := i + 2; j < len(array); j++ {
				if array[j].Date < array[minValue].Date {
					minValue = j
				}
			}
			if array[i].Date > array[minValue].Date {
				temp := array[i]
				array[i] = array[minValue]
				array[minValue] = temp
			}
		}

	}
	return array
}

// Function to order documents and parse them back to the database
func OrderDocs(w http.ResponseWriter, r *http.Request) {
	option := r.FormValue("orderby")

	email, err := utilities.DecodeSecureCookie("email", r)
	if err != nil {
		panic(err)
	}

	rows, err := utilities.QueryRow("select * from users")
	if err != nil {
		panic(err)
	}

	fiscalCode := ""
	for r := range rows {
		if rows[r].Email == email["email"] {
			fiscalCode = rows[r].FiscalCode
		}
	}

	docs, err := utilities.CollectDocuments(utilities.GetSettings().ContentPath + string(os.PathSeparator) + fiscalCode)
	if err != nil {
		panic(err)
	}

	prep := utilities.DataToPass{Email: email["email"], Data: selectionSort(docs, option)}

	t, _ := template.ParseFiles("./static/data.html")
	t.Execute(w, prep)
}
