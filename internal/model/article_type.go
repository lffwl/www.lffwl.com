package model

import "www.lffwl.com/internal/model/entity"

type ArticleTypeIndexInput struct {
	Page  int
	Limit int
	Name  string
}

type ArticleTypeIndexOutput struct {
	List  []entity.ArticleType
	Total int
}

type ArticleTypeStoreInput struct {
	Name string
}

type ArticleTypeUpdateInput struct {
	Id   int
	Name string
}
