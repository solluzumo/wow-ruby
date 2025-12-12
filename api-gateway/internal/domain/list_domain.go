package domain

type ListDomain struct {
	Page     int
	PageSize int

	SortBy    string
	SortOrder string

	Filters map[string]interface{}

	Search string
}

type ListResponseDomain[T any] struct {
	Data       []T
	Total      int
	Page       int
	PageSize   int
	TotalPages int
}
