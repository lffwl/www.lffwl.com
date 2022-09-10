package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUploadInput struct {
	File *ghttp.UploadFile
}
type FileUploadOutput struct {
	File string
	Name string
}
