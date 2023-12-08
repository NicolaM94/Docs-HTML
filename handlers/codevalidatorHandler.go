package handlers

import (
	"docs/utilities"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func CodeValidatorHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieve insertedCode from form and decodes code cookie
	var insertedCode string = r.FormValue("inserted-code")
	cookieCode, err := utilities.DecodeSecureCookie("code", r)
	if err != nil {
		fmt.Println("Cannot decode code cookie")
	}

	// Checks if the two codes are equal. If not, redirect to register page
	if insertedCode != cookieCode["code"] {
		http.Redirect(w, r, "/register.html", http.StatusFound)
		return
	}

	// Decode cookies and register values in the database
	fmt.Println("Registration request")
	cookieName, err := utilities.DecodeSecureCookie("name", r)
	if err != nil {
		panic(err)
	}
	cookieSurname, err := utilities.DecodeSecureCookie("surname", r)
	if err != nil {
		panic(err)
	}
	cookieFiscalCode, err := utilities.DecodeSecureCookie("fiscalcode", r)
	if err != nil {
		panic(err)
	}
	cookieMail, err := utilities.DecodeSecureCookie("email", r)
	if err != nil {
		panic(err)
	}
	cookiePassword, err := utilities.DecodeSecureCookie("password", r)
	if err != nil {
		panic(err)
	}
	err = utilities.InsertRow(cookieName["name"], cookieSurname["surname"], cookieFiscalCode["fiscalcode"], cookieMail["email"], cookiePassword["password"])
	if err != nil {
		panic(err)
	}

	// Check if folder already present: if so panic. Then create user folder to store documents in
	existance, err := utilities.VerifyFolderExistance(utilities.GetSettings().ContentPath + string(os.PathSeparator) + cookieFiscalCode["fiscalcode"])
	if err != nil {
		panic(err)
	}
	if !existance {
		fmt.Println("Folder for ", cookieFiscalCode["fiscalcode"], " does not exist. Creating one...")
		err = os.Mkdir(utilities.GetSettings().ContentPath+string(os.PathSeparator)+cookieFiscalCode["fiscalcode"], 0777)
		if err != nil {
			panic(err)
		}
	}
	//Puts welcome guide into folder
	rdr, err := os.ReadFile("Benvenuto.pdf")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(utilities.GetSettings().ContentPath+string(os.PathSeparator)+cookieFiscalCode["fiscalcode"]+string(os.PathSeparator)+"Benvenuto.pdf", rdr, 0777)
	if err != nil {
		panic(err)
	}

	// Parses confirm
	t, _ := template.ParseFiles("./static/registration-confirm.html")
	t.Execute(w, nil)
}
