package model

import "www.lffwl.com/internal/model/entity"

type RoleIndexInput struct {
	Page  int
	Limit int
	Name  string
}

type RoleIndexOutput struct {
	List  []entity.Role
	Total int
}

type RoleStoreInput struct {
	Name string
	Apis []int
}

type RoleUpdateInput struct {
	Id   int
	Name string
	Apis []int
}

type RoleShowOutput struct {
	entity.Role
	Apis []int
}
