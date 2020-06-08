package database

import (
	"database/sql"
	"log"
	"time"
)

// DBConn :
var DBConn *sql.DB

// SetupDatabase :
func SetupDatabase() {
	var err error

	DBConn, err = sql.Open("mysql", "root:3210@tcp(127.0.0.1:3306)/gocrud")

	if err != nil {
		log.Fatal(err)
	}
	DBConn.SetMaxOpenConns(3)
	DBConn.SetMaxIdleConns(3)
	DBConn.SetConnMaxLifetime(60 * time.Second)
}
