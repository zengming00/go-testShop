package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() *sql.DB {
	var db, err = sql.Open("sqlite3", "./models/test.db")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	return db
}
