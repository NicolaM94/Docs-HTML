package utilities

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Base insertion of a row into the upchardb.db
func InsertRow(username, password string) error {
	db, err := sql.Open("sqlite3", "upchardb.db")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO users (username,password) VALUES(?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, password)
	if err != nil {
		return err
	}
	return nil
}

// Base delition of a row from the upchardb.db
func DeleteRow(username, password string) error {
	db, err := sql.Open("sqlite3", "upchardb.db")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("DELETE from users where username=? and password=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
