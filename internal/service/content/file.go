package content

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
	"time"
	"www.lffwl.com/internal/model"
)

type file struct {
}

func File() *file {
	return &file{}
}

func (s *file) Upload(ctx context.Context, input model.FileUploadInput) (output *model.FileUploadOutput, err error) {

	fileDir := g.Cfg("content").MustGet(ctx, "file.articlePath").String() + time.Now().Format("/2006-01-02/")
	fileName, err := input.File.Save(fileDir, true)
	if err != nil {
		return nil, err
	}

	output = &model.FileUploadOutput{File: s.ToUrl(fileDir + fileName), Name: input.File.Filename}

	return
}

func (s *file) ToUrl(filePath string) string {
	return strings.TrimLeft(filePath, ".")
}
