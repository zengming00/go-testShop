package admin

import (
	"github.com/zengming00/go-testShop/lib"
)

func Login(ctx *lib.HandlerContext) {
	if ctx.R.Method == "GET" {
		var filename = "./views/admin/login.html"
		ctx.Render(filename, nil)
		return
	}

	if ctx.R.Method == "POST" {
		var username = ctx.R.FormValue("username")
		var password = ctx.R.FormValue("password")
		var yzm = ctx.R.FormValue("yzm")

		var verifyOk = lib.Verify(yzm, ctx)
		if username == "admin" && password == "admin123" && verifyOk {
			ctx.SetSessionVal("isAdmin", true)
		}
		ctx.Redirect("./index.go")
	}
}
