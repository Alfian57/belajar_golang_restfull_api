package services

import (
	"context"

	"github.com/Alfian57/belajar_golang_restfull_api/models/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId uint64)
	FindById(ctx context.Context, categoryId uint64) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
