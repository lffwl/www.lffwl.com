package manage

import (
	"context"
	"www.lffwl.com/api/manage"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/service/auth"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct{}

func (c *cAdmin) Index(ctx context.Context, req *manage.AdminIndexReq) (res *manage.AdminIndexRes, err error) {

	data, err := auth.Admin(ctx).Index(model.AdminIndexInput{
		Page:     req.Page,
		Limit:    req.Size,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	res = &manage.AdminIndexRes{
		List: data.List,
		CommonPageRes: manage.CommonPageRes{
			Page:  req.Page,
			Size:  req.Size,
			Total: data.Total,
		},
	}

	return
}

func (c *cAdmin) Store(ctx context.Context, req *manage.AdminStoreReq) (res *manage.AdminStoreRes, err error) {

	err = auth.Admin(ctx).Store(model.AdminStoreInput{
		Username: req.Username,
		Password: req.Password,
		Roles:    req.Roles,
	})

	return
}

func (c *cAdmin) Update(ctx context.Context, req *manage.AdminUpdateReq) (res *manage.AdminUpdateRes, err error) {

	err = auth.Admin(ctx).Update(model.AdminUpdateInput{
		Id:       int(req.Id),
		Username: req.Username,
		Password: req.Password,
		Roles:    req.Roles,
	})

	return
}

func (c *cAdmin) Delete(ctx context.Context, req *manage.AdminDeleteReq) (res *manage.AdminDeleteRes, err error) {

	err = auth.Admin(ctx).PkDelete(int(req.Id))

	return
}

func (c *cAdmin) Show(ctx context.Context, req *manage.AdminShowReq) (res *manage.AdminShowRes, err error) {

	data, err := auth.Admin(ctx).Show(int(req.Id))

	res = &manage.AdminShowRes{
		Id:        data.Id,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Roles:     data.Roles,
	}

	return
}
