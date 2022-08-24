package manage

import (
	"github.com/gogf/gf/v2/frame/g"
	"www.lffwl.com/internal/model/entity"
)

type RoleIndexReq struct {
	g.Meta `path:"/role" tags:"Role" method:"get" summary:"角色列表"`
	Name   string `json:"name" in:"query"  dc:"角色名称"`
	CommonPageReq
}
type RoleIndexRes struct {
	List []entity.Role `json:"list" dc:"集合"`
	CommonPageRes
}

type RoleStoreReq struct {
	g.Meta `path:"/role" tags:"Role" method:"post" summary:"新增角色"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Apis   []int  `json:"apis" dc:"角色权限"`
}
type RoleStoreRes struct{}

type RoleUpdateReq struct {
	g.Meta `path:"/role/{id}" tags:"Role" method:"post" summary:"更新角色"`
	Id     uint   `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Apis   []int  `json:"apis" dc:"角色权限"`
}
type RoleUpdateRes struct{}

type RoleDeleteReq struct {
	g.Meta `path:"/role/{id}/delete" tags:"Role" method:"post" summary:"删除角色"`
	Id     uint `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
}
type RoleDeleteRes struct{}

type RoleShowReq struct {
	g.Meta `path:"/role/{id}" tags:"Role" method:"get" summary:"角色详情"`
	Id     uint `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
}
type RoleShowRes struct {
	entity.Role
	Apis []int
}
