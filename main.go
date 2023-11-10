package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Alfian57/belajar_golang_restfull_api/app"
	controllers "github.com/Alfian57/belajar_golang_restfull_api/controllers/category"
	"github.com/Alfian57/belajar_golang_restfull_api/helper"
	repositories "github.com/Alfian57/belajar_golang_restfull_api/repositories/category"
	services "github.com/Alfian57/belajar_golang_restfull_api/services/category"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controllers.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	// router.PanicHandler = exceptions.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
