package lib

import (
	"strings"
)

func Verify(yzm string, ctx *HandlerContext) bool {
	var v, ok = ctx.GetSessionVal("__verify")
	if ok {
		if v == strings.ToUpper(yzm) {
			//清空，防止多次使用
			ctx.SetSessionVal("__verify", nil)
			return true
		}
	}
	return false
}
