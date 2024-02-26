// Settings module containing all the type and functions to manage settings in DocShelf workflow.
package managers

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

/*
Settings class used in fetching settings.
  - Server[Port] stays for the port of the local server address.
  - ServerMail[Address-Pass] are needed for the authentication in order to send service auth codes.
  - ServerSMTP[Host-Port] are config data for the mail server used. Check online based on which provider is serving you.
  - DocBasePath is the base folder containing all the documents serverd to the users.
*/
type Settings struct {
	ServerPort        string
	ServerMailAddress string
	ServerMailPass    string
	ServerSMTPHost    string
	ServerSMTPPort    string
	DocBasePath       string
	UDBLocation       string
}

// Populates the settings struct instance with the current data from settings.json.
// It shold never log err from os.ReadFile since the settings existance is checked in main before this.
//
// Write back directly in s: one should call this every time a Settings struct is instantiated.
func (s Settings) Populate() Settings {
	reader, err := os.ReadFile("settings.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(reader, &s)
	if err != nil {
		log.Fatalln(err)
	}
	return s
}

// Checks if settings file exists in the path of the executable.
//
// Internal to settings.go module, used in [CheckSettings] func.
func settingsExist() bool {
	var found bool = false
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			if d.Name() == "settings.json" {
				found = true
			}
		}
		return nil
	})
	return found
}

// Function used to check if settings file exists. If it does not, tries to create the file.
// Manages the error calling to panic if any error occurs.
//
// Public to project. Always returns nil.
func CheckSettings() error {
	if !settingsExist() {
		var payload []byte
		payload, err := json.MarshalIndent(Settings{}, "", "    ")
		if err != nil {
			return errors.New("Settings file was not found. Cannot write. Called by CheckSettings, managers module")
		}
		os.WriteFile("settings.json", payload, 0777)
		return errors.New("Settings file was not found. Please restart the server after populating the settings.json file. Called by CheckSettings, managers module")
	}
	return nil
}
