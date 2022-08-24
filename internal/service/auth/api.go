package auth

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"www.lffwl.com/internal/dao"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/model/entity"
)

type api struct {
	model *gdb.Model
	ctx   context.Context
}

func Api(ctx context.Context) *api {
	return &api{
		model: dao.Api.Ctx(ctx),
		ctx:   ctx,
	}
}

func (s *api) AllIndex() (output []entity.Api, err error) {

	if err = s.model.Scan(&output); err != nil {
		return nil, err
	}

	return

}

func (s *api) CheckPathMethodExist(path string, method int, id ...int) (bool, error) {

	m := s.model

	if len(id) > 0 {
		m = m.WhereNot(dao.Api.Columns().Id, id[0])
	}

	if num, err := m.Where(dao.Api.Columns().Path, path).Where(dao.Api.Columns().Method, method).Count(); err != nil {
		return false, err
	} else if num > 0 {
		return true, nil
	}

	return false, nil
}

func (s *api) CheckFieldExist(field string, value interface{}, id ...int) (bool, error) {
	m := s.model

	if len(id) > 0 {
		m = m.WhereNot(dao.Api.Columns().Id, id[0])
	}

	if num, err := m.Where(field, value).Count(); err != nil {
		return false, err
	} else if num > 0 {
		return true, nil
	}

	return false, nil
}

func (s *api) CheckIdExist(id int) (bool, error) {
	return s.CheckFieldExist(dao.Api.Columns().Id, id)
}

func (s *api) CheckKeyExist(key string, id ...int) (bool, error) {
	return s.CheckFieldExist(dao.Api.Columns().Key, key, id...)
}

func (s *api) Store(input model.ApiStoreInput) (err error) {

	if input.Pid > 0 {
		if exist, err := s.CheckIdExist(input.Pid); err != nil {
			return err
		} else if !exist {
			return errors.New("上级不存在")
		}
	}

	if exist, err := s.CheckKeyExist(input.Key); err != nil {
		return err
	} else if exist {
		return errors.New("标识已存在")
	}

	if input.Path != "" || input.Method > 0 {
		if exist, err := s.CheckPathMethodExist(input.Path, input.Method); err != nil {
			return err
		} else if exist {
			return errors.New("请求地址和请求类型已存在")
		}
	}

	if _, err = s.model.Data(input).Save(); err != nil {
		return err
	}

	return nil
}

func (s *api) Update(input model.ApiUpdateInput) (err error) {

	if exist, err := s.CheckIdExist(input.Id); err != nil {
		return err
	} else if !exist {
		return errors.New("ID不存在")
	}

	if input.Pid > 0 {
		if exist, err := s.CheckIdExist(input.Pid); err != nil {
			return err
		} else if !exist {
			return errors.New("上级不存在")
		}
	}

	if exist, err := s.CheckKeyExist(input.Key, input.Id); err != nil {
		return err
	} else if exist {
		return errors.New("标识已存在")
	}

	if input.Path != "" || input.Method > 0 {
		if exist, err := s.CheckPathMethodExist(input.Path, input.Method, input.Id); err != nil {
			return err
		} else if exist {
			return errors.New("请求地址和请求类型已存在")
		}
	}

	if _, err = s.model.FieldsEx(dao.Api.Columns().Id).Where(dao.Api.Columns().Id, input.Id).Data(input).Update(); err != nil {
		return err
	}

	return nil
}

func (s *api) PkDelete(id int) error {

	if exist, err := s.CheckIdExist(id); err != nil {
		return err
	} else if !exist {
		return errors.New("ID不存在")
	}

	if _, err := s.model.Where(dao.Api.Columns().Id, id).Delete(); err != nil {
		return err
	}

	return nil
}
