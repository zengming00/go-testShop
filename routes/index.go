package routes

import (
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func Index(ctx *framework.HandlerContext) {
	if ctx.R.URL.Path != "/" {
		return
	}
	if ctx.R.Method == "GET" {
		var data = map[string]interface{}{
			"history": ctx.GetHistroys(),

			"tree": ctx.Cates.GetTree(ctx.Cates.Find()),
			"bestGs": ctx.Goods.Find(
				map[string]interface{}{
					"is_best":    1,
					"is_on_sale": 1,
				},
				map[string]interface{}{
					"sort":  map[string]interface{}{"id": "desc"},
					"skip":  0,
					"limit": 3,
				}),
			"newGs": ctx.Goods.Find(
				map[string]interface{}{
					"is_new":     1,
					"is_on_sale": 1,
				},
				map[string]interface{}{
					"sort":  map[string]interface{}{"id": "desc"},
					"skip":  0,
					"limit": 3,
				}),
			"hotGs": ctx.Goods.Find(
				map[string]interface{}{
					"is_hot":     1,
					"is_on_sale": 1,
				},
				map[string]interface{}{
					"sort":  map[string]interface{}{"id": "desc"},
					"skip":  0,
					"limit": 3,
				}),

			"user": ctx.GetUser(),
			"cart": ctx.GetCart(),
		}
		ctx.Render("./views/index.html", data, lib.FuncMap)
	}
}
