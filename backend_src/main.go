package main

import (
	"tiger_app/backend_src/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	initRouting(e)
	e.Logger.Fatal(e.Start(":8000"))
}

func initRouting(e *echo.Echo) {
	e.GET("/terms", handlers.GetTerms)
	e.GET("/categories", handlers.GetCategories)

	// 　todo認証
	// e.POST("/login", handlers.Login)
}
