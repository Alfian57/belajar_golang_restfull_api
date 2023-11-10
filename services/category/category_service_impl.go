package services

import (
	"context"
	"database/sql"

	"github.com/Alfian57/belajar_golang_restfull_api/exceptions"
	"github.com/Alfian57/belajar_golang_restfull_api/helper"
	"github.com/Alfian57/belajar_golang_restfull_api/models/domain"
	"github.com/Alfian57/belajar_golang_restfull_api/models/web"
	repositories "github.com/Alfian57/belajar_golang_restfull_api/repositories/category"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repositories.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryReponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryReponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId uint64) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId uint64) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryReponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryReponses(categories)
}
