package routes

import (
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func Goods(ctx *framework.HandlerContext) {

	if ctx.R.Method == "GET" {
		var oid = ctx.R.FormValue("oid")
		var gs = ctx.Goods.GetByOid(oid)
		if gs == nil {
			ctx.Send("goods not found")
			return
		}

		var cats = ctx.Cates.Find()

		//   //添加到历史记录
		//   var gsData = {
		//     oid: gs.oid,
		//     goods_img: gs.goods_img,
		//     goods_name: gs.goods_name,
		//     shop_price: gs.shop_price
		//   };
		//   var history = session.get('history') ? JSON.parse(session.get('history')) : [];
		//   var temp = [];
		//   for (var i = 0; i < history.length; i++) { //过滤重复数据
		//     if (history[i]._id != gsData._id) {
		//       temp.push(history[i]);
		//     }
		//   }
		//   temp.unshift(gsData);//插入到前面
		//   if (temp.length > 5) {
		//     temp.pop();
		//   }
		//   session.set('history', JSON.stringify(history));

		//渲染
		var data = map[string]interface{}{
			"gs":     gs,
			"tree":   ctx.Cates.GetTree(cats),
			"family": ctx.Cates.GetFamily(cats, gs.Cat_id),
			"user":   ctx.GetUser(),
			"cart":   ctx.GetCart(),
		}
		ctx.Render("./views/goods.html", data, lib.FuncMap)
	}
}
