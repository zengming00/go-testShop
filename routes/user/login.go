package user

import (
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func Login(ctx *framework.HandlerContext) {
	if ctx.R.Method == "GET" {
		var allData = ctx.Cates.Find()
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(allData),
			"user": ctx.GetUser(),
			"cart": ctx.GetCart(),
		}
		ctx.Render("./views/user/login.html", data, lib.FuncMap)
		return
	}

	if ctx.R.Method == "POST" {
		var yzm = ctx.R.FormValue("yzm")
		if !ctx.Verify(yzm) {
			ctx.Json(map[string]interface{}{"error": "验证码错误"})
			return
		}

		var username = ctx.R.FormValue("username")
		var password = ctx.R.FormValue("password")

		// 允许用户名、手机号、邮箱登录
		var cond = map[string]interface{}{
			"$or": []map[string]interface{}{
				{"userName": username},
				{"phone": username},
				{"email": username},
			},
		}
		var us = ctx.Users.Find(cond, nil)
		if len(us) > 0 {
			var doc = us[0]
			if doc.Password == lib.EncodePassword(password, doc.Salt) {
				ctx.SetUser(doc)
				ctx.Json(map[string]string{"success": "登录成功！"})
				return
			}
			ctx.Json(map[string]string{"error": "密码错误！"})
			return
		}
		ctx.Json(map[string]string{"error": "不存在的用户"})
		return
	}
}
