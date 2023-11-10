package helper

import (
	"github.com/Alfian57/belajar_golang_restfull_api/models/domain"
	"github.com/Alfian57/belajar_golang_restfull_api/models/web"
)

func ToCategoryReponse(category domain.Category) web.CategoryResponse {
	categoryReponse := web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}

	return categoryReponse
}

func ToCategoryReponses(categories []domain.Category) []web.CategoryResponse {
	var categoryReponses []web.CategoryResponse
	for _, category := range categories {
		categoryReponses = append(categoryReponses, ToCategoryReponse(category))
	}

	return categoryReponses
}
