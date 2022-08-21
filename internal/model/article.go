package model

import "www.lffwl.com/internal/model/entity"

type ArticleIndexInput struct {
	Page   int
	Limit  int
	Name   string
	TypeId int
}

type ArticleIndexOutput struct {
	List  []entity.Article
	Total int
}

type ArticleStoreInput struct {
	TypeId  int
	Content string
	Name    string
}

type ArticleUpdateInput struct {
	Id      int
	Name    string
	TypeId  int
	Content string
}
