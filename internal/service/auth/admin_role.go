package auth

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"www.lffwl.com/internal/dao"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/model/entity"
)

type adminRole struct {
	model *gdb.Model
	ctx   context.Context
}

func AdminRole(ctx context.Context) *adminRole {
	return &adminRole{
		model: dao.AdminRole.Ctx(ctx),
		ctx:   ctx,
	}
}

func (s *adminRole) Save(input model.AdminRoleSaveInput) error {

	return s.model.Transaction(s.ctx, func(ctx context.Context, tx *gdb.TX) error {

		if _, err := s.model.Where(dao.AdminRole.Columns().AdminId, input.AdminId).Delete(); err != nil {
			return err
		}

		if len(input.Roles) > 0 {
			var data []entity.AdminRole
			for _, val := range input.Roles {
				data = append(data, entity.AdminRole{
					AdminId: input.AdminId,
					RoleId:  val,
				})
			}

			if _, err := s.model.Data(data).Insert(); err != nil {
				return err
			}
		}

		return nil
	})

}

func (s *adminRole) GetRoles(adminId int) (output []int, err error) {

	if err = s.model.Where(dao.AdminRole.Columns().AdminId, adminId).Fields(dao.AdminRole.Columns().RoleId).Scan(&output); err != nil {
		return nil, err
	}

	return
}
