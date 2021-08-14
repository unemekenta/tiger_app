package handlers

import (
	"net/http"

	"tiger_app/backend_src/models"

	"github.com/labstack/echo"
)

func GetTerms(c echo.Context) error {
	var terms = new(models.Terms)
	Db := models.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	Db.Order("id").Find(terms)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, terms)
}
