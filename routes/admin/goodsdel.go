package admin

import (
	"github.com/zengming00/go-testShop/framework"
)

func GoodsDel(ctx *framework.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	var oid = ctx.R.FormValue("oid")
	ctx.Goods.DelByOid(oid)
	ctx.Redirect("./goodslist.go")
}
