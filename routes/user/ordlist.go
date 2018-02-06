package user

import (
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func Ordlist(ctx *framework.HandlerContext) {
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
			"ords": ctx.Ordinfos.Find(
				map[string]interface{}{"userId": user.Oid},
				map[string]interface{}{
					"sort": map[string]interface{}{"id": "desc"},
				},
			),
		}
		ctx.Render("./views/user/ordlist.html", data, lib.FuncMap)
		return
	}
}
