package handler

import (
	"net/http"
	"strconv"
	"time"

	"backend/usecase"

	"github.com/labstack/echo"
)

// WebsiteHandler website handlerのinterface
type WebsiteHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetByCategory() echo.HandlerFunc
	SearchByName() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type websiteHandler struct {
	websiteUsecase usecase.WebsiteUsecase
}

// WebsiteHandler website handlerのコンストラクタ
func NewWebsiteHandler(websiteUsecase usecase.WebsiteUsecase) WebsiteHandler {
	return &websiteHandler{websiteUsecase: websiteUsecase}
}

type requestWebsite struct {
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	CompanyName string    `json:"companyName"`
	Image       string    `json:"image"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type responseWebsite struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	CompanyName string    `json:"companyName"`
	Image       string    `json:"image"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Post websiteを保存するときのハンドラー
func (wh *websiteHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestWebsite
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdWebsite, err := wh.websiteUsecase.Create(req.Name, req.Url, req.CompanyName, req.Image, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseWebsite{
			ID:          createdWebsite.ID,
			Name:        createdWebsite.Name,
			Url:         createdWebsite.Url,
			CompanyName: createdWebsite.CompanyName,
			Image:       createdWebsite.Image,
			UpdatedAt:   createdWebsite.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get websiteを取得するときのハンドラー
func (wh *websiteHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundWebsite, err := wh.websiteUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseWebsite{
			ID:          foundWebsite.ID,
			Name:        foundWebsite.Name,
			Url:         foundWebsite.Url,
			CompanyName: foundWebsite.CompanyName,
			Image:       foundWebsite.Image,
			UpdatedAt:   foundWebsite.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetAll websiteを全て取得するときのハンドラー
func (wh *websiteHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundWebsites, err := wh.websiteUsecase.FindAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseWebsite

		for _, fw := range *foundWebsites {
			res = append(res, responseWebsite{
				ID:          fw.ID,
				Name:        fw.Name,
				Url:         fw.Url,
				CompanyName: fw.CompanyName,
				Image:       fw.Image,
				UpdatedAt:   fw.UpdatedAt,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetByCategory カテゴリからwebsiteを全て取得するときのハンドラー
func (wh *websiteHandler) GetByCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		foundWebsites, err := wh.websiteUsecase.FindByCategory(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseWebsite

		for _, fw := range *foundWebsites {
			res = append(res, responseWebsite{
				ID:          fw.ID,
				Name:        fw.Name,
				Url:         fw.Url,
				CompanyName: fw.CompanyName,
				Image:       fw.Image,
				UpdatedAt:   fw.UpdatedAt,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

// SearchByName nameでの検索結果を全て取得するときのハンドラー
func (wh *websiteHandler) SearchByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		query := c.QueryParam("q")
		foundWebsites, err := wh.websiteUsecase.SearchByName(query)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseWebsite

		for _, fw := range *foundWebsites {
			res = append(res, responseWebsite{
				ID:          fw.ID,
				Name:        fw.Name,
				Url:         fw.Url,
				CompanyName: fw.CompanyName,
				Image:       fw.Image,
				UpdatedAt:   fw.UpdatedAt,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put websiteを更新するときのハンドラー
func (wh *websiteHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestWebsite
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedWebsite, err := wh.websiteUsecase.Update(id, req.Name, req.Url, req.CompanyName, req.Image, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseWebsite{
			ID:          updatedWebsite.ID,
			Name:        updatedWebsite.Name,
			Url:         updatedWebsite.Url,
			CompanyName: updatedWebsite.CompanyName,
			Image:       updatedWebsite.Image,
			UpdatedAt:   updatedWebsite.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete websiteを削除するときのハンドラー
func (wh *websiteHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = wh.websiteUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
