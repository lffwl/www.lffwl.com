package manage

import (
	"context"
	"www.lffwl.com/api/manage"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/service/content"
)

var (
	Article = cArticle{}
)

type cArticle struct{}

func (c *cArticle) Index(ctx context.Context, req *manage.ArticleIndexReq) (res *manage.ArticleIndexRes, err error) {

	data, err := content.Article(ctx).Index(model.ArticleIndexInput{
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

	res.Config.ArticleType = content.ArticleType(ctx).ConfigAllIndex()

	return
}

func (c *cArticle) Store(ctx context.Context, req *manage.ArticleStoreReq) (res *manage.ArticleStoreRes, err error) {

	err = content.Article(ctx).Store(model.ArticleStoreInput{
		TypeId:  req.TypeId,
		Content: req.Content,
		Name:    req.Name,
		Desc:    req.Desc,
	})

	return
}

func (c *cArticle) Update(ctx context.Context, req *manage.ArticleUpdateReq) (res *manage.ArticleUpdateRes, err error) {

	err = content.Article(ctx).Update(model.ArticleUpdateInput{
		Id:      int(req.Id),
		TypeId:  req.TypeId,
		Content: req.Content,
		Name:    req.Name,
		Desc:    req.Desc,
	})

	return
}

func (c *cArticle) Delete(ctx context.Context, req *manage.ArticleDeleteReq) (res *manage.ArticleDeleteRes, err error) {

	err = content.Article(ctx).PkDelete(int(req.Id))

	return
}

func (c *cArticle) Show(ctx context.Context, req *manage.ArticleShowReq) (res *manage.ArticleShowRes, err error) {

	data, err := content.Article(ctx).Show(int(req.Id))
	if err != nil {
		return nil, err
	}

	res = &manage.ArticleShowRes{
		Id:      data.Id,
		Name:    data.Name,
		TypeId:  data.TypeId,
		Content: data.Content,
		Desc:    data.Desc,
	}

	return
}
