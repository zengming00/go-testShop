package admin

import (
	"github.com/zengming00/go-testShop/lib"
)

func GoodsList(ctx *lib.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	if ctx.R.Method == "GET" {
		// var query = ctx.GetQuery()
		// var cat_id = query["cat_id"]
		// var sortTime = "D"
		// if v, ok := query["sortTime"]; ok {
		// 	sortTime = v
		// }
		// var tplData = map[string]interface{}{ //传递给模板的数据
		// 	"cat_id":   cat_id,
		// 	"sortTime": sortTime,
		// 	// toDateStr: Common.toDateStr,
		// }

		//   var total = 0;
		//   if (cat_id != "") {
		//     total = goods.count({ cat_id: cat_id });
		//   } else { //未传递栏目ID，查找所有
		//     total = goods.count();
		//   }
		//   var req = {
		//     originalUrl: request.getUri(),
		//     query: query,
		//   }
		//   var page = new Page(req, total, 8);
		//   var allCates = cates.find();
		//   tplData.tree = cates.getTree(allCates);
		//   tplData.page = page.show();

		//   sortTime = sortTime == 'D' ? 'desc' : 'asc';
		//   var opt = { sort: { id: sortTime }, skip: page.firstRow, limit: page.listRows };
		//   if (cat_id) {
		//     tplData.goods = goods.find({ cat_id: cat_id }, opt);
		//   } else {
		//     tplData.goods = goods.find({}, opt);
		//   }
		//   if (query.json) {
		//     res.json(tplData);
		//     return;
		//   }
		//   res.render('./views/admin/goodslist.ejs', tplData);
	}
}
