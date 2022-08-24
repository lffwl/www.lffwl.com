package manage

import (
	"context"
	"www.lffwl.com/api/manage"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/service/article"
)

var (
	Article = cArticle{}
)

type cArticle struct{}

func (c *cArticle) Index(ctx context.Context, req *manage.ArticleIndexReq) (res *manage.ArticleIndexRes, err error) {

	data, err := article.Article(ctx).Index(model.ArticleIndexInput{
		Page:  req.Page,
		Limit: req.Size,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}

	res = &manage.ArticleIndexRes{
		List: data.List,
		CommonPageRes: manage.CommonPageRes{
			Page:  req.Page,
			Size:  req.Size,
			Total: data.Total,
		},
	}

	res.Config.ArticleType, err = article.ArticleType(ctx).ConfigAllIndex()
	if err != nil {
		return nil, err
	}

	return
}

func (c *cArticle) Store(ctx context.Context, req *manage.ArticleStoreReq) (res *manage.ArticleStoreRes, err error) {

	err = article.Article(ctx).Store(model.ArticleStoreInput{
		TypeId:  req.TypeId,
		Content: req.Content,
		Name:    req.Name,
	})

	return
}

func (c *cArticle) Update(ctx context.Context, req *manage.ArticleUpdateReq) (res *manage.ArticleUpdateRes, err error) {

	err = article.Article(ctx).Update(model.ArticleUpdateInput{
		Id:      int(req.Id),
		TypeId:  req.TypeId,
		Content: req.Content,
		Name:    req.Name,
	})

	return
}

func (c *cArticle) Delete(ctx context.Context, req *manage.ArticleDeleteReq) (res *manage.ArticleDeleteRes, err error) {

	err = article.Article(ctx).PkDelete(int(req.Id))

	return
}
