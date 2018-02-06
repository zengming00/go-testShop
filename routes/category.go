package routes

import (
	"html/template"

	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func Category(ctx *framework.HandlerContext) {
	if ctx.R.Method == "GET" {
		var catid = ctx.R.FormValue("cat_id")

		var allCates = ctx.Cates.Find()
		var childs = ctx.Cates.GetChildCates(allCates, catid)
		childs = append([]string{catid}, childs...) // 将当前栏目也包含在内

		var gsCond = map[string]interface{}{
			"cat_id": map[string]interface{}{
				"$in": childs,
			},
		}
		var total = ctx.Goods.Count(gsCond)

		var page = lib.NewPage(total, 9, ctx.R.RequestURI, ctx.GetQuery())

		var docs = ctx.Goods.Find(gsCond, map[string]interface{}{
			"skip":  page.FirstRow,
			"limit": page.ListRows,
		})
		var data = map[string]interface{}{
			"gs":      docs,
			"user":    ctx.GetUser(),
			"cart":    ctx.GetCart(),
			"history": ctx.GetHistroys(),
			"page":    template.HTML(page.Show()),
			"tree":    ctx.Cates.GetTree(allCates),
			"family":  ctx.Cates.GetFamily(allCates, catid),
		}
		ctx.Render("./views/category.html", data, lib.FuncMap)
	}

}
