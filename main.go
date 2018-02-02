package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/routes"
	"github.com/zengming00/go-testShop/routes/admin"
)

var sessionMgr = lib.NewSessionMgr("sid", 60*15)
var cacheMgr = lib.NewCacheMgr(60)

type HandleFunc = func(resp http.ResponseWriter, req *http.Request)

func mHandle(h func(ctx *lib.HandlerContext)) HandleFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		h(&lib.HandlerContext{
			W:          resp,
			R:          req,
			SessionMgr: sessionMgr,
			CacheMgr:   cacheMgr,
		})
	}
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/routes/admin/login.go", mHandle(admin.Login))
	http.HandleFunc("/routes/admin/catelist.go", mHandle(admin.CateList))
	http.HandleFunc("/routes/admin/cateadd.go", mHandle(admin.CateAdd))
	http.HandleFunc("/routes/admin/catedel.go", mHandle(admin.CateDel))
	http.HandleFunc("/routes/admin/cateedit.go", mHandle(admin.CateEdit))
	http.HandleFunc("/routes/admin/logout.go", mHandle(admin.Logout))
	http.HandleFunc("/routes/admin/index.go", mHandle(admin.Index))
	http.HandleFunc("/routes/admin/left.go", mHandle(admin.Left))
	http.HandleFunc("/routes/admin/main.go", mHandle(admin.Main))
	http.HandleFunc("/routes/admin/drag.go", mHandle(admin.Drag))
	http.HandleFunc("/routes/admin/top.go", mHandle(admin.Top))

	http.HandleFunc("/routes/capcha.go", mHandle(routes.Capcha))
	http.HandleFunc("/favicon.ico", func(resp http.ResponseWriter, req *http.Request) {
		// todo
	})

	http.HandleFunc("/test", mHandle(func(ctx *lib.HandlerContext) {
		if ctx.R.RequestURI == "/favicon.ico" {
			return
		}
		query := ctx.GetQuery()

		pg := lib.NewPage(10, 8, ctx.R.RequestURI, query)
		fmt.Println(pg.FirstRow, pg.ListRows)
		str := pg.Show()
		ctx.W.Write([]byte(str))
	}))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
