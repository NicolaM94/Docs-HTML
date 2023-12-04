package utilities

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Base insertion of a row into the upchardb.db
func InsertRow(name, surname, fiscalcode, email, password string) error {
	db, err := sql.Open("sqlite3", GetSettings().DBFilePath)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO users (name,surname,fiscalcode,email,password) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, surname, fiscalcode, email, password)
	if err != nil {
		return err
	}
	return nil
}

// Base delition of a row from the upchardb.db
func DeleteRow(username, password string) error {
	db, err := sql.Open("sqlite3", GetSettings().DBFilePath)
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

// Query db with statement
type Row struct {
	Id         int
	Name       string
	Surname    string
	FiscalCode string
	Email      string
	Password   string
}

func QueryRow(statement string) ([]Row, error) {
	db, err := sql.Open("sqlite3", GetSettings().DBFilePath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	var collector []Row
	var entry Row
	for rows.Next() {
		err = rows.Scan(&entry.Id, &entry.Name, &entry.Surname, &entry.FiscalCode, &entry.Email, &entry.Password)
		if err != nil {
			return nil, err
		}
		collector = append(collector, entry)
	}
	return collector, nil
}

// Search functions in rows
func SearchInRows(target string, rows []Row) bool {
	for g := range rows {
		curretRow := rows[g]
		if target == curretRow.FiscalCode || target == curretRow.Email {
			return true
		}
	}
	return false
}
