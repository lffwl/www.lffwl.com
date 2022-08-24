package manage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"www.lffwl.com/api/manage"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/service/auth"
)

var (
	Api = cApi{}
)

type cApi struct{}

func (c *cApi) Index(ctx context.Context, req *manage.ApiIndexReq) (res *manage.ApiIndexRes, err error) {

	data, err := auth.Api(ctx).AllIndex()
	if err != nil {
		return nil, err
	}

	res = &manage.ApiIndexRes{
		List: data,
	}

	res.Config.Method = g.Cfg().MustGet(ctx, "auth.method").MapStrStr()
	res.Config.Type = g.Cfg().MustGet(ctx, "auth.type").MapStrStr()

	return
}

func (c *cApi) Store(ctx context.Context, req *manage.ApiStoreReq) (res *manage.ApiStoreRes, err error) {

	err = auth.Api(ctx).Store(model.ApiStoreInput{
		Menu:   req.Menu,
		Name:   req.Name,
		Key:    req.Key,
		Path:   req.Path,
		Method: req.Method,
		Type:   req.Type,
		Pid:    req.Pid,
	})

	return
}

func (c *cApi) Update(ctx context.Context, req *manage.ApiUpdateReq) (res *manage.ApiUpdateRes, err error) {

	err = auth.Api(ctx).Update(model.ApiUpdateInput{
		Id:     int(req.Id),
		Menu:   req.Menu,
		Name:   req.Name,
		Key:    req.Key,
		Path:   req.Path,
		Method: req.Method,
		Type:   req.Type,
		Pid:    req.Pid,
	})

	return
}

func (c *cApi) Delete(ctx context.Context, req *manage.ApiDeleteReq) (res *manage.ApiDeleteRes, err error) {

	err = auth.Api(ctx).PkDelete(int(req.Id))

	return
}
