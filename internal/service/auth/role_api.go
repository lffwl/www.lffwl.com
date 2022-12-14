package auth

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
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

func (s *roleApi) Delete(roleId int) error {
	if _, err := s.model.Where(dao.RoleApi.Columns().RoleId, roleId).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *roleApi) Save(input model.RoleApiSaveInput) error {

	return s.model.Transaction(s.ctx, func(ctx context.Context, tx *gdb.TX) error {

		if err := s.Delete(input.RoleId); err != nil {
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

	apis, err := s.model.Where(dao.RoleApi.Columns().RoleId, roleId).Distinct().Fields(dao.RoleApi.Columns().ApiId).Array()
	if err != nil {
		return nil, err
	}

	return gconv.Ints(apis), nil
}

func (s *roleApi) GetApisByRoles(roleId []int) (output []int, err error) {

	apis, err := s.model.WhereIn(dao.RoleApi.Columns().RoleId, roleId).Distinct().Fields(dao.RoleApi.Columns().ApiId).Array()
	if err != nil {
		return nil, err
	}

	return gconv.Ints(apis), nil
}
