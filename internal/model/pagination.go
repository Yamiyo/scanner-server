package model

type PaginationResp struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	PageCount int `json:"pageCount"`
	Total     int `json:"total"`
}

type Pagination struct {
	PageIndex int    `form:"pageIndex" binding:"required"`
	PageSize  int    `form:"pageSize" binding:"required"`
	PageOrder string `form:"pageOrder"`
}

func (p *Pagination) TransferToPages(totalCount int) *PaginationResp {
	pageCount := totalCount / p.PageSize
	if totalCount%p.PageSize != 0 {
		pageCount += 1
	}

	return &PaginationResp{
		PageIndex: p.PageIndex,
		PageSize:  p.PageSize,
		PageCount: pageCount,
		Total:     totalCount,
	}
}
