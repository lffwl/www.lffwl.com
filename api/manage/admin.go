package manage

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"www.lffwl.com/internal/model"
)

type AdminIndexReq struct {
	g.Meta   `path:"/admin" tags:"Admin" method:"get" summary:"管理员列表"`
	Username string `json:"username" in:"query"  dc:"用户名"`
	CommonPageReq
}
type AdminIndexRes struct {
	List []model.AdminIndexOutputList `json:"list" dc:"集合"`
	CommonPageRes
}

type AdminStoreReq struct {
	g.Meta   `path:"/admin" tags:"Admin" method:"post" summary:"新增管理员"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#管理员密码不能为空" dc:"管理员密码"`
	Roles    []int  `json:"apis" dc:"管理员角色"`
}
type AdminStoreRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/admin/{id}" tags:"Admin" method:"post" summary:"更新管理员"`
	Id       uint   `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password"  dc:"管理员密码"`
	Roles    []int  `json:"apis" dc:"管理员角色"`
}
type AdminUpdateRes struct{}

type AdminDeleteReq struct {
	g.Meta `path:"/admin/{id}/delete" tags:"Admin" method:"post" summary:"删除管理员"`
	Id     uint `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
}
type AdminDeleteRes struct{}

type AdminShowReq struct {
	g.Meta `path:"/admin/{id}" tags:"Admin" method:"get" summary:"管理员详情"`
	Id     uint `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
}
type AdminShowRes struct {
	Id        int         `json:"id"        dc:""`
	Username  string      `json:"username"  dc:"用户名"`
	CreatedAt *gtime.Time `json:"createdAt" dc:""`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:""`
	Roles     []int
}
