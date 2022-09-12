package content

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"www.lffwl.com/internal/dao"
	"www.lffwl.com/internal/model"
)

type articleType struct {
	model *gdb.Model
	ctx   context.Context
	cache map[string]int
}

var inArticleType = &articleType{cache: map[string]int{}}

func ArticleType(ctx context.Context) *articleType {
	inArticleType.ctx = ctx
	inArticleType.model = dao.ArticleType.Ctx(ctx)
	return inArticleType
}

func (s *articleType) Index(input model.ArticleTypeIndexInput) (output *model.ArticleTypeIndexOutput, err error) {

	output = &model.ArticleTypeIndexOutput{}

	m := s.model

	if input.Name != "" {
		m = m.WhereLike(dao.ArticleType.Columns().Name, "%"+input.Name+"%")
	}

	if err = m.Page(input.Page, input.Limit).Scan(&output.List); err != nil {
		return nil, err
	}

	if output.Total, err = m.Count(); err != nil {
		return nil, err
	}

	return
}

func (s *articleType) AllIndex() (output []model.ArticleTypeAllIndexOutput, err error) {

	if err = s.model.Fields(dao.ArticleType.Columns().Id, dao.ArticleType.Columns().Name).Scan(&output); err != nil {
		return nil, err
	}

	return
}

func (s *articleType) LoadCache() error {

	data, err := s.AllIndex()
	if err != nil {
		return err
	}

	if len(data) > 0 {
		for _, item := range data {
			s.cache[item.Name] = item.Id
		}
	}

	return nil
}

func (s *articleType) ConfigAllIndex() map[int]string {

	output := map[int]string{}

	if len(s.cache) > 0 {
		for name, id := range s.cache {
			output[id] = name
		}
	}

	return output
}

func (s *articleType) CheckNameUnique(name string, id ...int) bool {

	for val, key := range s.cache {
		if len(id) > 0 {
			if id[0] != key && val == name {
				return true
			}
		} else {
			if val == name {
				return true
			}
		}

	}

	return false
}

func (s *articleType) GetIdByName(name string) int {

	for val, id := range s.cache {

		if val == name {
			return id
		}

	}

	return 0
}

func (s *articleType) Store(input model.ArticleTypeStoreInput) error {

	if exist := s.CheckNameUnique(input.Name); exist {
		return errors.New("文章分类已存在")
	}

	insertId, err := s.model.Data(input).InsertAndGetId()
	if err != nil {
		return err
	}

	s.cache[input.Name] = int(insertId)

	return nil
}

func (s *articleType) CheckIdExist(id int) bool {

	for _, val := range s.cache {
		if val == id {
			return true
		}
	}

	return false

}

func (s *articleType) Update(input model.ArticleTypeUpdateInput) error {

	if exist := s.CheckIdExist(input.Id); !exist {
		return errors.New("ID不存在")
	}

	if exist := s.CheckNameUnique(input.Name, input.Id); exist {
		return errors.New("文章分类已存在")
	}

	if _, err := s.model.Where(dao.ArticleType.Columns().Id, input.Id).FieldsEx(dao.ArticleType.Columns().Id).Data(input).Update(); err != nil {
		return err
	}

	s.deleteCache(input.Id)

	s.cache[input.Name] = input.Id

	return nil
}

func (s *articleType) PkDelete(id int) error {

	if exist := s.CheckIdExist(id); !exist {
		return errors.New("ID不存在")
	}

	if _, err := s.model.Where(dao.ArticleType.Columns().Id, id).Delete(); err != nil {
		return err
	}

	s.deleteCache(id)

	return nil
}

func (s *articleType) deleteCache(id int) {
	for name, valId := range s.cache {
		if id == valId {
			delete(s.cache, name)
			break
		}
	}
}
