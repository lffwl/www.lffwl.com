package manage

import (
	"github.com/gogf/gf/v2/frame/g"
	"www.lffwl.com/internal/model/entity"
)

type ArticleIndexReq struct {
	g.Meta `path:"/article" tags:"Article" method:"get" summary:"文章列表"`
	Name   string `json:"name" in:"query"  dc:"文章名称"`
	CommonPageReq
}
type ArticleIndexRes struct {
	List   []entity.Article      `json:"list" dc:"集合"`
	Config ArticleIndexResConfig `json:"config" dc:"配置"`
	CommonPageRes
}

type ArticleIndexResConfig struct {
	ArticleType map[int]string `json:"articleType" dc:"文章分类"`
}

type ArticleStoreReq struct {
	g.Meta  `path:"/article" tags:"Article" method:"post" summary:"新增文章"`
	Name    string `json:"name" v:"required#文章名称不能为空" dc:"文章名称"`
	Content string `json:"content" dc:"文章内容"`
	Desc    string `json:"desc" dc:"描叙"`
	TypeId  int    `json:"typeId" dc:"文章分类ID"`
}
type ArticleStoreRes struct{}

type ArticleUpdateReq struct {
	g.Meta  `path:"/article/{id}" tags:"Article" method:"post" summary:"更新文章"`
	Id      uint   `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
	Name    string `json:"name" v:"required#文章名称不能为空" dc:"文章名称"`
	Desc    string `json:"desc" dc:"描叙"`
	Content string `json:"content" dc:"文章内容"`
	TypeId  int    `json:"typeId" dc:"文章分类ID"`
}
type ArticleUpdateRes struct{}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/{id}/delete" tags:"Article" method:"post" summary:"删除文章"`
	Id     uint `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
}
type ArticleDeleteRes struct{}

type ArticleShowReq struct {
	g.Meta `path:"/article/{id}" tags:"Article" method:"get" summary:"文章详情"`
	Id     uint `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
}
type ArticleShowRes struct {
	Id      int    `json:"id"  dc:"id"`
	Name    string `json:"name" dc:"文章名称"`
	Desc    string `json:"desc" dc:"描叙"`
	Content string `json:"content" dc:"文章内容"`
	TypeId  int    `json:"typeId" dc:"文章分类ID"`
}
