// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Api is the golang structure of table api for DAO operations like Where/Data.
type Api struct {
	g.Meta    `orm:"table:api, do:true"`
	Id        interface{} //
	Meun      interface{} // 菜单地址
	Name      interface{} // 名称
	Key       interface{} // 标识
	Path      interface{} // 请求地址
	Method    interface{} // 请求类型
	Type      interface{} // 菜单类型
	Pid       interface{} // 上级ID
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
