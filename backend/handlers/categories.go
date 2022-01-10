package handlers

import (
	"net/http"
	"strconv"

	"backend/model"

	"github.com/labstack/echo"
)

func GetCategories(c echo.Context) error {
	var categories []model.Categories
	Db := model.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	Db.Order("name").Where("ancestor_id IS NULL").Find(&categories)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, categories)
}

func GetCategoriesByAncestor(c echo.Context) error {
	var categories []model.Categories
	Db := model.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	ancestor_id, err := strconv.Atoi(c.Param("ancestor_id"))
	if err != nil {
		return err
	}

	Db.Order("name").Where("ancestor_id = ?", ancestor_id).Find(&categories)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, categories)
}

func GetCategoriesByWebsites(c echo.Context) error {
	var category model.Categories
	var category_website model.CategoriesWebsites
	Db := model.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	website_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	Db.Where("website_id = ?", website_id).First(&category_website)
	if err != nil {
		return err
	}
	category_id := category_website.CategoryID

	Db.Where("id = ?", category_id).First(&category)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, category)
}
