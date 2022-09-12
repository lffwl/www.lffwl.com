package content

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type view struct {
	seoTitle string
	seoKey   string
	seoDesc  string
}

var inView = &view{}

func View() *view {
	return inView
}

func (s *view) Init(seo ...string) {

	if len(seo) > 0 {
		s.seoTitle = seo[0]
	}

	if len(seo) > 1 {
		s.seoKey = seo[1]
	}

	if len(seo) > 2 {
		s.seoDesc = seo[2]
	}
}

func (s *view) RenderTpl(ctx context.Context, tpl string, data ...map[string]interface{}) {

	res := map[string]interface{}{}

	if len(data) > 0 {
		res = data[0]
	}

	res["Title"] = gconv.String(res["Title"]) + s.seoTitle
	res["SeoKey"] = gconv.String(res["SeoKey"]) + s.seoKey
	res["SeoDesc"] = gconv.String(res["SeoDesc"]) + s.seoDesc

	res["typeMap"] = ArticleType(ctx).ConfigAllIndex()
	res["Process"] = &process{}

	request := g.RequestFromCtx(ctx)

	page := request.GetPage(gconv.Int(res["total"]), gconv.Int(res["pageSize"]))

	page.SpanStyle = ""
	page.LinkStyle = ""
	res["page"] = page.GetContent(3)
	request.Response.WriteTpl(tpl, res)
	request.Exit()
}

type process struct {
}

func (v *process) MapIntOutStr(data map[int]string, key int) string {
	return data[key]
}
