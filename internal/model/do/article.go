// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure of table article for DAO operations like Where/Data.
type Article struct {
	g.Meta    `orm:"table:article, do:true"`
	Id        interface{} //
	Name      interface{} // 文章名称
	TypeId    interface{} // 文章分类ID
	Content   interface{} // 内容路径
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}