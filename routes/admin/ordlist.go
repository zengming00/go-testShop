package admin

import (
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func OrdList(ctx *framework.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	var allData = ctx.Cates.Find()
	var data = map[string]interface{}{
		"tree": ctx.Cates.GetTree(allData),
		"ords": ctx.Ordinfos.Find(nil, map[string]interface{}{
			"sort": map[string]interface{}{"id": "desc"},
		}),
	}
	ctx.Render("./views/admin/ordlist.html", data, lib.FuncMap)
}
