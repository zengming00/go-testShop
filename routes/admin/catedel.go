package admin

import (
	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/models"
)

func CateDel(ctx *lib.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	var oid = ctx.R.FormValue("oid")
	models.Cates.DelByOid(oid)
	ctx.Redirect("./catelist.go")
}
