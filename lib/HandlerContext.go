package lib

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
)

type HandlerContext struct {
	W          http.ResponseWriter
	R          *http.Request
	SessionMgr *SessionMgr
	CacheMgr   *CacheMgr
	sessionID  string
}

func (c *HandlerContext) GetQuery() (map[string]string, error) {
	vs, err := url.ParseQuery(c.R.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	var vv = map[string][]string(vs)
	var query = make(map[string]string)

	for k, v := range vv {
		query[k] = v[0]
	}
	return query, nil
}

func (c *HandlerContext) Render(filename string, data map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.W.WriteHeader(http.StatusServiceUnavailable)
		}
	}()

	tplName := filepath.Base(filename)
	tpl, err := template.New(tplName).ParseFiles(filename)
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
