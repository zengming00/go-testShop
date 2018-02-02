package admin

import (
	"strings"

	"github.com/zengming00/go-testShop/framework"
)

func verify(yzm string, ctx *framework.HandlerContext) bool {
	var v, ok = ctx.GetSessionVal("__verify")
	if ok {
		if v == strings.ToUpper(yzm) {
			//清空，防止多次使用
			ctx.SetSessionVal("__verify", nil)
			return true
		}
	}
	return false
}

func Login(ctx *framework.HandlerContext) {
	if ctx.R.Method == "GET" {
		var filename = "./views/admin/login.html"
		ctx.Render(filename, nil, nil)
		return
	}

	if ctx.R.Method == "POST" {
		var username = ctx.R.FormValue("username")
		var password = ctx.R.FormValue("password")
		var yzm = ctx.R.FormValue("yzm")

		var verifyOk = verify(yzm, ctx)
		if username == "admin" && password == "admin123" && verifyOk {
			ctx.SetSessionVal("isAdmin", true)
		}
		ctx.Redirect("./index.go")
	}
}
