package model

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategoryRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
