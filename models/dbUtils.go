package models

import (
	"database/sql"
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
		if m, ok := v.(map[string]interface{}); ok {
			if ins, ok := m["$in"]; ok {
				if arr, ok := ins.([]interface{}); ok {
					conds = append(conds, k+" in ("+MakePlaceStr(len(arr))+")")
					args = append(args, arr...)
				}
			}
			continue
		}
		conds = append(conds, k+"=?")
		args = append(args, v)
	}
	return &BuildAndResult{
		And:  strings.Join(conds, " and "),
		Args: args,
	}
}
