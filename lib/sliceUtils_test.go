package lib

import (
	"reflect"
	"testing"
)

func TestUnshift(t *testing.T) {
	var a = []int{1, 2, 3}
	a = Unshift(a, 22).([]int)
	if !reflect.DeepEqual(a, []int{22, 1, 2, 3}) {
		t.Error("a:", a)
	}

	var b = []string{"a", "b"}
	b = Unshift(b, "haha").([]string)
	if !reflect.DeepEqual(b, []string{"haha", "a", "b"}) {
		t.Error("b:", b)
	}
}

func TestToInterfaceSlice(t *testing.T) {
	var a = ToInterfaceSlice([]int{12, 32, 0})
	if !reflect.DeepEqual(a, []interface{}{12, 32, 0}) {
		t.Error("a:", a)
	}
	var b = ToInterfaceSlice([]string{"a", "b"})
	if !reflect.DeepEqual(b, []interface{}{"a", "b"}) {
		t.Error("b:", b)
	}
}
