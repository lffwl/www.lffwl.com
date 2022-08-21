package manage

import (
	"github.com/gogf/gf/v2/frame/g"
	"www.lffwl.com/internal/model/entity"
)

type ApiIndexReq struct {
	g.Meta `path:"/api" tags:"Api" method:"get" summary:"api列表"`
}
type ApiIndexRes struct {
	List   []entity.Api      `json:"list" dc:"集合"`
	Config ApiIndexResConfig `json:"config" dc:"配置"`
}

type ApiIndexResConfig struct {
	Method map[string]string `json:"method" dc:"请求类型"`
	Type   map[string]string `json:"type" dc:"api类型"`
}

type ApiStoreReq struct {
	g.Meta `path:"/api" tags:"Api" method:"post" summary:"新增api"`
	Menu   string `json:"menu" dc:"菜单地址"`
	Name   string `json:"name" v:"required#api名称不能为空" dc:"api名称"`
	Key    string `json:"key" v:"required#标识不能为空" dc:"标识"`
	Path   string `json:"path" v:"required-with:method#标识不能为空" dc:"请求地址"`
	Method int    `json:"method" v:"required-with:path#请求类型不能为空" dc:"请求类型"`
	Type   int    `json:"type" v:"min:1#api类型必须要大于0" dc:"api类型"`
	Pid    int    `json:"pid" dc:"上级Api"`
}
type ApiStoreRes struct{}

type ApiUpdateReq struct {
	g.Meta `path:"/api/{id}" tags:"Api" method:"post" summary:"更新api"`
	Id     uint   `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
	Menu   string `json:"menu" dc:"菜单地址"`
	Name   string `json:"name" v:"required#api名称不能为空" dc:"api名称"`
	Key    string `json:"key" v:"required#标识不能为空" dc:"标识"`
	Path   string `json:"path" v:"required-with:method#标识不能为空" dc:"请求地址"`
	Method int    `json:"method" v:"required-with:path#请求类型不能为空" dc:"请求类型"`
	Type   int    `json:"type" v:"min:1#api类型必须要大于0" dc:"api类型"`
	Pid    int    `json:"pid" dc:"上级Api"`
}
type ApiUpdateRes struct{}

type ApiDeleteReq struct {
	g.Meta `path:"/api/{id}/delete" tags:"Api" method:"post" summary:"删除api"`
	Id     uint `in:"path" v:"min:1#Id必须要大于0" dc:"id"`
}
type ApiDeleteRes struct{}
