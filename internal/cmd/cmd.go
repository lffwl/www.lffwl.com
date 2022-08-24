package cmd

import (
	"context"
	"www.lffwl.com/internal/controller/manage"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "www.lffwl.com",
		Usage: "www.lffwl.com",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/manage", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					manage.ArticleType,
					manage.Article,
					manage.Api,
					manage.Role,
					manage.Admin,
				)
			})
			s.Run()
			return nil
		},
	}
)
