package framework

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/zengming00/go-testShop/lib"
	"github.com/zengming00/go-testShop/models"
)

type HandlerContext struct {
	sessionID  string
	W          http.ResponseWriter
	R          *http.Request
	SessionMgr *lib.SessionMgr
	CacheMgr   *lib.CacheMgr
	Cates      *models.CatesModel
	Goods      *models.GoodsModel
	Users      *models.UsersModel
}

func (c *HandlerContext) GetQuery() map[string]string {
	var query = make(map[string]string)
	vs, err := url.ParseQuery(c.R.URL.RawQuery)
	if err == nil {
		var vv = map[string][]string(vs)
		for k, v := range vv {
			query[k] = v[0]
		}
	}
	return query
}

func (c *HandlerContext) Send(text string) {
	c.W.Write([]byte(text))
}

func (c *HandlerContext) Render(filename string, data map[string]interface{}, funcMap template.FuncMap) {
	tpl, err := template.New(filepath.Base(filename)).Funcs(funcMap).ParseFiles(filename)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(c.W, data)
	if err != nil {
		panic(err)
	}
}

func (c *HandlerContext) Json(data interface{}) {
	c.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	var bts, err = json.Marshal(data)
	if err != nil {
		panic(err)
	}
	c.W.Write(bts)
}

func (c *HandlerContext) Html(html []byte) {
	c.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.W.Write(html)
}

func (c *HandlerContext) HtmlFile(path string) {
	var content, ok = c.CacheMgr.Get(path)
	if !ok {
		var err error
		content, err = ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		c.CacheMgr.Set(path, content, -1)
	}
	c.Html(content.([]byte))
}

func (c *HandlerContext) Redirect(path string) {
	c.W.Header().Set("location", path)
	c.W.WriteHeader(302)
}

func (c *HandlerContext) StartSession() string {
	if c.sessionID == "" {
		c.sessionID = c.SessionMgr.StartSession(c.W, c.R)
	}
	return c.sessionID
}

func (c *HandlerContext) SetSessionVal(key string, value interface{}) {
	if c.sessionID == "" {
		c.StartSession()
	}
	c.SessionMgr.SetSessionVal(c.sessionID, key, value)
}

func (c *HandlerContext) GetSessionVal(key string) (value interface{}, ok bool) {
	if c.sessionID == "" {
		c.StartSession()
	}
	return c.SessionMgr.GetSessionVal(c.sessionID, key)
}

const CART_KEY = "__cart"

func (c *HandlerContext) GetCart() *lib.Cart {
	if v, ok := c.GetSessionVal(CART_KEY); ok {
		if cart, ok := v.(*lib.Cart); ok {
			return cart
		}
	}
	var cart = lib.NewCart()
	c.SetSessionVal(CART_KEY, cart)
	return cart
}

func (c *HandlerContext) SetCart(cart *lib.Cart) {
	if cart == nil {
		return
	}
	c.SetSessionVal(CART_KEY, cart)
}

const USER_KEY = "__user"

func (c *HandlerContext) GetUser() *models.User {
	if v, ok := c.GetSessionVal(USER_KEY); ok {
		if u, ok := v.(*models.User); ok {
			return u
		}
	}
	return nil
}

func (c *HandlerContext) SetUser(user *models.User) {
	if user == nil {
		return
	}
	c.SetSessionVal(USER_KEY, user)
}

const HISTORY_KEY = "__histroy"

func (c *HandlerContext) AddHistroy(gs *models.Good) {
	var history = c.GetHistroys()
	if history == nil {
		history = make([]*models.Good, 0, 6)
	}
	var temp = make([]*models.Good, 0, 6)
	temp = append(temp, gs)     // 插入到前面
	for _, v := range history { // 过滤重复数据
		if v.Id != gs.Id {
			temp = append(temp, v)
		}
	}
	// 只保留5个
	var llen = len(temp)
	if llen > 5 {
		llen = 5
	}
	c.SetSessionVal(HISTORY_KEY, temp[:llen])
}

func (c *HandlerContext) GetHistroys() []*models.Good {
	if v, ok := c.GetSessionVal(HISTORY_KEY); ok {
		if g, ok := v.([]*models.Good); ok {
			return g
		}
	}
	return nil
}

func (ctx *HandlerContext) Verify(yzm string) bool {
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
