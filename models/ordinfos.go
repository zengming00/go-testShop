package models

import (
	"database/sql"
)

type OrdinfosModel struct {
	db *sql.DB
}

func NewOrdinfosModel(db *sql.DB) *OrdinfosModel {
	return &OrdinfosModel{db}
}

func (o *OrdinfosModel) Add(data map[string]interface{}) *DMLResult {
	var ret, err = Add(o.db, "ordinfos", data)
	if err != nil {
		panic(err)
	}
	return ret
}

func (o *OrdinfosModel) GetByOid(oid string) *Ordinfo {
	var v = o.Find(map[string]interface{}{"oid": oid}, nil)
	if 0 < len(v) {
		return v[0]
	}
	return nil
}

func (o *OrdinfosModel) Find(where, opt map[string]interface{}) []*Ordinfo {
	var sql = "select * from ordinfos"
	r, err := Find(sql, where, opt, o.Query)
	if err != nil {
		panic(err)
	}
	return r.([]*Ordinfo)
}

type Ordinfo struct {
	Id         int
	Oid        string
	OrdId      string
	UserId     string
	UserName   string
	Address    string
	PayType    string
	PayState   int
	Money      string
	Fuyan      string
	Created_at string
}

func (o *OrdinfosModel) Query(sql string, params ...interface{}) (interface{}, error) {
	rows, err := o.db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret = make([]*Ordinfo, 0)
	for rows.Next() {
		var ord = &Ordinfo{}
		var err = rows.Scan(&ord.Id, &ord.Oid, &ord.OrdId, &ord.UserId, &ord.UserName,
			&ord.Address, &ord.PayType, &ord.PayState, &ord.Money, &ord.Fuyan, &ord.Created_at)
		if err != nil {
			return nil, err
		}
		ret = append(ret, ord)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}
