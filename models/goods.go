package models

import (
	"database/sql"
)

type GoodsModel struct {
	db *sql.DB
}

type Good struct {
	Id           int
	Oid          string
	Goods_name   string
	Cat_id       string
	Shop_price   string
	Goods_img    string
	Goods_desc   string
	Goods_number int
	Is_best      int
	Is_new       int
	Is_hot       int
	Is_on_sale   int
	Created_at   string
	Updated_at   string
}

func (g *GoodsModel) Query(sql string, params []interface{}) (interface{}, error) {
	rows, err := g.db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret = make([]*Good, 0)
	for rows.Next() {
		var good = &Good{}
		var err = rows.Scan(&good.Id, &good.Oid, &good.Goods_name, &good.Cat_id, &good.Shop_price, &good.Goods_img, &good.Goods_desc,
			&good.Goods_number, &good.Is_best, &good.Is_new, &good.Is_hot, &good.Is_on_sale, &good.Created_at, &good.Updated_at)
		if err != nil {
			return nil, err
		}
		ret = append(ret, good)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}
