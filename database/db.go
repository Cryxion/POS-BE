package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "yaposUser"
	password = "yaposUser"
	dbName   = "YAPOS"
)

var db *sql.DB
var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

func InitDB() error {
	var err error
	fmt.Sprintln("Connecting")
	db, err = sql.Open("postgres", connectString)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
