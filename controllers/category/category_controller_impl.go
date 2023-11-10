package controllers

import (
	"net/http"
	"strconv"

	"github.com/Alfian57/belajar_golang_restfull_api/helper"
	"github.com/Alfian57/belajar_golang_restfull_api/models/web"
	services "github.com/Alfian57/belajar_golang_restfull_api/services/category"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		service: service,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.service.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebReponse{
		Code:   201,
		Status: "CREATED",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryResponse := controller.service.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebReponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	categoryIdNumber, err := strconv.ParseUint(categoryId, 10, 64)
	helper.PanicIfError(err)

	controller.service.Delete(request.Context(), categoryIdNumber)
	webResponse := web.WebReponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	categoryIdNumber, err := strconv.ParseUint(categoryId, 10, 64)
	helper.PanicIfError(err)

	categoryResponse := controller.service.FindById(request.Context(), categoryIdNumber)
	webResponse := web.WebReponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controller.service.FindAll(request.Context())

	webResponse := web.WebReponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
