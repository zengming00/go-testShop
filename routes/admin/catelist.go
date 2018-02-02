package admin

import (
	"github.com/zengming00/go-testShop/framework"
)

func CateList(ctx *framework.HandlerContext) {
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
	}
	ctx.Render("./views/admin/catelist.html", data, nil)
}
