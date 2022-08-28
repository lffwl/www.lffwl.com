package cmd

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"www.lffwl.com/internal/controller/manage"
	"www.lffwl.com/internal/service/auth"
	"www.lffwl.com/utility/response"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "www.lffwl.com",
		Usage: "www.lffwl.com",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			auth.Jwt = auth.NewJwt(ctx)
			s := g.Server()
			s.Group("/manage", func(group *ghttp.RouterGroup) {
				group.Middleware(auth.Auth, func(r *ghttp.Request) {

					r.Middleware.Next()

					// 如果已经有返回内容，那么该中间件什么也不做
					if r.Response.BufferLength() > 0 {
						return
					}

					var (
						err             = r.GetError()
						res             = r.GetHandlerResponse()
						code gcode.Code = gcode.CodeOK
					)
					if err != nil {
						code = gerror.Code(err)
						if code == gcode.CodeNil {
							code = gcode.CodeInternalError
						}
						response.JsonExit(r, code.Code(), err.Error())
					} else {
						if r.IsAjaxRequest() {
							response.JsonExit(r, code.Code(), "", res)
						} else {
							// 什么都不做，业务API自行处理模板渲染的成功逻辑。
						}
					}
				})
				group.Bind(
					manage.ArticleType,
					manage.Article,
					manage.Api,
					manage.Role,
					manage.Admin,
					manage.Auth,
				)
			})
			s.Run()
			return nil
		},
	}
)
