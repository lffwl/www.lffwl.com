package blog

import (
	"context"
	"www.lffwl.com/api/blog"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/service/content"
)

var (
	Home = cHome{}
)

type cHome struct{}

func (c *cHome) Index(ctx context.Context, req *blog.HomeIndexReq) (res *blog.HomeIndexRes, err error) {

	typeId := 0
	title := "首页"

	if req.TypeName != "" {
		title = req.TypeName
		typeId = content.ArticleType(ctx).GetIdByName(req.TypeName)
	}

	data, err := content.Article(ctx).Index(model.ArticleIndexInput{
		Page:   req.Page,
		Limit:  req.Size,
		TypeId: typeId,
	})
	if err != nil {
		return nil, err
	}

	content.View().RenderTpl(ctx, "index.html", map[string]interface{}{
		"Title":    title,
		"list":     data.List,
		"typeId":   typeId,
		"pageSize": req.Size,
		"total":    data.Total,
	})

	return
}

func (c *cHome) Search(ctx context.Context, req *blog.HomeSearchReq) (res *blog.HomeSearchRes, err error) {

	title := req.Name + " - 搜索"
	data, err := content.Article(ctx).Index(model.ArticleIndexInput{
		Page:  req.Page,
		Limit: req.Size,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}

	content.View().RenderTpl(ctx, "index.html", map[string]interface{}{
		"Title": title,
		"list":  data.List,
		"name":  req.Name,
	})

	return
}

func (c *cHome) About(ctx context.Context, req *blog.HomeAboutReq) (res *blog.HomeAboutRes, err error) {

	content.View().RenderTpl(ctx, "about.html", map[string]interface{}{
		"typeId": "about",
		"Title":  "关于本站",
	})

	return
}

func (c *cHome) Show(ctx context.Context, req *blog.HomeShowReq) (res *blog.HomeShowRes, err error) {

	typeId := 0
	title := "详情"

	if req.TypeName != "" {
		typeId = content.ArticleType(ctx).GetIdByName(req.TypeName)
	}

	info, err := content.Article(ctx).Show(req.Id)
	if err != nil {
		return nil, err
	}

	title = info.Name

	content.View().RenderTpl(ctx, "show.html", map[string]interface{}{
		"typeId": typeId,
		"Title":  title,
		"info":   info,
	})

	return
}
