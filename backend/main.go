package main

import (
	"backend/config"
	"backend/infra"
	"backend/interface/handler"
	"backend/usecase"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	termRepository := infra.NewTermRepository(config.OpenDBConn())
	termUsecase := usecase.NewTermUsecase(termRepository)
	termHandler := handler.NewTermHandler(termUsecase)

	categoryRepository := infra.NewCategoryRepository(config.OpenDBConn())
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	e := echo.New()
	e.Use(middleware.CORS())
	handler.InitRouting(e, categoryHandler)

	// heroku用
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = os.Getenv("API_PORT")
	}
	e.Logger.Fatal(e.Start(":" + port))
}
