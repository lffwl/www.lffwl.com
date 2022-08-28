package article

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lffwl/utility/file"
	"io/ioutil"
	"os"
	"time"
	"www.lffwl.com/internal/dao"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/model/entity"
)

type article struct {
	model *gdb.Model
	ctx   context.Context
}

func Article(ctx context.Context) *article {
	return &article{
		model: dao.Article.Ctx(ctx),
		ctx:   ctx,
	}
}

func (s *article) Index(input model.ArticleIndexInput) (output *model.ArticleIndexOutput, err error) {

	output = &model.ArticleIndexOutput{}
	m := s.model

	if input.Name != "" {
		m = m.WhereLike(dao.Article.Columns().Name, "%"+input.Name+"%")
	}

	if input.TypeId > 0 {
		m = m.Where(dao.Article.Columns().TypeId, input.TypeId)
	}

	if err = m.Page(input.Page, input.Limit).Scan(&output.List); err != nil {
		return nil, err
	}

	if output.Total, err = m.Count(); err != nil {
		return nil, err
	}

	for i, item := range output.List {
		if item.Content != "" {
			if bytes, err := ioutil.ReadFile(item.Content); err != nil {
				output.List[i].Content = err.Error()
			} else {
				output.List[i].Content = string(bytes[:])
			}
		}
	}

	return
}

func (s *article) CreateFilePath() string {
	now := time.Now()

	filePath := g.Cfg().MustGet(s.ctx, "file.articlePath").String() + now.Format("/2006-01-02/") + gconv.String(now.Second()) + gconv.String(now.Nanosecond()) + ".txt"

	if err := file.CreateFilePathDir(filePath); err != nil {
		g.Log().Error(s.ctx, err)
	}

	return filePath

}

func (s *article) SaveContent(content string, path ...string) (string, error) {

	var filePath string
	if len(path) > 0 && path[0] != "" {
		filePath = path[0]
	} else {
		filePath = s.CreateFilePath()
	}

	fileContent, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer fileContent.Close()
	fileContent.WriteString(content)

	return filePath, nil
}

func (s *article) Store(input model.ArticleStoreInput) (err error) {

	if input.Content != "" {
		if input.Content, err = s.SaveContent(input.Content); err != nil {
			return err
		}
	}

	if _, err = s.model.Data(input).Save(); err != nil {
		return err
	}

	return nil
}

func (s *article) CheckIdExist(id int) (bool, error) {

	if num, err := s.model.Where(dao.Article.Columns().Id, id).Count(); err != nil {
		return false, err
	} else if num > 0 {
		return true, nil
	}

	return false, nil
}

func (s *article) Info(id int, columns ...string) (output *entity.Article, err error) {

	output = &entity.Article{}

	if err = s.model.Where(dao.Article.Columns().Id, id).Fields(columns).Scan(output); err != nil {
		return nil, err
	}

	return output, nil
}

func (s *article) Update(input model.ArticleUpdateInput) (err error) {

	if exist, err := s.CheckIdExist(input.Id); err != nil {
		return err
	} else if !exist {
		return errors.New("ID不存在")
	}

	if input.Content != "" {
		info, err := s.Info(input.Id, dao.Article.Columns().Content)
		if err != nil {
			return err
		}
		if input.Content, err = s.SaveContent(input.Content, info.Content); err != nil {
			return err
		}
	}

	if _, err = s.model.Where(dao.Article.Columns().Id, input.Id).FieldsEx(dao.Article.Columns().Id).Data(input).Update(); err != nil {
		return err
	}

	return nil
}

func (s *article) PkDelete(id int) error {

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

func (s *article) Show(id int) (output *model.ArticleShowOutput, err error) {

	if exist, err := s.CheckIdExist(id); err != nil {
		return nil, err
	} else if !exist {
		return nil, errors.New("ID不存在")
	}

	info, err := s.Info(id)
	if err != nil {
		return nil, err
	}

	output = &model.ArticleShowOutput{
		Id:     info.Id,
		Name:   info.Name,
		TypeId: info.TypeId,
		Desc:   info.Desc,
	}

	if info.Content != "" {
		bytes, err := ioutil.ReadFile(info.Content)
		if err != nil {
			return nil, err
		}

		output.Content = string(bytes[:])
	}

	return
}
