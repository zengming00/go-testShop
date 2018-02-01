package admin

import (
	"github.com/zengming00/go-testShop/lib"
)

func Index(ctx *lib.HandlerContext) {
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		if v.(bool) {
			ctx.HtmlFile("./views/admin/static/index.html")
			return
		}
	}
	ctx.Redirect("/routes/admin/login.go")
}
