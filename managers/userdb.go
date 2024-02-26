package managers

import (
	"errors"
	"log"
	"os"
)

func checkUDBExistance() bool {
	var location string = Settings{}.Populate().UDBLocation
	log.Default().Printf("Checking for UDB in %v...\n", location)
	_, err := os.Stat(location)
	if errors.Is(err, os.ErrNotExist) {
		log.Default().Println("** WARNING ** : File not found. Trying to create one in the next call...")
		return false
	}
	return true
}

func InitUserDatabase() error {
	if !checkUDBExistance() {
		log.Default().Println("** WARNING ** : Catching exeption from the previous function. Trying to create a new db file...")

		_, err := os.Create(Settings{}.Populate().UDBLocation)
		if err != nil {
			log.Default().Println("** WARNING ** : Cannot create UDB file in the given settings location. Aborting...")
			panic(err)
		}

		// TODO : Scrivi sql init per la prima inizializzazione del database

	}

}
