package admin

import "github.com/zengming00/go-testShop/lib"

func Logout(ctx *lib.HandlerContext) {
	ctx.SetSessionVal("isAdmin", false)
	ctx.Redirect("./index.go")
}