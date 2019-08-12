package sqlite

import (
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

//database used as interface to connect
var database *dbr.Session

//New used for setup connection to sqlite
func New(driver, path string) error {
	db, _ := dbr.Open(driver, path, nil)

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)

	database = db.NewSession(nil)
	err := database.Ping()
	if err != nil {
		return err
	}
	return nil
}

//GetDB used for get current database connection
func GetDB() *dbr.Session {
	return database
}
