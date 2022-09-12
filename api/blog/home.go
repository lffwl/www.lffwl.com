package blog

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HomeIndexReq struct {
	g.Meta   `path:"/*typeName" method:"get" summary:"分类列表页面" tags:"Home"`
	TypeName string `json:"typeName" in:"path" dc:"分类名称"`
	CommonPageReq
}
type HomeIndexRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type HomeSearchReq struct {
	g.Meta `path:"/search" method:"get" summary:"搜索页面" tags:"Home"`
	Name   string `json:"name" in:"query" dc:"名称"`
	CommonPageReq
}
type HomeSearchRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type HomeAboutReq struct {
	g.Meta `path:"/about.html" method:"get" summary:"关于本站" tags:"Home"`
}
type HomeAboutRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type HomeShowReq struct {
	g.Meta   `path:"/{typeName}/{id}.html" method:"get" summary:"详情页面" tags:"Home"`
	TypeName string `json:"typeName" in:"path" dc:"分类名称"`
	Id       int    `json:"id" in:"path" dc:"ID"`
}
type HomeShowRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
