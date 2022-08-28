package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type AdminIndexInput struct {
	Page     int
	Limit    int
	Username string
}

type AdminIndexOutput struct {
	List  []AdminIndexOutputList
	Total int
}

type AdminIndexOutputList struct {
	Id        int         `json:"id"        dc:""`
	Username  string      `json:"username"  dc:"用户名"`
	CreatedAt *gtime.Time `json:"createdAt" dc:""`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:""`
}

type AdminStoreInput struct {
	Username string
	Password string
	Roles    []int
}

type AdminUpdateInput struct {
	Id       int
	Username string
	Password string
	Roles    []int
}

type AdminShowOutput struct {
	Id        int
	Username  string
	CreatedAt *gtime.Time
	UpdatedAt *gtime.Time
	Roles     []int
}

type AdminAuthData struct {
	Id       int
	Username string
}
