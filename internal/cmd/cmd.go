package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"www.lffwl.com/internal/controller/blog"
	"www.lffwl.com/internal/controller/manage"
	"www.lffwl.com/internal/service/auth"
	"www.lffwl.com/internal/service/content"
)

var (
	Main = gcmd.Command{
		Name:  "www.lffwl.com",
		Usage: "www.lffwl.com",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			auth.Jwt = auth.NewJwt(ctx)
			// ArticleType Cache
			if err := content.ArticleType(ctx).LoadCache(); err != nil {
				return err
			}

			// seo
			content.View().Init(
				g.Cfg("content").MustGet(ctx, "seo.title").String(),
				g.Cfg("content").MustGet(ctx, "seo.seoKey").String(),
				g.Cfg("content").MustGet(ctx, "seo.seoDesc").String(),
			)

			s := g.Server()
			s.Group("/manage", func(group *ghttp.RouterGroup) {
				group.Middleware(auth.Auth, ghttp.MiddlewareHandlerResponse)
				group.Bind(
					manage.ArticleType,
					manage.Article,
					manage.Api,
					manage.Role,
					manage.Admin,
					manage.Auth,
					manage.File,
				)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					blog.Home,
				)
			})
			s.Run()
			return nil
		},
	}
)
