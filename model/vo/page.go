package vo

type ListQuery struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

func (this ListQuery) Offset() int {
	if this.PageNum == 0 {
		this.PageNum = 1
		this.PageSize = 10
	}
	if this.PageNum > 0 {
		return (this.PageNum - 1) * this.PageSize
	}
	return this.PageSize
}

type PageQueryReply struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}
