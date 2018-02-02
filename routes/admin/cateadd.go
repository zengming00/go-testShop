package admin

import (
	"html/template"

	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/models"
)

func CateAdd(ctx *lib.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	if ctx.R.Method == "GET" {
		var allData = models.Cates.Find()
		var data = map[string]interface{}{
			"tree": models.Cates.GetTree(allData),
		}
		ctx.Render("./views/admin/cateadd.html", data, template.FuncMap{
			"genList": func(v int) []int {
				return make([]int, v)
			},
		})
		return
	}

	if ctx.R.Method == "POST" {
		var data = map[string]interface{}{
			"cat_name":  ctx.R.FormValue("cat_name"),
			"intro":     ctx.R.FormValue("cat_desc"),
			"parent_id": ctx.R.FormValue("parent_id"),
			"oid":       models.MakeOid(),
		}
		models.Cates.Add(data)
		ctx.Redirect("./catelist.go")
	}
}
