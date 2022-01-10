package handlers

import (
	"net/http"
	"strconv"

	"backend/model"

	"github.com/labstack/echo"
)

func GetTerms(c echo.Context) error {
	var terms []model.Terms
	Db := model.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	Db.Order("name").Find(&terms)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, terms)
}

func GetTerm(c echo.Context) error {
	term := new(model.Terms)
	Db := model.OpenDBConn()
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
