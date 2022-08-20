package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"www.lffwl.com/internal/model/entity"
)

type ArticleTypeIndexReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
	Name   string `json:"name" in:"query"  dc:"文章分类名称"`
	CommonPageReq
}
type ArticleTypeIndexRes struct {
	List []entity.ArticleType `json:"list" dc:"列表"`
	CommonPageRes
}
