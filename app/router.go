package app

import (
	controllers "github.com/Alfian57/belajar_golang_restfull_api/controllers/category"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controllers.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:cateogoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:cateogoryId", categoryController.Update)
	router.DELETE("/api/categories/:cateogoryId", categoryController.Delete)

	return router
}
