package user

import (
	"fmt"
	"regexp"

	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/models"
)

func Reg(ctx *framework.HandlerContext) {
	if ctx.R.Method == "GET" {
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(ctx.Cates.Find()),
			"user": ctx.GetUser(),
			"cart": ctx.GetCart(),
		}
		ctx.Render("./views/user/reg.html", data, lib.FuncMap)
		return
	}

	if ctx.R.Method == "POST" {
		var yzm = ctx.R.FormValue("yzm")
		if !ctx.Verify(yzm) {
			ctx.Json(map[string]string{"error": "验证码错误"})
			return
		}

		var username = ctx.R.FormValue("username")
		var email = ctx.R.FormValue("email")
		var password = ctx.R.FormValue("password")
		var repassword = ctx.R.FormValue("repassword")
		var phone = ctx.R.FormValue("phone")

		if !regexp.MustCompile(`^(13[0-9]|14[5|7]|15\d|18\d)\d{8}$`).Match([]byte(phone)) {
			ctx.Json(map[string]string{"error": "手机号无法接受"})
			return
		}

		if !regexp.MustCompile(`^[a-zA-Z0-9]{6,16}$`).Match([]byte(username)) || username == "admin" { // 直接拒绝admin账号
			ctx.Json(map[string]string{"error": "用户名只能是6-16位英文字母+数字"})
			return
		}
		if !regexp.MustCompile(`^[a-zA-z0-9]+@[a-zA-z0-9]+(\.[a-zA-z0-9]+)+$`).Match([]byte(email)) {
			ctx.Json(map[string]string{"error": "邮箱无法接受"})
			return
		}
		if !regexp.MustCompile(`^[a-zA-Z0-9]{6,16}$`).Match([]byte(password)) {
			ctx.Json(map[string]string{"error": "密码只能是6-16位英文字母+数字"})
			return
		} else if password != repassword {
			ctx.Json(map[string]string{"error": "两次密码不一致"})
			return
		}

		// 看用户名、手机、邮箱是否已存在
		var cond = map[string]interface{}{
			"$or": []map[string]interface{}{
				{"userName": username},
				{"phone": phone},
				{"email": email},
			},
		}
		var r = ctx.Users.Find(cond, nil)
		fmt.Println(r)
		if len(r) == 0 { // 未找到返回空数组，此时才可以入库
			var salt = lib.MakeSalt(6)
			var data = map[string]interface{}{
				"oid":      models.MakeOid(),
				"userName": username,
				"phone":    phone,
				"email":    email,
				"password": lib.EncodePassword(password, salt),
				"salt":     salt,
			}
			var r = ctx.Users.Add(data)
			fmt.Println(r)
			ctx.Json(map[string]string{"success": "注册成功！"})
		} else {
			ctx.Json(map[string]string{"error": "注册失败！请更换用户名、手机或邮箱"})
		}
	}

}
