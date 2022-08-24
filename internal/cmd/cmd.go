package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"www.lffwl.com/internal/controller"
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
					controller.ArticleType,
					controller.Article,
					controller.Api,
					controller.Role,
				)
			})
			s.Run()
			return nil
		},
	}
)
