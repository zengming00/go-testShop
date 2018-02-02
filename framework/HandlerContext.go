package framework

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"

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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.W.WriteHeader(http.StatusServiceUnavailable)
		}
	}()

	tplName := filepath.Base(filename)
	tpl, err := template.New(tplName).Funcs(funcMap).ParseFiles(filename)
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