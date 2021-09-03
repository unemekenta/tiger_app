package handlers

import (
	"net/http"
	"strconv"

	"tiger_app/backend_src/models"

	"github.com/labstack/echo"
)

func GetWebsites(c echo.Context) error {
	var websites []models.Websites
	Db := models.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	Db.Order("name").Find(&websites)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, websites)
}

func GetWebsite(c echo.Context) error {
	website := new(models.Websites)
	Db := models.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	website_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	Db.Where("id = ?", website_id).First(&website)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, website)
}
