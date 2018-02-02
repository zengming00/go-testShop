package admin

import "github.com/zengming00/go-testShop/framework"

func Logout(ctx *framework.HandlerContext) {
	// ctx.SetSessionVal("isAdmin", false)
	ctx.SetSessionVal("isAdmin", true)
	ctx.Redirect("./index.go")
}
