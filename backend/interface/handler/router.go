package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouting(e *echo.Echo, termHandler TermHandler, categoryHandler CategoryHandler, websiteHandler WebsiteHandler, userHandler UserHandler, websiteContentHandler WebsiteContentHandler, moneyAccountHandler MoneyAccountHandler) {

	api := e.Group("/api")
	api.GET("/terms", termHandler.GetAll())
	api.GET("/term/:id", termHandler.Get())

	api.GET("/websites", websiteHandler.GetAll())
	api.GET("/website/:id", websiteHandler.Get())
	api.GET("/websites_category/:id", websiteHandler.GetByCategory())
	api.GET("/websites/search", websiteHandler.SearchByName())

	api.GET("/website_content/:website_id", websiteContentHandler.GetByWebsiteID())

	api.GET("/categories", categoryHandler.GetAllParentCategories())
	api.GET("/categories_by_ancestor/:id", categoryHandler.GetCategoriesByAncestor())
	api.GET("/categories_website/:id", categoryHandler.GetCategoryByWebsite())

	api.POST("/signup", userHandler.Signup())
	api.POST("/login", userHandler.Login())
	api.POST("/login_check", userHandler.LoginCheck())

	api.GET("/money_account/user/:id", moneyAccountHandler.GetByUser())
	api.POST("/money_account/user", moneyAccountHandler.Post())
	api.DELETE("/money_account/user/:id", moneyAccountHandler.Delete())

	auth := api.Group("/auth")
	auth.Use(middleware.JWT([]byte("secret")))
	// /auth下はJWTの認証が必要
	// curlで接続する場合は curl http://localhost:8000/api/auth/mypage -H "Authorization: Bearer {login時に発行されるtoken}"
	auth.POST("/authCheck", userHandler.AuthCheck())

	auth.GET("/user/:id", userHandler.Get())
}
