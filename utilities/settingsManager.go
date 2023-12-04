package utilities

import (
	"encoding/json"
	"os"
)

// Class used from GetSettings func below
type Settings struct {
	Mail        string
	Password    string
	ServerSMTP  string
	PortSMTP    string
	ContentPath string
	LogFilePath string
}

// Retrieves settings from .json file and stores them into a class
func GetSettings() Settings {
	rdr, err := os.ReadFile("settings.json")
	if err != nil {
		panic(err)
	}
	var s Settings
	err = json.Unmarshal(rdr, &s)
	if err != nil {
		panic(err)
	}
	return s
}
