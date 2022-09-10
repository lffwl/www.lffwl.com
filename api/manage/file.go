package manage

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file/upload" method:"post" mime:"multipart/form-data" tags:"File" summary:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}
type FileUploadRes struct {
	Url  string `json:"url"   dc:"访问URL，可能只是URI"`
	Name string `json:"name"  dc:"文件名"`
}
