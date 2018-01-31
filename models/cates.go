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

// func GetTree($rows, $pid, $level) {
//   $pid = $pid || 0;
//   $level = $level || 0;

//   var $tree = [];
//   for (var k in $rows) {
//     var $row = $rows[k];
//     if ($row.parent_id == $pid) {
//       $tree.push({  //$row是模型实体对象无法添加level属性
//         oid: $row.oid,
//         cat_name: $row.cat_name,
//         intro: $row.intro,
//         parent_id: $row.parent_id,
//         level: $level
//       });
//       //if($row.oid != $row.parent_id){ //原本以为会死循环，但实际上不会出现
//       $tree = $tree.concat(getTree($rows, $row.oid, $level + 1));
//       //}
//     }
//   }
//   return $tree;
// }

// exports.getFamily = function getFamily($rows, $catid) {
//   var $arr = [], k, row, isFind;
//   while ($catid != 0) {
//     isFind = false;
//     for (k in $rows) {
//       row = $rows[k];
//       if (row.oid == $catid) {
//         $arr.unshift(row);
//         $catid = row.parent_id;
//         isFind = true;//避免死循环
//         break;
//       }
//     }
//     if (!isFind) break;
//   }
//   return $arr;
// }

// exports.getChildCates = function getChilds($rows, $catid) {
//   var arr = [], k, r;
//   for (k in $rows) {
//     r = $rows[k];
//     if (r.parent_id == $catid) {
//       arr.push(r.oid);
//       arr = arr.concat(getChilds($rows, r.oid));
//     }
//   }
//   return arr;
// }

func (c *CatesModel) Find() ([]map[string]interface{}, error) {
	return c.Query("select * from cates", []interface{}{})
}

func (c *CatesModel) Query(sql string, params []interface{}) ([]map[string]interface{}, error) {
	rows, err := c.db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret = make([]map[string]interface{}, 0)
	for rows.Next() {
		var id int
		var oid string
		var cat_name string
		var intro string
		var parent_id string
		var created_at string

		var err = rows.Scan(&id, &oid, &cat_name, &intro, &parent_id, &created_at)
		if err != nil {
			return nil, err
		}

		ret = append(ret, map[string]interface{}{
			"id":         id,
			"oid":        oid,
			"cat_name":   cat_name,
			"intro":      intro,
			"parent_id":  parent_id,
			"created_at": created_at,
		})
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

func (c *CatesModel) GetByOid(oid string) ([]map[string]interface{}, error) {
	return c.Query("select * from cates where oid = ?", []interface{}{oid})
}

// exports.updateByOid = function updateByOid(oid, data) {
//   var r = dbUtils.exportKeyValues(data)
//   var sql = dbUtils.makeUpdateSql('cates', r.keys, ['oid'])
//   // 添加where条件值
//   r.values.push(oid)
//   return dbUtils.update(db, sql, r.values)
// }
