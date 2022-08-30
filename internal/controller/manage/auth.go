package manage

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lffwl/utility/dataproc"
	"www.lffwl.com/api/manage"
	"www.lffwl.com/internal/consts"
	"www.lffwl.com/internal/model"
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

	info, err := adminService.GetInfoByUserName(req.Username)
	if err != nil {
		return nil, err
	}

	if !dataproc.SecretHashCompare(info.Password, req.Password) {
		return nil, errors.New("账号密码错误")
	}

	token, err := auth.Jwt.CreateToken(model.AdminAuthData{
		Id:       info.Id,
		Username: info.Username,
	})
	if err != nil {
		return nil, err
	}

	res = &manage.AuthLoginRes{
		Username: info.Username,
		Token:    token,
	}

	return
}

func (c *cAuth) Center(ctx context.Context, req *manage.AuthCenterReq) (res *manage.AuthCenterRes, err error) {

	info, err := auth.Admin(ctx).Show(gconv.Int(ctx.Value(consts.AuthCtxAdminIdKey)))
	if err != nil {
		return nil, err
	}

	res = &manage.AuthCenterRes{
		Id:        info.Id,
		Username:  info.Username,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
		Roles:     info.Roles,
	}

	if info.Id != g.Cfg("auth").MustGet(ctx, "superRoleId").Int() {
		apis, err := auth.RoleApi(ctx).GetApisByRoles(info.Roles)
		if err != nil {
			return nil, err
		}

		if len(apis) > 0 {
			if res.Auth, err = auth.Api(ctx).Index(model.ApiIndexInput{Ids: apis}); err != nil {
				return nil, err
			}
		}
	} else {
		if res.Auth, err = auth.Api(ctx).AllIndex(); err != nil {
			return nil, err
		}
	}

	return
}
