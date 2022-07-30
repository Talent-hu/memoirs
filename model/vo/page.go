package vo

type BasePage struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

func (this BasePage) Offset() int {
	if this.PageNum == 0 {
		this.PageNum = 1
		this.PageSize = 10
	}
	if this.PageNum > 0 {
		return (this.PageNum - 1) * this.PageSize
	}
	return this.PageSize
}
