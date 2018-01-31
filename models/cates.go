package models

import (
	"database/sql"
)

type CatesModel struct {
	db *sql.DB
}

func NewCatesModel(db *sql.DB) *CatesModel {
	return &CatesModel{db}
}

type CateTreeItem struct {
	Oid       string
	Cat_name  string
	Intro     string
	Parent_id string
	Level     int
}

func (c *CatesModel) GetTree(rows []*Cate, pid string, level int) []*CateTreeItem {
	var tree = make([]*CateTreeItem, 0)
	for _, row := range rows {
		if row.Parent_id == pid {
			tree = append(tree, &CateTreeItem{
				Oid:       row.Oid,
				Cat_name:  row.Cat_name,
				Intro:     row.Intro,
				Parent_id: row.Parent_id,
				Level:     level,
			})
			tree = append(tree, c.GetTree(rows, row.Oid, level+1)...)
		}
	}
	return tree
}

func (c *CatesModel) GetFamily(rows []*Cate, catid string) []*Cate {
	var arr = make([]*Cate, 0)
	var isFind bool
	for catid != "0" {
		isFind = false
		for _, row := range rows {
			if row.Oid == catid {
				var tmp = make([]*Cate, len(arr)+1)
				tmp[0] = row
				copy(tmp[1:], arr)
				arr = tmp
				catid = row.Parent_id
				isFind = true //避免死循环
				break
			}
		}
		if !isFind {
			break
		}
	}
	return arr
}

func (c *CatesModel) GetChildCates(rows []*Cate, catid string) []string {
	var arr = make([]string, 0)
	for _, r := range rows {
		if r.Parent_id == catid {
			arr = append(arr, r.Oid)
			arr = append(arr, c.GetChildCates(rows, r.Oid)...)
		}
	}
	return arr
}

func (c *CatesModel) Find() ([]*Cate, error) {
	r, err := c.Query("select * from cates", []interface{}{})
	if err != nil {
		return nil, err
	}
	return r.([]*Cate), nil
}

type Cate struct {
	Id         int
	Oid        string
	Cat_name   string
	Intro      string
	Parent_id  string
	Created_at string
}

func (c *CatesModel) Query(sql string, params []interface{}) (interface{}, error) {
	rows, err := c.db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret = make([]*Cate, 0)
	for rows.Next() {
		cate := &Cate{}
		var err = rows.Scan(&cate.Id, &cate.Oid, &cate.Cat_name, &cate.Intro, &cate.Parent_id, &cate.Created_at)
		if err != nil {
			return nil, err
		}
		ret = append(ret, cate)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// exports.add = function add(data) {
//   var r = dbUtils.exportKeyValues(data);
//   var sql = dbUtils.makeInsertSql('cates', r.keys);
//   return dbUtils.insert(db, sql, r.values);
// }

// exports.delByOid = function delByOid(oid) {
//   return dbUtils.delete(db, 'delete from cates where oid=?', [oid]);
// }

func (c *CatesModel) GetByOid(oid string) ([]*Cate, error) {
	r, err := c.Query("select * from cates where oid = ?", []interface{}{oid})
	if err != nil {
		return nil, err
	}
	return r.([]*Cate), nil
}

// exports.updateByOid = function updateByOid(oid, data) {
//   var r = dbUtils.exportKeyValues(data)
//   var sql = dbUtils.makeUpdateSql('cates', r.keys, ['oid'])
//   // 添加where条件值
//   r.values.push(oid)
//   return dbUtils.update(db, sql, r.values)
// }
