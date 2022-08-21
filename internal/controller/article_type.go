package controller

import (
	"context"
	marticle "www.lffwl.com/internal/model"
	"www.lffwl.com/internal/service/article"

	"www.lffwl.com/api/manage"
)

var (
	ArticleType = cArticleType{}
)

type cArticleType struct{}

func (c *cArticleType) Index(ctx context.Context, req *manage.ArticleTypeIndexReq) (res *manage.ArticleTypeIndexRes, err error) {

	data, err := article.ArticleType(ctx).Index(marticle.ArticleTypeIndexInput{
		Page:  req.Page,
		Limit: req.Size,
		Name:  req.Name,
	})

	res = &manage.ArticleTypeIndexRes{
		List: data.List,
		CommonPageRes: manage.CommonPageRes{
			Page:  req.Page,
			Size:  req.Size,
			Total: data.Total,
		},
	}

	return
}

func (c *cArticleType) Store(ctx context.Context, req *manage.ArticleTypeStoreReq) (res *manage.ArticleTypeStoreRes, err error) {

	err = article.ArticleType(ctx).Store(marticle.ArticleTypeStoreInput{
		Name: req.Name,
	})

	return
}

func (c *cArticleType) Update(ctx context.Context, req *manage.ArticleTypeUpdateReq) (res *manage.ArticleTypeUpdateRes, err error) {

	err = article.ArticleType(ctx).Update(marticle.ArticleTypeUpdateInput{
		Id:   int(req.Id),
		Name: req.Name,
	})

	return
}

func (c *cArticleType) Delete(ctx context.Context, req *manage.ArticleTypeDeleteReq) (res *manage.ArticleTypeDeleteRes, err error) {

	err = article.ArticleType(ctx).PkDelete(int(req.Id))

	return
}
