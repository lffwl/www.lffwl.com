package auth

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"www.lffwl.com/internal/dao"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/model/entity"
)

type admin struct {
	model *gdb.Model
	ctx   context.Context
}

func Admin(ctx context.Context) *admin {
	return &admin{
		model: dao.Admin.Ctx(ctx),
		ctx:   ctx,
	}
}

func (s *admin) Index(input model.AdminIndexInput) (output *model.AdminIndexOutput, err error) {

	output = &model.AdminIndexOutput{}
	m := s.model

	if input.Username != "" {
		m = m.WhereLike(dao.Admin.Columns().Username, "%"+input.Username+"%")
	}

	if err = m.Fields(model.AdminIndexOutputList{}).Page(input.Page, input.Limit).Scan(&output.List); err != nil {
		return nil, err
	}

	if output.Total, err = m.Count(); err != nil {
		return nil, err
	}

	return

}

func (s *admin) CheckFieldExist(field string, value interface{}, id ...int) (bool, error) {
	m := s.model

	if len(id) > 0 {
		m = m.WhereNot(dao.Admin.Columns().Id, id[0])
	}

	if num, err := m.Where(field, value).Count(); err != nil {
		return false, err
	} else if num > 0 {
		return true, nil
	}

	return false, nil
}

func (s *admin) CheckIdExist(id int) (bool, error) {
	return s.CheckFieldExist(dao.Admin.Columns().Id, id)
}

func (s *admin) CheckUsernameExist(name string, id ...int) (bool, error) {
	return s.CheckFieldExist(dao.Admin.Columns().Username, name, id...)
}

func (s *admin) Store(input model.AdminStoreInput) (err error) {

	if exist, err := s.CheckUsernameExist(input.Username); err != nil {
		return err
	} else if exist {
		return errors.New("管理员名称已存在")
	}

	return s.model.Transaction(s.ctx, func(ctx context.Context, tx *gdb.TX) error {

		insertId, err := s.model.Data(input).InsertAndGetId()
		if err != nil {
			return err
		}

		return AdminRole(ctx).Save(model.AdminRoleSaveInput{
			AdminId: int(insertId),
			Roles:   input.Roles,
		})

	})
}

func (s *admin) Update(input model.AdminUpdateInput) (err error) {

	if exist, err := s.CheckIdExist(input.Id); err != nil {
		return err
	} else if !exist {
		return errors.New("ID不存在")
	}

	if exist, err := s.CheckUsernameExist(input.Username, input.Id); err != nil {
		return err
	} else if exist {
		return errors.New("管理员名称已存在")
	}

	return s.model.Transaction(s.ctx, func(ctx context.Context, tx *gdb.TX) error {

		if _, err = s.model.Data(input).Update(); err != nil {
			return err
		}

		return AdminRole(ctx).Save(model.AdminRoleSaveInput{
			AdminId: input.Id,
			Roles:   input.Roles,
		})

	})
}

func (s *admin) PkDelete(id int) error {

	if exist, err := s.CheckIdExist(id); err != nil {
		return err
	} else if !exist {
		return errors.New("ID不存在")
	}

	return s.model.Transaction(s.ctx, func(ctx context.Context, tx *gdb.TX) error {

		if _, err := s.model.Where(dao.Admin.Columns().Id, id).Delete(); err != nil {
			return err
		}

		return AdminRole(ctx).Delete(id)
	})
}

func (s *admin) Info(id int, columnsEx ...string) (output *entity.Admin, err error) {

	output = &entity.Admin{}

	if err = s.model.Where(dao.Admin.Columns().Id, id).FieldsEx(columnsEx).Scan(output); err != nil {
		return nil, err
	}

	return output, nil
}

func (s *admin) Show(id int) (output *model.AdminShowOutput, err error) {

	info, err := s.Info(id, dao.Admin.Columns().Password)
	if err != nil {
		return nil, err
	}

	roles, err := AdminRole(s.ctx).GetRoles(id)
	if err != nil {
		return nil, err
	}

	output = &model.AdminShowOutput{
		Id:        info.Id,
		Username:  info.Username,
		CreatedAt: info.CreatedAt,
		UpdatedAt: info.UpdatedAt,
		Roles:     roles,
	}

	return
}
