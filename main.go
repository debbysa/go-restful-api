package main

import (
	"github.com/debbysa/go-restful-api/app"
	"github.com/debbysa/go-restful-api/controller"
	"github.com/debbysa/go-restful-api/helper"
	"github.com/debbysa/go-restful-api/middleware"
	"github.com/debbysa/go-restful-api/repository"
	"github.com/debbysa/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	// http server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
