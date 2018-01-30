package models

import (
	"database/sql"
)

var reUseDb *sql.DB

func GetConn() (*sql.DB, error) {
	if reUseDb != nil {
		return reUseDb, nil
	}
	db, err := sql.Open("sqlite3", "./models/test.db")
	if err != nil {
		return nil, err
	}
	reUseDb = db
	db.SetMaxOpenConns(100)
	return reUseDb, nil
}
