package main

import (
	"log"
	"net/http"
	"runtime/debug"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zengming00/go-testShop/framework"
	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/models"
	"github.com/zengming00/go-testShop/routes"
	"github.com/zengming00/go-testShop/routes/admin"
	"github.com/zengming00/go-testShop/routes/flow"
	"github.com/zengming00/go-testShop/routes/user"
)

var cates *models.CatesModel
var goods *models.GoodsModel
var users *models.UsersModel
var ordinfos *models.OrdinfosModel
var ordgoods *models.OrdgoodsModel

var sessionMgr = lib.NewSessionMgr("sid", 60*15)
var cacheMgr = lib.NewCacheMgr(60)

type HandleFunc = func(resp http.ResponseWriter, req *http.Request)

func mHandle(h func(ctx *framework.HandlerContext)) HandleFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var ctx = &framework.HandlerContext{
			W:          resp,
			R:          req,
			SessionMgr: sessionMgr,
			CacheMgr:   cacheMgr,
			Cates:      cates,
			Goods:      goods,
			Users:      users,
			Ordinfos:   ordinfos,
			Ordgoods:   ordgoods,
		}

		defer func() {
			if err := recover(); err != nil {
				ctx.Json(map[string]string{"error": "服务器忙"})
				log.Println(err, string(debug.Stack()))
			}
		}()
		h(ctx)
	}
}

func main() {
	var db = models.OpenDB()
	cates = models.NewCatesModel(db)
	goods = models.NewGoodsModel(db)
	users = models.NewUsersModel(db)
	ordinfos = models.NewOrdinfosModel(db)
	ordgoods = models.NewOrdgoodsModel(db)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/routes/admin/goodsdel.go", mHandle(admin.GoodsDel))
	http.HandleFunc("/routes/admin/goodsadd.go", mHandle(admin.GoodsAdd))
	http.HandleFunc("/routes/admin/goodslist.go", mHandle(admin.GoodsList))
	http.HandleFunc("/routes/admin/login.go", mHandle(admin.Login))
	http.HandleFunc("/routes/admin/ordlist.go", mHandle(admin.OrdList))
	http.HandleFunc("/routes/admin/userlist.go", mHandle(admin.UserList))
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

	http.HandleFunc("/routes/user/ordlist.go", mHandle(user.Ordlist))
	http.HandleFunc("/routes/user/address.go", mHandle(user.Address))
	http.HandleFunc("/routes/user/liuyan.go", mHandle(user.Liuyan))
	http.HandleFunc("/routes/user/favor.go", mHandle(user.Favor))
	http.HandleFunc("/routes/user/info.go", mHandle(user.Info))
	http.HandleFunc("/routes/user/reg.go", mHandle(user.Reg))
	http.HandleFunc("/routes/user/repwd.go", mHandle(user.Repwd))
	http.HandleFunc("/routes/user/login.go", mHandle(user.Login))
	http.HandleFunc("/routes/user/logout.go", mHandle(user.Logout))

	http.HandleFunc("/routes/flow/checkout.go", mHandle(flow.Checkout))
	http.HandleFunc("/routes/flow/cartApi.go", mHandle(flow.CartApi))
	http.HandleFunc("/routes/flow/cart.go", mHandle(flow.Cart))
	http.HandleFunc("/routes/flow/done.go", mHandle(flow.Done))

	http.HandleFunc("/routes/category.go", mHandle(routes.Category))
	http.HandleFunc("/routes/goods.go", mHandle(routes.Goods))
	http.HandleFunc("/routes/capcha.go", mHandle(routes.Capcha))
	http.HandleFunc("/favicon.ico", func(resp http.ResponseWriter, req *http.Request) {
		// todo
		resp.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/", mHandle(routes.Index))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
