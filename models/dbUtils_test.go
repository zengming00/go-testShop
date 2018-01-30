package models

import (
	"reflect"
	"testing"
)

func TestMakePlaceStr(t *testing.T) {
	str := MakePlaceStr(3)
	if str != "?,?,?" {
		t.Fail()
	}
}

func TestMakeFieldPlaceStr(t *testing.T) {
	fields := []string{"hello", "world"}
	str := MakeFieldPlaceStr(fields)
	if str != "hello=?,world=?" {
		t.Error(str)
	}
}

func TestMakeOid(t *testing.T) {
	str := MakeOid()
	t.Log(str)
}

func TestBuildAnd(t *testing.T) {
	fields := []string{"a", "b"}
	values := []interface{}{
		"haha",
		map[string][]interface{}{
			"$in": {1, 3, 5},
		},
	}
	r := BuildAnd(fields, values)
	if r.And != "a=? and b in (?,?,?)" {
		t.Error(r.And)
	}
	if !reflect.DeepEqual(r.Args, []interface{}{"haha", 1, 3, 5}) {
		t.Error("args err")
	}
}

func TestBuildWhere(t *testing.T) {
	where := map[string]interface{}{
		"a": "haha",
		"b": map[string][]interface{}{
			"$in": {1, 3, 5},
		},
	}
	r := BuildWhere(where)
	if r.Where != " where a=? and b in (?,?,?)" {
		t.Error(r.Where)
	}
	if !reflect.DeepEqual(r.Args, []interface{}{"haha", 1, 3, 5}) {
		t.Error("args err")
	}
	where = map[string]interface{}{
		"$or": []map[string]interface{}{
			{"a": "haha"},
			{
				"b": map[string][]interface{}{
					"$in": {1, 3, 5},
				},
			},
		},
	}
	r = BuildWhere(where)
	if r.Where != " where (a=?) or (b in (?,?,?))" {
		t.Error(r.Where)
	}
	if !reflect.DeepEqual(r.Args, []interface{}{"haha", 1, 3, 5}) {
		t.Error("args err")
	}
}
