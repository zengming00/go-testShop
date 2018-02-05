package lib

// 自定义模板函数
import (
	"bytes"
	"html/template"
	"math"
	"path/filepath"
	"strconv"
	"time"
)

var FuncMap template.FuncMap

func init() {
	FuncMap = template.FuncMap{
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
		"shopPrice": func(price string) string {
			var p = float64(ParsePositiveInt(price, 0))
			var i = p + math.Ceil(p*0.15)
			return strconv.Itoa(int(i))
		},
		"include": include,
	}
}

func include(file string, dot map[string]interface{}) template.HTML {
	var buffer = &bytes.Buffer{}
	tpl, err := template.New(filepath.Base(file)).Funcs(FuncMap).ParseFiles(file)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(buffer, dot)
	if err != nil {
		panic(err)
	}
	return template.HTML(buffer.String())
}
