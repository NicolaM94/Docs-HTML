package managers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Function used to check if the db file exists in the location set in the settings.json
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

		// Prepare statement to write users table
		statement, err := db.Prepare(`CREATE TABLE "users" (
			"USERMAIL"	TEXT NOT NULL,
			"PASSWORD"	TEXT NOT NULL,
			"NAME"	TEXT NOT NULL,
			"SURNAME"	TEXT NOT NULL
		);`)
		if err != nil {
			return errors.New("** FATAL ERRROR ** : Some errors occured while preparing the usr table init stmt :" + err.Error())
		}

		res, err := statement.Exec()
		if err != nil {
			return errors.New("** FATAL ERRROR ** : Some errors occured while executing the usr table init stmt :" + err.Error())
		}

		_, err = res.RowsAffected()
		if err != nil {
			return errors.New("** FATAL ERRROR ** : Some errors occured while retrieving usr table affected rows  :" + err.Error())
		}

		db.Close()

		db, err = sql.Open("sqlite3", Settings{}.Populate().UDBLocation)
		if err != nil {
			log.Default().Fatalf("Cannot sql open %v. This error was found %v: \n", Settings{}.Populate().UDBLocation, err)
		}

		// Prepare statement to write auth tokens table
		statement, err = db.Prepare(`CREATE TABLE "tokens" (
			"TOKEN"	TEXT NOT NULL,
			"TTL"	BLOB NOT NULL
		);`)

		if err != nil {
			return errors.New("** FATAL ERRROR ** : Some errors occured while preparing the tokens table init stmt :" + err.Error())
		}

		res, err = statement.Exec()
		if err != nil {
			return errors.New("** FATAL ERRROR ** : Some errors occured while executing the tokens table init stmt :" + err.Error())
		}

		_, err = res.RowsAffected()
		if err != nil {
			return errors.New("** FATAL ERRROR ** : Some errors occured while retrieving tokens table affected rows  :" + err.Error())
		}

		db.Close()

	}
	log.Default().Printf("UDB generated at %v. No exceptions caught.\n", Settings{}.Populate().UDBLocation)
	return nil
}

// Base struct to catch udb rows
type UDBrow struct {
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
		err = rows.Scan(&temp.Mail, &temp.Password, &temp.Name, &temp.Surname)
		if err != nil {
			return nil, err
		}
		container = append(container, temp)
	}
	// If no row is found for the given query, implement a new error stating that.
	// Return nil as container
	return container, nil
}

func QueryByMail(mail string) ([]UDBrow, error) {
	db, err := sql.Open("sqlite3", Settings{}.Populate().UDBLocation)
	if err != nil {
		log.Fatal(err)
	}
	qry, err := db.Query(fmt.Sprintf(`SELECT * FROM users WHERE usermail="%v"`, mail))
	if err != nil {
		log.Fatal(err)
	}
	collector := []UDBrow{}
	for qry.Next() {
		temp := UDBrow{}
		qry.Scan(&temp.Mail, &temp.Password, &temp.Name, &temp.Surname)
		collector = append(collector, temp)
	}
	if len(collector) == 0 {
		return nil, errors.New("no users found with the given email")
	}
	if len(collector) > 1 {
		return nil, errors.New("!!! SERIOUS WARNING !!!! multiple users found with the given mail. Shutdown the server and check for any duplicates")
	}
	return collector, nil
}

// Register a new user into the databases
func RegisterUserUDB(mail, password, name, surname string) error {

	// Verify that user is not already present
	users, err := NormalQueryDB("select * from users")
	if err != nil {
		return fmt.Errorf("RegisterUserDB: %v", err)
	}
	if len(users) != 0 {
		for u := range users {
			if users[u].Mail == mail {
				return fmt.Errorf("RegisterUserDB: %v", errors.New("user already present"))
			}
		}
	}

	// Starts registration process
	db, err := sql.Open("sqlite3", Settings{}.Populate().UDBLocation)
	if err != nil {
		return fmt.Errorf("RegisterUserDB: %v", err)
	}
	stmt, err := db.Prepare("INSERT INTO users(USERMAIL, PASSWORD, NAME, SURNAME) values (?,?,?,?)")
	if err != nil {
		return fmt.Errorf("RegisterUserDB: %v", err)
	}
	res, err := stmt.Exec(mail, password, name, surname)
	if err != nil {
		return fmt.Errorf("RegisterUserDB: %v", err)
	}
	rws, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("RegisterUserDB: %v", err)
	}
	log.Default().Print(">> Rows affecteb by db insertion: ", rws)
	return nil
}

type TokenEntry struct {
	Token string
	Ttl   string
}

// Register a new auth token to db
func RegisterToken(token string, ttl time.Time) error {

	db, err := sql.Open("sqlite3", Settings{}.Populate().UDBLocation)
	if err != nil {
		log.Fatal("Cannot write token to DB: ", err)
	}
	stmt, err := db.Prepare("INSERT INTO tokens(token, ttl) values(?,?)")
	if err != nil {
		log.Fatal("Cannot prepare token to DB: ", err)
	}
	res, err := stmt.Exec(token, ttl)
	if err != nil {
		log.Fatal("Cannot execute statement to DB: ", err)
	}
	rws, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Cannot retrieve rows affected by DB call: ", err)
	}
	if rws == 0 {
		log.Fatal("No rows affected after DB insertion. Check for errors.")
	}
	return nil
}

// Function to check wheater the token is present in the database
func IsTokenPresent(token string) (bool, error) {
	db, err := sql.Open("sqlite3", Settings{}.Populate().UDBLocation)
	if err != nil {
		return false, err
	}
	query, err := db.Query(fmt.Sprintf(`select * from tokens where token="%v"`, token))
	if err != nil {
		return false, fmt.Errorf("query error: %v", err)
	}

	counter := 0
	for query.Next() {
		counter++
	}
	if counter != 1 {
		return false, err
	}
	return true, nil
}
