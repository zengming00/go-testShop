package flow

import (
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func Cart(ctx *framework.HandlerContext) {
	if ctx.R.Method == "GET" {
		var cats = ctx.Cates.Find()
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(cats),
			"user": ctx.GetUser(),
			"cart": ctx.GetCart(),
		}
		ctx.Render("./views/flow/cart.html", data, lib.FuncMap)
	}
}
