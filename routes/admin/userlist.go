package admin

import (
	"github.com/zengming00/go-testShop/framework"
)

func UserList(ctx *framework.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	var us = ctx.Users.Find(nil, nil)
	var data = map[string]interface{}{
		"users": us,
	}
	ctx.Render("./views/admin/userlist.html", data, nil)
}
