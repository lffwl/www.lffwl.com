package controller

import (
	"context"
	marticle "www.lffwl.com/internal/model/article"
	"www.lffwl.com/internal/service/article"

	"www.lffwl.com/api/admin"
)

var (
	ArticleType = cArticleType{}
)

type cArticleType struct{}

func (c *cArticleType) Index(ctx context.Context, req *admin.ArticleTypeIndexReq) (res *admin.ArticleTypeIndexRes, err error) {

	data, err := article.ArticleType(ctx).Index(marticle.ArticleTypeIndexInput{
		Page:  req.Page,
		Limit: req.Size,
		Name:  req.Name,
	})

	res = &admin.ArticleTypeIndexRes{
		List: data.List,
		CommonPageRes: admin.CommonPageRes{
			Page:  req.Page,
			Size:  req.Size,
			Total: data.Total,
		},
	}

	return
}
