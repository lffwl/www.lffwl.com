package manage

type CommonPageReq struct {
	Page int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size int `json:"size" in:"query" d:"10" v:"max:50#分页数量最大50条" dc:"分页大小，最大50"`
}

type CommonPageRes struct {
	Page  int `json:"page" dc:"当前页"`
	Size  int `json:"size" dc:"分页大小"`
	Total int `json:"total" dc:"总数"`
}
