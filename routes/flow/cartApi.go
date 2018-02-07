package flow

import (
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func CartApi(ctx *framework.HandlerContext) {
	if ctx.R.Method == "GET" {
		var cart = ctx.GetCart()
		var a = ctx.R.FormValue("a")
		var oid = ctx.R.FormValue("oid")
		switch a {
		case "add":
			var gs = ctx.Goods.GetByOid(oid)
			cart.Add(gs.Oid, gs.Goods_name, gs.Goods_img, lib.ParsePositiveInt(gs.Shop_price, 0), 1)
			break
		case "del":
			cart.Del(oid)
			break
		case "incr":
			cart.Incr(oid)
			break
		case "decr":
			cart.Decr(oid)
			break
		default:
		}
		ctx.SetCart(cart)
		ctx.Redirect("./cart.go")
	}
}
