package utils

import "math"

type PaginationQueryString struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type Pagination struct {
	Limit  int
	Offset int
	Order  string
	Page   int
}

type PaginationResult struct {
	DataCurrentPage int `json:"data_current_page"`
	DataLimit       int `json:"data_limit"`
	DataTotal       int `json:"data_total"`
	PageCurrent     int `json:"page_current"`
	PageTotal       int `json:"page_total"`
}

func NewPagination(req *PaginationQueryString) *Pagination {
	page := req.Page
	if page <= 0 {
		page = 1
	}

	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	if offset <= 0 {
		offset = 0
	}

	return &Pagination{
		Limit:  limit,
		Offset: offset,
		Order:  req.Sort,
		Page:   req.Page,
	}
}

func GeneratePagination(len int, count int, pagination *Pagination) *PaginationResult {
	totalPages := int(math.Ceil(float64(count) / float64(pagination.Limit)))
	if totalPages <= 0 {
		totalPages = 1
	}

	return &PaginationResult{
		DataCurrentPage: len,
		DataLimit:       pagination.Limit,
		DataTotal:       count,
		PageCurrent:     pagination.Page,
		PageTotal:       totalPages,
	}
}
