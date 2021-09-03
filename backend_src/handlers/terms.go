package handlers

import (
	"net/http"
	"strconv"

	"tiger_app/backend_src/models"

	"github.com/labstack/echo"
)

func GetTerms(c echo.Context) error {
	var terms []models.Terms
	Db := models.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	Db.Order("name").Find(&terms)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, terms)
}

func GetTerm(c echo.Context) error {
	term := new(models.Terms)
	Db := models.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	term_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	Db.Where("id = ?", term_id).First(&term)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, term)
}
