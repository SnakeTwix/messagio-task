package entity

type PaginateResponse[T any] struct {
	Values []T `json:"values"`
	Total  int `json:"total"`
}

type PaginateRequest struct {
	Limit int `query:"limit" validate:"lte=200,required"`
	Page  int `query:"page" validate:"gte=0"`
}

type Paginate struct {
	Limit  int
	Offset int
}
