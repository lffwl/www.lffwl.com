package auth

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"www.lffwl.com/internal/dao"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/model/entity"
)

type role struct {
	model *gdb.Model
	ctx   context.Context
}

func Role(ctx context.Context) *role {
	return &role{
		model: dao.Role.Ctx(ctx),
		ctx:   ctx,
	}
}

func (s *role) Index(input model.RoleIndexInput) (output *model.RoleIndexOutput, err error) {

	output = &model.RoleIndexOutput{}
	m := s.model

	if input.Name != "" {
		m = m.WhereLike(dao.Role.Columns().Name, "%"+input.Name+"%")
	}

	if err = m.Page(input.Page, input.Limit).Scan(&output.List); err != nil {
		return nil, err
	}

	if output.Total, err = m.Count(); err != nil {
		return nil, err
	}

	return

}

func (s *role) CheckFieldExist(field string, value interface{}, id ...int) (bool, error) {
	m := s.model

	if len(id) > 0 {
		m = m.WhereNot(dao.Role.Columns().Id, id[0])
	}

	if num, err := m.Where(field, value).Count(); err != nil {
		return false, err
	} else if num > 0 {
		return true, nil
	}

	return false, nil
}

func (s *role) CheckIdExist(id int) (bool, error) {
	return s.CheckFieldExist(dao.Role.Columns().Id, id)
}

func (s *role) CheckNameExist(name string, id ...int) (bool, error) {
	return s.CheckFieldExist(dao.Role.Columns().Name, name, id...)
}

func (s *role) Store(input model.RoleStoreInput) (err error) {

	if exist, err := s.CheckNameExist(input.Name); err != nil {
		return err
	} else if exist {
		return errors.New("角色名称已存在")
	}

	return s.model.Transaction(s.ctx, func(ctx context.Context, tx *gdb.TX) error {

		insertId, err := s.model.Data(input).InsertAndGetId()
		if err != nil {
			return err
		}

		return RoleApi(ctx).Save(model.RoleApiSaveInput{
			RoleId: int(insertId),
			Apis:   input.Apis,
		})

	})
}

func (s *role) Update(input model.RoleUpdateInput) (err error) {

	if exist, err := s.CheckIdExist(input.Id); err != nil {
		return err
	} else if !exist {
		return errors.New("ID不存在")
	}

	if exist, err := s.CheckNameExist(input.Name, input.Id); err != nil {
		return err
	} else if exist {
		return errors.New("角色名称已存在")
	}

	return s.model.Transaction(s.ctx, func(ctx context.Context, tx *gdb.TX) error {

		if _, err = s.model.FieldsEx(dao.Role.Columns().Id).Where(dao.Role.Columns().Id, input.Id).Data(input).Update(); err != nil {
			return err
		}

		return RoleApi(ctx).Save(model.RoleApiSaveInput{
			RoleId: input.Id,
			Apis:   input.Apis,
		})

	})
}

func (s *role) PkDelete(id int) error {

	if id == g.Cfg("auth").MustGet(s.ctx, "superRoleId").Int() {
		return errors.New("超级管理员不能删除")
	}

	if exist, err := s.CheckIdExist(id); err != nil {
		return err
	} else if !exist {
		return errors.New("ID不存在")
	}

	return s.model.Transaction(s.ctx, func(ctx context.Context, tx *gdb.TX) error {

		if _, err := s.model.Where(dao.Role.Columns().Id, id).Delete(); err != nil {
			return err
		}

		return RoleApi(ctx).Delete(id)
	})

}

func (s *role) Info(id int, columns ...string) (output *entity.Role, err error) {

	output = &entity.Role{}

	if err = s.model.Where(dao.Role.Columns().Id, id).Fields(columns).Scan(output); err != nil {
		return nil, err
	}

	return output, nil
}

func (s *role) Show(id int, columns ...string) (output *model.RoleShowOutput, err error) {

	info, err := s.Info(id, columns...)
	if err != nil {
		return nil, err
	}

	apis, err := RoleApi(s.ctx).GetApis(id)
	if err != nil {
		return nil, err
	}

	output = &model.RoleShowOutput{
		Role: entity.Role{
			Id:        info.Id,
			Name:      info.Name,
			CreatedAt: info.CreatedAt,
			UpdatedAt: info.UpdatedAt,
		},
		Apis: apis,
	}

	return
}
