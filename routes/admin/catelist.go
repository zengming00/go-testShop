package admin

import "github.com/zengming00/go-testShop/lib"
import "github.com/zengming00/go-testShop/models"

func CateList(ctx *lib.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	var allData = models.Cates.Find()
	var data = map[string]interface{}{
		"tree": models.Cates.GetTree(allData),
	}
	ctx.Render("./views/admin/catelist.html", data, nil)
}
