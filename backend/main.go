package main

import (
	"backend/handlers"

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
	e.GET("/api/terms/:id", handlers.GetTerm)

	e.GET("/api/websites", handlers.GetWebsites)
	e.GET("/api/websites/:id", handlers.GetWebsite)
	e.GET("/api/websites_category/:id", handlers.GetWebsitesByCategory)
	e.GET("/api/websites/search", handlers.SearchWebsite)

	e.GET("/api/website_content/:website_id", handlers.GetContentByWebsite)

	e.GET("/api/categories", handlers.GetCategories)
	e.GET("/api/categories_by_ancestor/:ancestor_id", handlers.GetCategoriesByAncestor)

	e.GET("/api/categories_website/:id", handlers.GetCategoriesByWebsites)

	e.POST("/api/signup", handlers.UserSignup)
	e.POST("/api/login", handlers.UserLogin)

	a := e.Group("/auth")
	a.Use(middleware.JWT([]byte("secret")))
	// /auth下はJWTの認証が必要
	// curlで接続する場合は curl http://localhost:8000/auth/api/mypage -H "Authorization: Bearer {login時に発行されるtoken}"
	a.POST("/api/authCheck", handlers.GetAuthCheck)

	a.GET("/api/user/:id", handlers.GetUser)
}
