package lib

import (
	"math"
	"net/url"
	"strconv"
	"strings"
)

func strReplace(search []string, replace []string, str string) string {
	if len(search) != len(replace) {
		panic("len(search) != len(replace)")
	}
	var oldnew = make([]string, len(search)+len(replace))
	for i, v := range search {
		oldnew[i*2] = v
		oldnew[i*2+1] = replace[i]
	}
	r := strings.NewReplacer(oldnew...)
	return r.Replace(str)
}

type PageConfig struct {
	Header string
	Prev   string
	Next   string
	First  string
	Last   string
	Theme  string
}

type Page struct {
	pageTag    string
	p          string
	uri        string
	lastSuffix bool
	rollPage   int
	totalRows  int
	nowPage    int
	ListRows   int
	FirstRow   int
	Config     PageConfig
}

func NewPage(totalRows, listRows int, uri string, query map[string]string) *Page {
	var p = &Page{
		pageTag:    "PAGE",    // url参数替换标记
		p:          "p",       // 分页参数名
		uri:        uri,       // 当前链接URL
		rollPage:   11,        // 分页栏每页显示的页数
		lastSuffix: true,      // 最后一页是否显示总页数
		totalRows:  totalRows, // 设置总记录数
		Config: PageConfig{ // 分页显示定制
			Header: `<span class="rows">共 %TOTAL_ROW% 条记录</span>`,
			Prev:   "<<",
			Next:   ">>",
			First:  "1...",
			Last:   "...%TOTAL_PAGE%",
			Theme:  "%FIRST% %UP_PAGE% %LINK_PAGE% %DOWN_PAGE% %END%",
		},
	}

	if i := strings.Index(uri, "?"); i != -1 { //  "/page/haha?p=1234&a=1"
		p.uri = uri[:i]
	}
	p.uri += "?"
	if query != nil {
		p.nowPage = ParsePositiveInt(query["p"], 1)
		query[p.p] = p.pageTag
		for k, v := range query {
			p.uri += url.QueryEscape(k) + "=" + url.QueryEscape(v) + "&"
		}
	} else {
		p.uri += p.p + "=" + p.pageTag
		p.nowPage = 1
	}
	if p.nowPage <= 0 {
		p.nowPage = 1
	}

	p.ListRows = listRows                     // 设置每页显示行数
	p.FirstRow = p.ListRows * (p.nowPage - 1) // 起始行数
	if totalRows <= p.FirstRow {
		p.FirstRow = 0
		p.nowPage = 1
	}
	return p
}

func (p *Page) makeURL(page int) string {
	return strings.Replace(p.uri, p.pageTag, strconv.Itoa(page), -1)
}

func (p *Page) Show() string {
	if p.totalRows == 0 {
		return ""
	}
	var v = float64(p.totalRows) / float64(p.ListRows)
	var totalPages = int(math.Ceil(v))
	if totalPages < p.nowPage {
		p.nowPage = totalPages
	}

	var nowCoolPage = float64(p.rollPage) / 2
	var nowCoolPageCeil = int(math.Ceil(nowCoolPage))

	// ?
	if p.lastSuffix {
		p.Config.Last = strconv.Itoa(totalPages)
	}

	// 上一页
	var upPage = ""
	var upRow = p.nowPage - 1
	if 0 < upRow {
		upPage = `<a class="prev" href="` + p.makeURL(upRow) + `">` + p.Config.Prev + "</a>"
	}

	// 下一页
	var downRow = p.nowPage + 1
	var downPage = ""
	if downRow <= totalPages {
		downPage = `<a class="next" href="` + p.makeURL(downRow) + `">` + p.Config.Next + "</a>"
	}

	// 第一页
	var theFirst = ""
	// 最后一页
	var theEnd = ""

	if totalPages > p.rollPage {
		if (float64(p.nowPage) - nowCoolPage) >= 1 {
			theFirst = `<a class="first" href="` + p.makeURL(1) + `">` + p.Config.First + "</a>"
		}

		if (float64(p.nowPage) + nowCoolPage) < float64(totalPages) {
			theEnd = `<a class="end" href="` + p.makeURL(totalPages) + `">` + p.Config.Last + "</a>"
		}
	}

	// 数字链接
	var linkPage = ""
	var page = 0
	for i := 1; i <= p.rollPage; i++ {
		if (float64(p.nowPage) - nowCoolPage) <= 0 {
			page = i
		} else if (float64(p.nowPage) + nowCoolPage - 1) >= float64(totalPages) {
			page = totalPages - p.rollPage + i
		} else {
			page = p.nowPage - nowCoolPageCeil + i
		}
		if page > 0 && page != p.nowPage {
			if page <= totalPages {
				linkPage += `<a class="num" href="` + p.makeURL(page) + `">` + strconv.Itoa(page) + `</a>`
			} else {
				break
			}
		} else {
			if page > 0 && totalPages != 1 {
				linkPage += `<span class="current">` + strconv.Itoa(page) + `</span>`
			}
		}
	}

	// 替换分页内容
	var pageStr = strReplace(
		[]string{`%HEADER%`, `%NOW_PAGE%`, `%UP_PAGE%`, `%DOWN_PAGE%`, `%FIRST%`, `%LINK_PAGE%`, `%END%`, `%TOTAL_ROW%`, `%TOTAL_PAGE%`},
		[]string{p.Config.Header, strconv.Itoa(p.nowPage), upPage, downPage, theFirst, linkPage, theEnd, strconv.Itoa(p.totalRows), strconv.Itoa(totalPages)},
		p.Config.Theme)
	return `<div class="page">` + pageStr + `</div>`
}
