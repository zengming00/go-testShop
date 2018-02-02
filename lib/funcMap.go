package lib

// 自定义模板函数
import (
	"html/template"
	"time"
)

var FuncMap = template.FuncMap{
	"genList": func(v int) []int {
		return make([]int, v)
	},
	"toDateStr": func(s string) string {
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			panic(err)
		}
		var cst = time.FixedZone("CST", 28800)
		var t2 = t.In(cst)
		return t2.Format("2006-01-02 15:04:05")
	},
}
