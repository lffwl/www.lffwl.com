package controller

import (
	"context"
	"www.lffwl.com/api/manage"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/model/entity"
	"www.lffwl.com/internal/service/auth"
)

var (
	Role = cRole{}
)

type cRole struct{}

func (c *cRole) Index(ctx context.Context, req *manage.RoleIndexReq) (res *manage.RoleIndexRes, err error) {

	data, err := auth.Role(ctx).Index(model.RoleIndexInput{
		Page:  req.Page,
		Limit: req.Size,
		Name:  req.Name,
	})
	if err != nil {
		return nil, err
	}

	res = &manage.RoleIndexRes{
		List: data.List,
		CommonPageRes: manage.CommonPageRes{
			Page:  req.Page,
			Size:  req.Size,
			Total: data.Total,
		},
	}

	return
}

func (c *cRole) Store(ctx context.Context, req *manage.RoleStoreReq) (res *manage.RoleStoreRes, err error) {

	err = auth.Role(ctx).Store(model.RoleStoreInput{
		Name: req.Name,
		Apis: req.Apis,
	})

	return
}

func (c *cRole) Update(ctx context.Context, req *manage.RoleUpdateReq) (res *manage.RoleUpdateRes, err error) {

	err = auth.Role(ctx).Update(model.RoleUpdateInput{
		Id:   int(req.Id),
		Name: req.Name,
		Apis: req.Apis,
	})

	return
}

func (c *cRole) Delete(ctx context.Context, req *manage.RoleDeleteReq) (res *manage.RoleDeleteRes, err error) {

	err = auth.Role(ctx).PkDelete(int(req.Id))

	return
}

func (c *cRole) Show(ctx context.Context, req *manage.RoleShowReq) (res *manage.RoleShowRes, err error) {

	data, err := auth.Role(ctx).Show(int(req.Id))

	res = &manage.RoleShowRes{
		Role: entity.Role{
			Id:        data.Id,
			Name:      data.Name,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
		Apis: res.Apis,
	}

	return
}
