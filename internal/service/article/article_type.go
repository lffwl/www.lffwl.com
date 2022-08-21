package article

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"www.lffwl.com/internal/dao"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/model/entity"
)

type articleType struct {
	model *gdb.Model
	ctx   context.Context
}

func ArticleType(ctx context.Context) *articleType {
	return &articleType{
		model: dao.ArticleType.Ctx(ctx),
		ctx:   ctx,
	}
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

func (s *articleType) AllIndex() (output []entity.ArticleType, err error) {

	if err = s.model.Scan(&output); err != nil {
		return nil, err
	}

	return
}

func (s *articleType) ConfigAllIndex(ctx context.Context) (output map[int]string, err error) {

	data, err := s.AllIndex()
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		for _, item := range data {
			output[item.Id] = item.Name
		}
	}

	return
}

func (s *articleType) CheckNameUnique(name string, id ...int) (bool, error) {

	m := s.model.Where(dao.ArticleType.Columns().Name, name)

	if len(id) > 0 {
		m = m.WhereNot(dao.ArticleType.Columns().Id, id[0])
	}

	if num, err := m.Count(); err != nil {
		return false, err
	} else if num > 0 {
		return true, nil
	}

	return false, nil
}

func (s *articleType) Store(input model.ArticleTypeStoreInput) error {

	if exist, err := s.CheckNameUnique(input.Name); err != nil {
		return err
	} else if exist {
		return errors.New("文章分类已存在")
	}

	if _, err := s.model.Data(input).Save(); err != nil {
		return err
	}

	return nil
}

func (s *articleType) CheckIdExist(id int) (bool, error) {

	if num, err := s.model.Where(dao.ArticleType.Columns().Id, id).Count(); err != nil {
		return false, err
	} else if num > 0 {
		return true, nil
	}

	return false, nil

}

func (s *articleType) Update(input model.ArticleTypeUpdateInput) error {

	if exist, err := s.CheckIdExist(input.Id); err != nil {
		return err
	} else if !exist {
		return errors.New("ID不存在")
	}

	if exist, err := s.CheckNameUnique(input.Name, input.Id); err != nil {
		return err
	} else if exist {
		return errors.New("文章分类已存在")
	}

	if _, err := s.model.Where(dao.ArticleType.Columns().Id, input.Id).FieldsEx(dao.ArticleType.Columns().Id).Data(input).Update(); err != nil {
		return err
	}

	return nil
}

func (s *articleType) PkDelete(id int) error {

	if exist, err := s.CheckIdExist(id); err != nil {
		return err
	} else if !exist {
		return errors.New("ID不存在")
	}

	if _, err := s.model.Where(dao.ArticleType.Columns().Id, id).Delete(); err != nil {
		return err
	}

	return nil
}
