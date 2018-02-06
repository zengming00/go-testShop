package lib

import "reflect"

// data := []string{"A", "B", "C", "D"}
// data = append([]string{"Prepend Item"}, data...)

// 实现类似javascript中array的unshift功能
func Unshift(slice, v interface{}) interface{} {
	var typ = reflect.TypeOf(slice)
	if typ.Kind() == reflect.Slice {
		var vv = reflect.ValueOf(slice)
		var tmp = reflect.MakeSlice(typ, vv.Len()+1, vv.Cap()+1)
		tmp.Index(0).Set(reflect.ValueOf(v))
		var dst = tmp.Slice(1, tmp.Len())
		reflect.Copy(dst, vv)
		return tmp.Interface()
	}
	panic("not a slice")
}

func ToInterfaceSlice(slice interface{}) []interface{} {
	var typ = reflect.TypeOf(slice)
	if typ.Kind() == reflect.Slice {
		var vv = reflect.ValueOf(slice)
		var tmp = make([]interface{}, vv.Len())
		for i := range tmp {
			tmp[i] = vv.Index(i).Interface()
		}
		return tmp
	}
	panic("not a slice")
}
