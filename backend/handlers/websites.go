package handlers

import (
	"net/http"
	"strconv"

	"backend/domain/model"

	"github.com/labstack/echo"
)

func GetWebsites(c echo.Context) error {
	var websites []model.Website
	Db := model.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	Db.Order("name").Find(&websites)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, websites)
}

func GetWebsitesByCategory(c echo.Context) error {
	var websites []model.Website
	Db := model.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	category_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	type Results struct {
		ID int
	}
	var results []Results
	Db.Table("categories_websites").Where("category_id = ?", category_id).Scan(&results)
	if err != nil {
		return err
	} else if len(results) == 0 {
		return c.JSON(http.StatusOK, websites)
	}

	var websites_by_category_ids []int

	for i := 0; i < len(results); i++ {
		websites_by_category_ids = append(websites_by_category_ids, results[i].ID)
	}

	Db.Where(websites_by_category_ids).Find(&websites)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, websites)
}

func GetWebsite(c echo.Context) error {
	website := new(model.Website)
	Db := model.OpenDBConn()
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

func SearchWebsite(c echo.Context) error {
	var websites []model.Website
	Db := model.OpenDBConn()
	db, err := Db.DB()
	defer db.Close()

	serch_query := c.QueryParam("q")

	Db.Where("name ILIKE ?", "%"+serch_query+"%").Find(&websites)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, websites)
}
