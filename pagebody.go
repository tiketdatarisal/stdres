package stdres

type pageBody struct {
	PageNumber int64       `json:"pageNum,omitempty"`
	PageSize   int64       `json:"pageSize,omitempty"`
	PageTotal  int64       `json:"pageTotal,omitempty"`
	Total      int64       `json:"total"`
	Data       interface{} `json:"data"`
}

// GetLimitOffset returns limit and offset value for this params.
// These limit and offset will be used by SQL query to limit data retrieved from database.
func GetLimitOffset(pageNumber, pageSize, total int64) (limit, offset int) {
	pageNumber, pageSize, _, total = normalizeParam(pageNumber, pageSize, total)
	return int(pageSize), int((pageNumber - 1) * pageSize)
}

// NewPageBody return a new paged data body.
// This function is usually used in combination with default response body.
func NewPageBody(pageData interface{}, pageNumber, pageSize, total int64) interface{} {
	pageNumber, pageSize, pageTotal, total := normalizeParam(pageNumber, pageSize, total)
	return pageBody{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		PageTotal:  pageTotal,
		Total:      total,
		Data:       pageData,
	}
}
