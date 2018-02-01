package routes

import (
	"image/png"

	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/lib/image"
)

func Capcha(ctx *lib.HandlerContext) {
	var r = image.MakeCapcha()

	ctx.SetSessionVal("__verify", r.Str)

	ctx.W.Header().Set("content-type", "image/png")
	err := png.Encode(ctx.W, r.Img)
	if err != nil {
		panic(err)
	}
}
