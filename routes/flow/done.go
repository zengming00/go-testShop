package flow

import (
	"fmt"
	"html/template"
	"math/rand"
	"time"

	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/models"
)

func Done(ctx *framework.HandlerContext) {
	var user = ctx.GetUser()
	if user == nil {
		ctx.Redirect("/routes/user/login.go")
		return
	}
	if ctx.R.Method == "POST" {
		var cart = ctx.GetCart()

		var date = time.Now()
		var ordId = fmt.Sprintf("%.4d%.2d%.2d%.2d%.2d%.2d%d", date.Year(), int(date.Month()), date.Day(),
			date.Hour(), date.Minute(), date.Second(), rand.Uint32())

		var carts = cart.Items()
		if len(carts) == 0 {
			ctx.Send("购物车为空")
			return
		}
		for _, gs := range carts {
			if 0 < gs.Num {
				var ordgood = map[string]interface{}{
					"oid":       models.MakeOid(),
					"ordId":     ordId,
					"goodsId":   gs.Id,
					"goodsName": gs.Name,
					"price":     gs.Price,
					"num":       gs.Num,
				}
				var r = ctx.Goods.DecrGoodsNum(gs.Id, gs.Num)
				if r.RowsAffected == 0 {
					ctx.Send("下单失败！    " + gs.Name + "   商品不存在或库存不足！")
					return
				}
				ctx.Ordgoods.Add(ordgood)
			}
		}

		var money = cart.GetTotalMoney()
		var ordinfo = map[string]interface{}{
			"oid":      models.MakeOid(),
			"ordId":    ordId,
			"userId":   user.Oid,
			"userName": user.UserName,
			"address":  user.Address,
			"payType":  "RMB",
			"payState": false,
			"money":    money,
			"fuyan":    ctx.R.FormValue("fuyan"),
		}
		ctx.Ordinfos.Add(ordinfo)

		cart.Clear()
		ctx.SetCart(cart)

		var allData = ctx.Cates.Find()
		var data = map[string]interface{}{
			"tree":    ctx.Cates.GetTree(allData),
			"user":    user,
			"cart":    cart,
			"ordId":   ordId,
			"money":   money,
			"payForm": template.HTML(lib.GetPayForm(ordId, money)),
		}
		ctx.Render("./views/flow/done.html", data, lib.FuncMap)
	}
}
