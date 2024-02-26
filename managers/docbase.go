package managers

import (
	"errors"
	"log"
	"os"
)

// Checks if docbase exists in  the first place.
//
// Internal to mangers module
func docbaseExists() bool {
	var settings Settings = Settings{}.Populate()
	info, err := os.Stat(settings.DocBasePath)
	if errors.Is(err, os.ErrNotExist) {
		log.Default().Println("** WARNING **: Docbase does not exists. Initializing a directory in the next call...")
		return false
	}
	if !info.IsDir() {
		log.Default().Println("** WARNING **: Docbase found but not as a dir. Initializing a directory in the next call...")
		return false
	}
	return true
}

// Uses docbaseExist function to check wheather docbase folder exists. If not, tries to create one.
//
// Public
func CheckDocBase() error {
	if !docbaseExists() {
		log.Default().Println("WARNING: Catching error from previous function. Trying to create a folder...")
		err := os.Mkdir(Settings{}.Populate().DocBasePath, os.ModePerm)
		if err != nil {
			return errors.New("Error called by CheckDockBase on line 36: " + err.Error())
		}
	}
	log.Default().Println("Docbase folder correctly created.")
	return nil
}
