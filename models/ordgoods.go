package models

import (
	"database/sql"
)

type OrdgoodsModel struct {
	db *sql.DB
}

func NewOrdgoodsModel(db *sql.DB) *OrdgoodsModel {
	return &OrdgoodsModel{db}
}

func (o *OrdgoodsModel) Add(data map[string]interface{}) *DMLResult {
	var ret, err = Add(o.db, "ordgoods", data)
	if err != nil {
		panic(err)
	}
	return ret
}

type Ordgood struct {
	Id         int
	Oid        string
	OrdId      string
	GoodsId    string
	GoodsName  string
	Price      string
	Num        int
	Created_at string
}

func (o *OrdgoodsModel) Query(sql string, params ...interface{}) (interface{}, error) {
	rows, err := o.db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret = make([]*Ordgood, 0)
	for rows.Next() {
		var ordgood = &Ordgood{}
		var err = rows.Scan(&ordgood.Id, &ordgood.Oid, &ordgood.OrdId, &ordgood.GoodsId,
			&ordgood.GoodsName, &ordgood.Price, &ordgood.Num, &ordgood.Created_at)
		if err != nil {
			return nil, err
		}
		ret = append(ret, ordgood)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}
