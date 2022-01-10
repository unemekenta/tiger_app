package handlers

import (
	"net/http"
	"strconv"

	"backend/model"

	"github.com/labstack/echo"
)

func GetContentByWebsite(c echo.Context) error {
	var website_contents model.WebsiteContents
	Db := model.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	website_id, err := strconv.Atoi(c.Param("website_id"))
	if err != nil {
		return err
	}

	Db.Where("website_id = ?", website_id).First(&website_contents)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, website_contents)
}
