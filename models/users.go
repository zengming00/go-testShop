package models

import (
	"database/sql"
)

type UsersModel struct {
	db *sql.DB
}

func NewUsersModel(db *sql.DB) *UsersModel {
	return &UsersModel{db}
}

func (u *UsersModel) Count(where map[string]interface{}) int {
	var ret, err = Count(u.db, "users", where)
	if err != nil {
		panic(err)
	}
	return ret
}

func (u *UsersModel) Add(data map[string]interface{}) *DMLResult {
	var ret, err = Add(u.db, "users", data)
	if err != nil {
		panic(err)
	}
	return ret
}

func (u *UsersModel) DelByOid(oid string) *DMLResult {
	var ret, err = DML(u.db, "users", "delete from goods where oid=?", oid)
	if err != nil {
		panic(err)
	}
	return ret
}

func (u *UsersModel) GetByOid(oid string) *User {
	var us = u.Find(map[string]interface{}{"oid": oid}, nil)
	if 0 < len(us) {
		return us[0]
	}
	return nil
}

func (u *UsersModel) Find(where, opt map[string]interface{}) []*User {
	var sql = "select * from users"
	var ret, err = Find(sql, where, opt, u.Query)
	if err != nil {
		panic(err)
	}
	return ret.([]*User)
}

type User struct {
	Id         int
	Oid        string
	UserName   string
	Phone      string
	Email      string
	Password   string
	Salt       string
	Address    string
	Created_at string
}

func (u *UsersModel) Query(sql string, params ...interface{}) (interface{}, error) {
	rows, err := u.db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret = make([]*User, 0)
	for rows.Next() {
		user := &User{}
		var err = rows.Scan(&user.Id, &user.Oid, &user.UserName, &user.Phone, &user.Email,
			&user.Password, &user.Salt, &user.Address, &user.Created_at)
		if err != nil {
			return nil, err
		}
		ret = append(ret, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (u *UsersModel) UpdateByOid(oid string, data map[string]interface{}) *DMLResult {
	var r = ExportKeyValues(data)
	var sql = MakeUpdateSql("users", r.Keys, "oid")
	var ret, err = DML(u.db, sql, append(r.Values, oid)...)
	if err != nil {
		panic(err)
	}
	return ret
}
