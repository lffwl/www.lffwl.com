package manage

import (
	"context"
	"errors"
	"www.lffwl.com/api/manage"
	"www.lffwl.com/internal/service/auth"
)

var (
	Auth = cAuth{}
)

type cAuth struct{}

func (c *cAuth) Login(ctx context.Context, req *manage.AuthLoginReq) (res *manage.AuthLoginRes, err error) {

	adminService := auth.Admin(ctx)

	if exist, err := adminService.CheckUsernameExist(req.Username); err != nil {
		return nil, err
	} else if !exist {
		return nil, errors.New("账号密码错误")
	}

	return
}
