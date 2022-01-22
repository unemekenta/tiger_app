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

	websiteRepository := infra.NewWebsiteRepository(config.OpenDBConn())
	websiteUsecase := usecase.NewWebsiteUsecase(websiteRepository)
	websiteHandler := handler.NewWebsiteHandler(websiteUsecase)

	userRepository := infra.NewUserRepository(config.OpenDBConn())
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	websiteContentRepository := infra.NewWebsiteContentRepository(config.OpenDBConn())
	websiteContentUsecase := usecase.NewWebsiteContentUsecase(websiteContentRepository)
	websiteContentHandler := handler.NewWebsiteContentHandler(websiteContentUsecase)

	e := echo.New()
	e.Use(middleware.CORS())
	handler.InitRouting(e, termHandler, categoryHandler, websiteHandler, userHandler, websiteContentHandler)

	// heroku用
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = os.Getenv("API_PORT")
	}
	e.Logger.Fatal(e.Start(":" + port))
}
