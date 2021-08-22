package main

import (
	"tiger_app/backend_src/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	initRouting(e)
	e.Logger.Fatal(e.Start(":8000"))
}

func initRouting(e *echo.Echo) {
	e.GET("/api/terms", handlers.GetTerms)
	e.GET("/api/categories", handlers.GetCategories)

	e.POST("/api/signup", handlers.UserSignup)
	e.POST("/api/login", handlers.UserLogin)

	a := e.Group("/auth")
	a.Use(middleware.JWT([]byte("secret"))) // /auth下はJWTの認証が必要
	a.GET("/api/mypage", handlers.GetMypage)
}
