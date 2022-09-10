package manage

import (
	"context"
	"www.lffwl.com/internal/model"
	"www.lffwl.com/internal/service/content"

	"www.lffwl.com/api/manage"
)

var (
	File = cFile{}
)

type cFile struct{}

func (c *cFile) Upload(ctx context.Context, req *manage.FileUploadReq) (res *manage.FileUploadRes, err error) {

	data, err := content.File().Upload(ctx, model.FileUploadInput{File: req.File})
	if err != nil {
		return nil, err
	}

	res = &manage.FileUploadRes{Url: data.File, Name: data.Name}

	return
}
