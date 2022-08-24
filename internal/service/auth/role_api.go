package auth

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"www.lffwl.com/internal/dao"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/model/entity"
)

type roleApi struct {
	model *gdb.Model
	ctx   context.Context
}

func RoleApi(ctx context.Context) *roleApi {
	return &roleApi{
		model: dao.RoleApi.Ctx(ctx),
		ctx:   ctx,
	}
}

func (s *roleApi) Save(input model.RoleApiSaveInput) error {

	return s.model.Transaction(s.ctx, func(ctx context.Context, tx *gdb.TX) error {

		if _, err := s.model.Where(dao.RoleApi.Columns().RoleId, input.RoleId).Delete(); err != nil {
			return err
		}

		if len(input.Apis) > 0 {
			var data []entity.RoleApi
			for _, val := range input.Apis {
				data = append(data, entity.RoleApi{
					RoleId: input.RoleId,
					ApiId:  val,
				})
			}

			if _, err := s.model.Data(data).Insert(); err != nil {
				return err
			}
		}

		return nil
	})

}

func (s *roleApi) GetApis(roleId int) (output []int, err error) {

	if err = s.model.Where(dao.RoleApi.Columns().RoleId, roleId).Fields(dao.RoleApi.Columns().ApiId).Scan(&output); err != nil {
		return nil, err
	}

	return
}
