package user

import (
	"github.com/zengming00/go-testShop/framework"
)

func Logout(ctx *framework.HandlerContext) {
	ctx.SetUser(nil)
	ctx.Redirect("./login.go")
}
