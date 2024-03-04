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

// Base function that registers a user in the system
func Register(w http.ResponseWriter, r *http.Request) error {

	// Retrieves cookies needed for the registation
	ck, err := r.Cookie("email")
	if err != nil {
		return err
	}
	email := secmanagers.DecodeSecCk(*ck)

	ck, err = r.Cookie("password")
	if err != nil {
		return err
	}
	password := secmanagers.DecodeSecCk(*ck)

	ck, err = r.Cookie("name")
	if err != nil {
		return err
	}
	name := secmanagers.DecodeSecCk(*ck)

	ck, err = r.Cookie("surname")
	if err != nil {
		return err
	}
	surname := secmanagers.DecodeSecCk(*ck)

	// Checks if the user is already present in the database
	users, err := managers.NormalQueryDB("select * from users")
	if err != nil {
		return err
	}
	for u := range users {
		if users[u].Mail == email {
			// TODO: Manage existing folder, the server cannot fault for this
			return fmt.Errorf("%v Register function - User already registered {%v, %v, %v}", r.RemoteAddr, email, name, surname)
		}
	}

	// Checks if the folder already exists in the docbase folder path
	// Formats the folder path given the settings declaration for it
	settings := managers.Settings{}.Populate()
	folderPath := fmt.Sprintf("%v_%v/", name, surname)
	if settings.DocBasePath[len(settings.DocBasePath)-1] == '/' {
		folderPath = settings.DocBasePath + folderPath
	} else {
		folderPath = settings.DocBasePath + "/" + folderPath
	}
	// Actual check if the folder exists.
	// The foulder should not exist, so a os.ErrNotExists should be thrown.
	// If not, raises a critical warning to the logger and returns to index.
	//TODO : Debug
	_, err = os.Stat(folderPath)
	if !errors.Is(err, os.ErrNotExist) {
		// TODO: Manage existing folder, the server cannot fault for this
		return fmt.Errorf("%v Register function - User folder already present for {%v, %v, %v} in %v", r.RemoteAddr, email, name, surname, folderPath)
	}

	// If the function reached this point, the user and the user folder should not be present.
	// Now tries to register and create the folder
	log.Default().Printf("%v - Starting registration for user\n", r.RemoteAddr)

	// Tries to create the folder
	log.Default().Printf("%v - Creating a folder for user %v %v...\n", r.RemoteAddr, name, surname)
	err = os.Mkdir(folderPath, os.ModePerm)
	if err != nil {
		log.Default().Printf("%v Register function : %v\n", r.RemoteAddr, err)
	}
	log.Default().Printf("%v - User folder created with name %v_%v\n", r.RemoteAddr, name, surname)

	// ... tries to register the user in the udb
	err = managers.RegisterUserUDB(email, password, name, surname)
	if err != nil {
		log.Default().Printf("%v Register function : %v\n", r.RemoteAddr, err)
	}
	log.Default().Printf("%v - User %v %v added to UDB\n", r.RemoteAddr, name, surname)

	// ...route to confirm registration
	t, _ := template.ParseFiles("./static/registration-confirm.html")
	t.Execute(w, nil)
	return nil
}
