package web

type CategoryUpdateRequest struct {
	Id   uint64 `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=1,max=255"`
}
