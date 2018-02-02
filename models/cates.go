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

func (c *CatesModel) GetTree(rows []*Cate) []*CateTreeItem {
	return c.getTree(rows, "0", 0)
}

func (c *CatesModel) getTree(rows []*Cate, pid string, level int) []*CateTreeItem {
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
			tree = append(tree, c.getTree(rows, row.Oid, level+1)...)
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

func (c *CatesModel) Find() []*Cate {
	r, err := c.Query("select * from cates")
	if err != nil {
		panic(err)
	}
	return r.([]*Cate)
}

type Cate struct {
	Id         int
	Oid        string
	Cat_name   string
	Intro      string
	Parent_id  string
	Created_at string
}

func (c *CatesModel) Query(sql string, params ...interface{}) (interface{}, error) {
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

func (c *CatesModel) Add(data map[string]interface{}) *DMLResult {
	var r = ExportKeyValues(data)
	var sql = MakeInsertSql("cates", r.Keys)
	var ret, err = DML(c.db, sql, r.Values...)
	if err != nil {
		panic(err)
	}
	return ret
}

func (c *CatesModel) DelByOid(oid string) *DMLResult {
	r, err := DML(c.db, "delete from cates where oid=?", oid)
	if err != nil {
		panic(err)
	}
	return r
}

func (c *CatesModel) GetByOid(oid string) []*Cate {
	r, err := c.Query("select * from cates where oid = ?", oid)
	if err != nil {
		panic(err)
	}
	return r.([]*Cate)
}

func (c *CatesModel) UpdateByOid(oid string, data map[string]interface{}) *DMLResult {
	var r = ExportKeyValues(data)
	var sql = MakeUpdateSql("cates", r.Keys, "oid")
	var ret, err = DML(c.db, sql, append(r.Values, oid)...)
	if err != nil {
		panic(err)
	}
	return ret
}
