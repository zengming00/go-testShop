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
		map[string]interface{}{
			"$in": []interface{}{1, 3, 5},
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
