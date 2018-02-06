package user

import (
	"fmt"
	"regexp"

	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
)

func Repwd(ctx *framework.HandlerContext) {
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
		ctx.Render("./views/user/repwd.html", data, lib.FuncMap)
		return
	}
	if ctx.R.Method == "POST" {
		var oldpassword = ctx.R.FormValue("old_password")
		var password = ctx.R.FormValue("new_password")
		var repassword = ctx.R.FormValue("comfirm_password")

		if !regexp.MustCompile(`^[a-zA-Z0-9]{6,16}$`).Match([]byte(password)) {
			ctx.Json(map[string]string{"error": "密码只能是6-16位英文字母+数字"})
			return
		} else if password != repassword {
			ctx.Json(map[string]string{"error": "两次密码不一致"})
			return
		}

		if user.Password == lib.EncodePassword(oldpassword, user.Salt) {
			// 旧密码验证成功
			password = lib.EncodePassword(password, user.Salt)
			var r = ctx.Users.UpdateByOid(user.Oid, map[string]interface{}{"password": password})
			fmt.Printf("%#v\n", r)
			user.Password = password
			ctx.SetUser(user)
			ctx.Json(map[string]string{"success": "修改成功"})
			return
		}
		ctx.Json(map[string]string{"error": "原密码错误"})
	}
}
