package managers

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func checkUDBExistance() bool {
	var location string = Settings{}.Populate().UDBLocation
	_, err := os.Stat(location)
	if errors.Is(err, os.ErrNotExist) {
		log.Default().Println("** WARNING ** : File not found. Trying to create one in the next call...")
		return false
	}
	return true
}

// Function used in the main one  to check for the initialization of the users database.
// Uses checkUDBExistance to verify the existance of the file in the first place.
// If it does not find it, creates a file with the path and name stored in the settings file.
// Then it uses the sql api calls to create the statement and execute it.
//
// Public
func InitUserDatabase() error {
	if !checkUDBExistance() {
		log.Default().Println("** WARNING ** : Catching exeption from the previous function. Trying to create a new db file...")

		// Should create the fi
		_, err := os.Create(Settings{}.Populate().UDBLocation)
		if err != nil {
			log.Default().Println("** WARNING ** : Cannot create UDB file in the given settings location. Aborting...")
			panic(err)
		}

		// Should open the file with sql api
		db, err := sql.Open("sqlite3", Settings{}.Populate().UDBLocation)
		if err != nil {
			log.Default().Fatalf("Cannot sql open %v. This error was found %v: \n", Settings{}.Populate().UDBLocation, err)
		}

		// Prepare statement to first writing
		statement, err := db.Prepare(`CREATE TABLE "users" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"USERMAIL"	TEXT NOT NULL,
			"PASSWORD"	TEXT NOT NULL,
			"NAME"	TEXT,
			"SURNAME"	TEXT,
			PRIMARY KEY('ID',"ID" AUTOINCREMENT)
		);`)
		if err != nil {
			return errors.New("** FATAL ERRROR ** : Some errors occured while preparing the db init stmt :" + err.Error())
		}

		res, err := statement.Exec()
		if err != nil {
			return errors.New("** FATAL ERRROR ** : Some errors occured while executing the db init stmt :" + err.Error())
		}

		_, err = res.RowsAffected()
		if err != nil {
			return errors.New("** FATAL ERRROR ** : Some errors occured while retrieving db affected rows  :" + err.Error())
		}
	}
	log.Default().Printf("UDB generated at %v. No exceptions caught.\n", Settings{}.Populate().UDBLocation)
	return nil
}

// Base struct to catch udb rows
type UDBrow struct {
	Id       int
	Mail     string
	Password string
	Name     string
	Surname  string
}

// Queries to the UDB.
// Returns an array of UDBrow instances and error
//
// If err != nil, UDBrow is nil
// Else err is nil and UDBrow will containt the rows upcoming from the query.
func NormalQueryDB(query string) ([]UDBrow, error) {
	db, err := sql.Open("sqlite3", Settings{}.Populate().UDBLocation)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	container := []UDBrow{}
	for rows.Next() {
		temp := UDBrow{}
		err = rows.Scan(&temp.Id, &temp.Mail, &temp.Password, &temp.Name, &temp.Surname)
		if err != nil {
			return nil, err
		}
		container = append(container, temp)
	}
	// If no row is found for the given query, implement a new error stating that.
	// Return nil as container
	if len(container) == 0 {
		return nil, errors.New("No row found for the given query")
	}
	return container, nil
}

// Register a new user into the databases
func RegisterUserUDB(mail, password, name, surname string) error {

	// Verify that user is not already present
	users, err := NormalQueryDB("select * from users")
	if err != nil {
		return err
	}
	for u := range users {
		if users[u].Mail == mail {
			return errors.New("user already present")
		}
	}

	// Starts registration process
	db, err := sql.Open("sqlite3", Settings{}.Populate().UDBLocation)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO users(ID, USERMAIL, PASSWORD, NAME, SURNAME) values (?,?,?,?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(nil, mail, password, name, surname)
	if err != nil {
		return err
	}
	rws, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Default().Print(">> Rows affecteb by db insertion: ", rws)
	return nil
}
