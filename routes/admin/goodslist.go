package admin

import (
	"html/template"

	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func GoodsList(ctx *framework.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	if ctx.R.Method == "GET" {
		var query = ctx.GetQuery()
		var cat_id = query["cat_id"]
		var sortTime = "desc"
		if v, ok := query["sortTime"]; ok {
			if v != "desc" {
				sortTime = "asc"
			}
		}

		var where map[string]interface{}
		if cat_id != "" {
			where = map[string]interface{}{"cat_id": cat_id}
		}
		var total = ctx.Goods.Count(where)
		var page = lib.NewPage(total, 8, ctx.R.RequestURI, ctx.GetQuery())
		var allCates = ctx.Cates.Find()

		var opt = map[string]interface{}{
			"sort":  map[string]interface{}{"id": sortTime},
			"skip":  page.FirstRow,
			"limit": page.ListRows,
		}

		var tplData = map[string]interface{}{ //传递给模板的数据
			"cat_id":   cat_id,
			"sortTime": sortTime,
			"tree":     ctx.Cates.GetTree(allCates),
			"page":     template.HTML(page.Show()),
			"goods":    ctx.Goods.Find(where, opt),
		}

		ctx.Render("./views/admin/goodslist.html", tplData, template.FuncMap{
			"genList": func(v int) []int {
				return make([]int, v)
			},
			"toDateStr": func(s string) string {
				// todo 转成中国时间
				return s
			},
		})
	}
}
