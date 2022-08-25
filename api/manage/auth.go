package manage

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AuthLoginReq struct {
	g.Meta   `path:"/auth/login" tags:"Auth" method:"post" summary:"登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#管理员密码不能为空" dc:"管理员密码"`
}
type AuthLoginRes struct{}
