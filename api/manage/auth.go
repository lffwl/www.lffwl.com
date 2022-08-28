package manage

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"www.lffwl.com/internal/model/entity"
)

type AuthLoginReq struct {
	g.Meta   `path:"/auth/login" tags:"Auth" method:"post" summary:"登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#管理员密码不能为空" dc:"管理员密码"`
}
type AuthLoginRes struct {
	Username string `json:"username" dc:"用户名"`
	Token    string `json:"token" dc:"Token"`
}

type AuthCenterReq struct {
	g.Meta `path:"/auth/center" tags:"Auth" method:"get" summary:"个人中心"`
}
type AuthCenterRes struct {
	Id        int          `json:"id"`
	Username  string       `json:"username"`
	CreatedAt *gtime.Time  `json:"createdAt"`
	UpdatedAt *gtime.Time  `json:"updatedAt"`
	Roles     []int        `json:"roles"`
	Auth      []entity.Api `json:"auth"`
}
