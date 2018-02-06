package user

import (
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func Favor(ctx *framework.HandlerContext) {
	var user = ctx.GetUser()
	if user == nil {
		ctx.Redirect("/routes/user/login.go")
		return
	}
	if ctx.R.Method == "GET" {
		var allData = ctx.Cates.Find()
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(allData),
			"user": user,
			"cart": ctx.GetCart(),
		}
		ctx.Render("./views/user/favor.html", data, lib.FuncMap)
		return
	}
}
