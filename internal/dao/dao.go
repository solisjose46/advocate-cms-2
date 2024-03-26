package dao

import (
	"fmt"
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
)

const (
	cmsDbPath = "db/cms.db"
	advocateDbPath = "db/cms.db"
)

var cmsDatabase *sql.DB
var advocateDatabase *sql.DB

func DatabaseInit() error {
	// open cms database
	cmsDatabase, err  := sql.Open("sqlite3", cmsDbPath)
	if err != nil {
		fmt.Println("Error trying to open cms database.")
		return err
	}
	defer cmsDatabase.Close()
    
	// open advocate database
	advocateDatabase, err := sql.Open("sqlite3", advocateDbPath)
	if err != nil {
		fmt.Println("Error trying to open advocate database.")
		return err
	}
	defer advocateDatabase.Close()

	return nil
}