package models

type GoodsStruct struct{}

var Goods GoodsStruct

func Count(where map[string]interface{}) {

}

func (*GoodsStruct) Query(sql string, params []interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret = make([]map[string]interface{}, 0)
	for rows.Next() {
		var id int
		var oid string
		var goods_name string
		var cat_id string
		var shop_price string
		var goods_img string
		var goods_desc string
		var goods_number int
		var is_best int
		var is_new int
		var is_hot int
		var is_on_sale int
		var created_at string
		var updated_at string

		var err = rows.Scan(&id, &oid, &goods_name, &cat_id, &shop_price, &goods_img, &goods_desc, &goods_number, &is_best, &is_new, &is_hot, &is_on_sale, &created_at, &updated_at)
		if err != nil {
			return nil, err
		}

		ret = append(ret, map[string]interface{}{
			"id":           id,
			"oid":          oid,
			"goods_name":   goods_name,
			"cat_id":       cat_id,
			"shop_price":   shop_price,
			"goods_img":    goods_img,
			"goods_desc":   goods_desc,
			"goods_number": goods_number,
			"is_best":      is_best,
			"is_new":       is_new,
			"is_hot":       is_hot,
			"is_on_sale":   is_on_sale,
			"created_at":   created_at,
			"updated_at":   updated_at,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}
