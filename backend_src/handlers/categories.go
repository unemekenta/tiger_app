package handlers

import (
	"net/http"

	"tiger_app/backend_src/models"

	"github.com/labstack/echo"
)

func GetCategories(c echo.Context) error {
	var categories []models.Categories
	Db := models.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	Db.Order("id").Find(&categories)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, categories)
}
