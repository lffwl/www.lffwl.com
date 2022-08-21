package manage

import (
	"github.com/gogf/gf/v2/frame/g"
	"www.lffwl.com/internal/model/entity"
)

type ArticleTypeIndexReq struct {
	g.Meta `path:"/manage/article_type" tags:"ArticleType" method:"get" summary:"文章分类列表"`
	Name   string `json:"name" in:"query"  dc:"文章分类名称"`
	CommonPageReq
}
type ArticleTypeIndexRes struct {
	List []entity.ArticleType `json:"list" dc:"集合"`
	CommonPageRes
}

type ArticleTypeStoreReq struct {
	g.Meta `path:"/manage/article_type" tags:"ArticleType" method:"post" summary:"新增文章分类"`
	Name   string `json:"name"  dc:"文章分类名称"`
}
type ArticleTypeStoreRes struct{}

type ArticleTypeUpdateReq struct {
	g.Meta `path:"/manage/article_type/{Id}" tags:"ArticleType" method:"post" summary:"更新文章分类"`
	Id     uint   `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
	Name   string `json:"name" dc:"文章分类名称"`
}
type ArticleTypeUpdateRes struct{}

type ArticleTypeDeleteReq struct {
	g.Meta `path:"/manage/article_type/{Id}/delete" tags:"ArticleType" method:"post" summary:"删除文章分类"`
	Id     uint `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
}
type ArticleTypeDeleteRes struct{}
