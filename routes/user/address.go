package user

import (
	"fmt"
	"strings"

	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func Address(ctx *framework.HandlerContext) {
	var user = ctx.GetUser()
	if user == nil {
		ctx.Redirect("/routes/user/login.go")
		return
	}
	if ctx.R.Method == "GET" {
		var allData = ctx.Cates.Find()
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(allData),
			"user": user,
			"cart": ctx.GetCart(),
		}
		ctx.Render("./views/user/address.html", data, lib.FuncMap)
		return
	}
	if ctx.R.Method == "POST" {
		var address = ctx.R.FormValue("address")
		var name = ctx.R.FormValue("name")
		var phone = ctx.R.FormValue("phone")
		var addr = strings.TrimSpace(address + "  " + name + "  " + phone)

		var r = ctx.Users.UpdateByOid(user.Oid, map[string]interface{}{"address": addr})
		fmt.Printf("%#v\n", r)
		user.Address = addr
		ctx.SetUser(user)
		ctx.Redirect("./address.go")
	}
}
