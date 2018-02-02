package admin

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/models"
)

func GoodsAdd(ctx *framework.HandlerContext) {
	var isAdmin = false
	if v, ok := ctx.GetSessionVal("isAdmin"); ok {
		isAdmin = v.(bool)
	}
	if !isAdmin {
		ctx.Redirect("/routes/admin/login.go")
	}

	if ctx.R.Method == "GET" {
		var allData = ctx.Cates.Find()
		var data = map[string]interface{}{
			"tree": ctx.Cates.GetTree(allData),
		}
		ctx.Render("./views/admin/goodsadd.html", data, lib.FuncMap)
		return
	}

	if ctx.R.Method == "POST" {
		var err = ctx.R.ParseMultipartForm(1024 * 1024)
		if err != nil {
			panic(err)
		}

		formFile, fh, err := ctx.R.FormFile("goods_img")
		if err == http.ErrMissingFile {
			ctx.Send("未上传图片")
			return
		}
		if err != nil {
			panic(err)
		}
		defer formFile.Close()

		var contentType = fh.Header.Get("Content-Type")
		if !strings.HasPrefix(contentType, "image") {
			ctx.Send("上传的不是图片")
			return
		}
		var ext = filepath.Ext(fh.Filename)
		var oidFilename = models.MakeOid() + ext
		var dateDir = lib.GetDateDir()
		// 文件的本地磁盘路径
		var filename = filepath.Join(dateDir.Fullpath, oidFilename)
		// web访问路径
		var publicPath = dateDir.Dir + oidFilename

		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.Copy(file, formFile)
		if err != nil {
			panic(err)
		}

		var data = map[string]interface{}{
			"oid":          models.MakeOid(),
			"goods_name":   ctx.R.FormValue("goods_name"),
			"cat_id":       ctx.R.FormValue("cat_id"),
			"shop_price":   ctx.R.FormValue("shop_price"),
			"goods_desc":   ctx.R.FormValue("goods_desc"),
			"goods_number": lib.ParsePositiveInt(ctx.R.FormValue("goods_number"), 0),
			"is_best":      lib.ParsePositiveInt(ctx.R.FormValue("is_best"), 0),
			"is_new":       lib.ParsePositiveInt(ctx.R.FormValue("is_new"), 0),
			"is_hot":       lib.ParsePositiveInt(ctx.R.FormValue("is_hot"), 0),
			"is_on_sale":   lib.ParsePositiveInt(ctx.R.FormValue("is_on_sale"), 0),
			"goods_img":    publicPath,
		}

		var result = ctx.Goods.Add(data)
		fmt.Printf("%#v\n", result)
		// ctx.Json(data)
		ctx.Redirect("./goodslist.go")
	}
}
