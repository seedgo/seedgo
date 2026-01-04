package seedgo

type Page struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

func (v *Page) FillPageDefault() {
	if v.PageSize < 2 {
		v.PageSize = 10
	}
	if v.PageSize > 1000 {
		v.PageSize = 1000
	}

	if v.PageNum < 1 {
		v.PageNum = 1
	}
}
