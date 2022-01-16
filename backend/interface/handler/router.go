package handler

import (
	"backend/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouting(e *echo.Echo, termHandler TermHandler, categoryHandler CategoryHandler) {

	api := e.Group("/api")
	api.GET("/terms", termHandler.GetAll())
	api.GET("/term/:id", termHandler.Get())

	api.GET("/websites", handlers.GetWebsites)
	api.GET("/websites/:id", handlers.GetWebsite)
	api.GET("/websites_category/:id", handlers.GetWebsitesByCategory)
	api.GET("/websites/search", handlers.SearchWebsite)

	api.GET("/website_content/:website_id", handlers.GetContentByWebsite)

	api.GET("/categories", categoryHandler.GetAllParentCategories())
	api.GET("/categories_by_ancestor/:id", categoryHandler.GetCategoriesByAncestor())
	api.GET("/categories_website/:id", categoryHandler.GetCategoriesByAncestor())

	api.POST("/signup", handlers.UserSignup)
	api.POST("/login", handlers.UserLogin)
	api.GET("/login_check", handlers.UserLoginCheck)

	a := e.Group("/auth")
	a.Use(middleware.JWT([]byte("secret")))
	// /auth下はJWTの認証が必要
	// curlで接続する場合は curl http://localhost:8000/auth/api/mypage -H "Authorization: Bearer {login時に発行されるtoken}"
	a.POST("/api/authCheck", handlers.GetAuthCheck)

	a.GET("/api/user/:id", handlers.GetUser)
}
