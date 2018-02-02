package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var Cates *CatesModel
var Goods *GoodsModel

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./models/test.db")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	Cates = NewCatesModel(db)
	Goods = NewGoodsModel(db)
}
