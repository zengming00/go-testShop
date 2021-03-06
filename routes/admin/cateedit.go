package admin

import (
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func CateEdit(ctx *framework.HandlerContext) {
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
		var allData = ctx.Cates.Find()
		var cate = ctx.Cates.GetByOid(oid)
		if len(cate) == 0 {
			ctx.Send("cate not found")
			return
		}
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(allData),
			"cate": cate[0],
		}
		ctx.Render("./views/admin/cateedit.html", data, lib.FuncMap)
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
		ctx.Cates.UpdateByOid(cat_id, data)
		ctx.Redirect("./catelist.go")
	}
}
