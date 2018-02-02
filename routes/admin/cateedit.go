package admin

import (
	"fmt"
	"html/template"

	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/models"
)

func CateEdit(ctx *lib.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	if ctx.R.Method == "GET" {
		var oid = ctx.R.FormValue("oid")
		if oid == "" {
			ctx.Send("oid is null")
			return
		}
		var allData = models.Cates.Find()
		var cate = models.Cates.GetByOid(oid)
		if len(cate) == 0 {
			ctx.Send("cate not found")
			return
		}
		var data = map[string]interface{}{
			"tree": models.Cates.GetTree(allData),
			"cate": cate[0],
		}
		fmt.Printf("%#v \n", data["cate"])
		ctx.Render("./views/admin/cateedit.html", data, template.FuncMap{
			"genList": func(v int) []int {
				return make([]int, v)
			},
		})
		return
	}

	if ctx.R.Method == "POST" {
		var cat_id = ctx.R.FormValue("cat_id")
		var parent_id = ctx.R.FormValue("parent_id")

		var data = map[string]interface{}{
			"cat_name":  ctx.R.FormValue("cat_name"),
			"intro":     ctx.R.FormValue("cat_desc"),
			"parent_id": parent_id,
		}
		if cat_id == parent_id {
			ctx.Send("错误的上级栏目")
			return
		}
		models.Cates.UpdateByOid(cat_id, data)
		ctx.Redirect("./catelist.go")
	}
}
