package dao

import (
	"fmt"
	"database/sql"
	"crypto/sha256"
    _ "github.com/mattn/go-sqlite3"
)

const (
	cmsDBPath = "db/cms.db"
	advDBPath = "db/advocate.db"
	saltSuffix = ":salt"
	loginQuery = "SELECT Users.password FROM Users WHERE Users.username = ?"
)

type Dao struct {
	cmsDB *sql.DB
	advDB *sql.DB
}

func DatabaseInit() (*Dao, error) {
	// open cms database
	cmsDB, err  := sql.Open("sqlite3", cmsDBPath)
	if err != nil {
		fmt.Println("Error trying to open cms database.")
		return nil, err
	}
    
	// open advocate database
	advDB, err := sql.Open("sqlite3", advDBPath)
	if err != nil {
		fmt.Println("Error trying to open advocate database.")
		cmsDB.Close()
		return nil, err
	}

	return &Dao{
		cmsDB: cmsDB,
		advDB: advDB,
	}, nil
}

func (db *Dao) CloseDatabase() {
	if db.cmsDB != nil {
		db.cmsDB.Close()
	}

	if db.advDB != nil {
		db.advDB.Close()
	}
}

func (db *Dao) IsValidLogin(username, password string) (bool, error) {
	userRow, err := db.cmsDB.Query(loginQuery, username)
	
	if err != nil {
		fmt.Println("Error with login query")
		return false, err
	}

	defer userRow.Close()

	// hash and salt password
	hasher := sha256.New()
	hasher.Write([]byte(password + saltSuffix))
	hashedBytes := hasher.Sum(nil)
	hashedUserPassword := fmt.Sprintf("%x", hashedBytes)

	var dbPassword string

	for userRow.Next() {
		userRow.Scan(&dbPassword)
	}

	return dbPassword == hashedUserPassword, nil
}