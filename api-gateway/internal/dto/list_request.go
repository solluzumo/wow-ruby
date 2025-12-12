package dto

//TODO: реализовать метод базового репозитория для пополучения списка объектов

type ListRequest struct {
	Page     int `json:"page" form:"page" query:"page"`
	PageSize int `json:"page_size" form:"page_size" query:"page_size"`

	SortBy    string `json:"sort_by" form:"sort_by" query:"sort_by"`
	SortOrder string `json:"sort_order" form:"sort_order" query:"sort_order"` // "asc" или "desc"

	Filters map[string]interface{} `json:"filters" form:"filters" query:"filters"`

	Search string `json:"search" form:"search" query:"search"`
}
