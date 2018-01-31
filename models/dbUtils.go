package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/zengming00/go-testShop/lib"
)

type DMLResult struct {
	LastInsertId int64
	RowsAffected int64
}

func DML(db *sql.DB, sql string, args []interface{}) (*DMLResult, error) {
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	ret := &DMLResult{}
	r, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	ret.LastInsertId = r

	r, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}
	ret.RowsAffected = r
	return ret, nil
}

// 生成len个?长度的?,?字符串
func MakePlaceStr(len int) string {
	arr := make([]string, len)
	for i := 0; i < len; i++ {
		arr[i] = "?"
	}
	return strings.Join(arr, ",")
}

// 生成xx=?,yy=?字符串
func MakeFieldPlaceStr(fields []string) string {
	length := len(fields)
	arr := make([]string, length)
	for i := 0; i < length; i++ {
		arr[i] = fields[i] + "=?"
	}
	return strings.Join(arr, ",")
}

func MakeInsertSql(table string, fields []string) string {
	return fmt.Sprintf("insert into %s(%s) values(%s)", table, strings.Join(fields, ","), MakePlaceStr(len(fields)))
}

func MakeUpdateSql(table string, fields []string, wheres []string) string {
	return fmt.Sprintf("update %s set %s where %s", table, MakeFieldPlaceStr(fields), MakeFieldPlaceStr(wheres))
}

func MakeOid() string {
	t := lib.CurrentTimeMillis()
	return fmt.Sprintf("%x", t)
}

type BuildAndResult struct {
	And  string
	Args []interface{}
}

func BuildAnd(fields []string, values []interface{}) *BuildAndResult {
	length := len(fields)
	args := make([]interface{}, 0)
	conds := make([]string, 0)
	for i := 0; i < length; i++ {
		k := fields[i]
		v := values[i]
		if m, ok := v.(map[string][]interface{}); ok {
			if ins, ok := m["$in"]; ok {
				conds = append(conds, k+" in ("+MakePlaceStr(len(ins))+")")
				args = append(args, ins...)
				continue
			}
			panic(errors.New("not support"))
		}
		conds = append(conds, k+"=?")
		args = append(args, v)
	}
	return &BuildAndResult{
		And:  strings.Join(conds, " and "),
		Args: args,
	}
}

type KeyValues struct {
	Keys   []string
	Values []interface{}
}

func ExportKeyValues(data map[string]interface{}) *KeyValues {
	kv := &KeyValues{
		Keys:   make([]string, 0, len(data)),
		Values: make([]interface{}, 0, len(data)),
	}
	for k, v := range data {
		kv.Keys = append(kv.Keys, k)
		kv.Values = append(kv.Values, v)
	}
	return kv
}

type BuildWhereResult struct {
	Where string
	Args  []interface{}
}

func BuildWhere(where map[string]interface{}) *BuildWhereResult {
	if vs, ok := where["$or"]; ok {
		if arr, ok := vs.([]map[string]interface{}); ok {
			length := len(arr)
			args := make([]interface{}, 0)
			conds := make([]string, 0)
			for i := 0; i < length; i++ {
				var r = ExportKeyValues(arr[i])
				var rr = BuildAnd(r.Keys, r.Values)
				conds = append(conds, "("+rr.And+")")
				args = append(args, rr.Args...)
			}
			return &BuildWhereResult{
				Where: " where " + strings.Join(conds, " or "),
				Args:  args,
			}
		}
		panic(errors.New("not support"))
	}
	var r = ExportKeyValues(where)
	var rr = BuildAnd(r.Keys, r.Values)
	return &BuildWhereResult{
		Where: " where " + rr.And,
		Args:  rr.Args,
	}
}

type QueryFunc = func(sql string, params []interface{}) (interface{}, error)

func Find(sql string, where, opt map[string]interface{}, queryFunc QueryFunc) (interface{}, error) {
	var params = make([]interface{}, 0)
	if where != nil {
		var r = BuildWhere(where)
		sql += r.Where
		params = append(params, r.Args...)
	}

	if sortv, ok := opt["sort"]; ok {
		if sort, ok := sortv.(map[string]interface{}); ok {
			var r = ExportKeyValues(sort)
			if v, ok := r.Values[0].(string); ok {
				sql += " order by " + r.Keys[0] + " " + v
			}
			// todo
		}
		// todo
	}
	var skip, skipOk = opt["skip"]
	var limit, limitOk = opt["limit"]
	if skipOk && limitOk {
		sql += " limit ?,?"
		params = append(params, skip)
		params = append(params, limit)
	}
	return queryFunc(sql, params)
}

func Add(db *sql.DB, table string, data map[string]interface{}) (*DMLResult, error) {
	var r = ExportKeyValues(data)
	var sql = MakeInsertSql(table, r.Keys)
	return DML(db, sql, r.Values)
}
