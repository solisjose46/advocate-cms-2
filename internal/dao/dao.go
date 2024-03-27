package dao

import (
	"fmt"
	"database/sql"
	"crypto/sha256"
    _ "github.com/mattn/go-sqlite3"
)

const (
	cmsDbPath = "db/cms.db"
	advocateDbPath = "db/advocate.db"
	saltSuffix = ":salt"
	loginQuery = "SELECT Users.password FROM Users WHERE Users.username = ?"
)

type Dao struct {
	Cms *sql.DB
	Advocate *sql.DB
}

func DatabaseInit() (*Dao, error) {
	// open cms database
	cmsDb, err  := sql.Open("sqlite3", cmsDbPath)
	if err != nil {
		fmt.Println("Error trying to open cms database.")
		return nil, err
	}
    
	// open advocate database
	advocateDb, err := sql.Open("sqlite3", advocateDbPath)
	if err != nil {
		fmt.Println("Error trying to open advocate database.")
		cmsDb.Close()
		return nil, err
	}

	return &Dao{
		Cms: cmsDb,
		Advocate: advocateDb,
	}, nil
}

func (db *Dao) IsValidLogin(username, password string) (bool, error) {
	userRow, err := db.Cms.Query(loginQuery, username)
	
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