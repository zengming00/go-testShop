package models

import (
	"database/sql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./models/test.db")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
}
