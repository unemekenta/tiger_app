package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/unemekenta/tiger_app/backend/usecase"
)

// WebsiteContentHandler websiteContent handlerのinterface
type WebsiteContentHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetByWebsiteID() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type websiteContentHandler struct {
	websiteContentUsecase usecase.WebsiteContentUsecase
}

// WebsiteContentHandler websiteContent handlerのコンストラクタ
func NewWebsiteContentHandler(websiteContentUsecase usecase.WebsiteContentUsecase) WebsiteContentHandler {
	return &websiteContentHandler{websiteContentUsecase: websiteContentUsecase}
}

type requestWebsiteContent struct {
	WebsiteID int       `json:"websiteId"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type responseWebsiteContent struct {
	ID        int    `json:"id"`
	WebsiteID int    `json:"websiteId"`
	Title     string `json:"title"`
	Contents  string `json:"contents"`
}

// Post websiteContentを保存するときのハンドラー
func (wch *websiteContentHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestWebsiteContent
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdWebsiteContent, err := wch.websiteContentUsecase.Create(req.WebsiteID, req.Title, req.Contents, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseWebsiteContent{
			ID:        createdWebsiteContent.ID,
			WebsiteID: createdWebsiteContent.WebsiteID,
			Title:     createdWebsiteContent.Title,
			Contents:  createdWebsiteContent.Contents,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get websiteContentを取得するときのハンドラー
func (wch *websiteContentHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundWebsiteContent, err := wch.websiteContentUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseWebsiteContent{
			ID:        foundWebsiteContent.ID,
			WebsiteID: foundWebsiteContent.WebsiteID,
			Title:     foundWebsiteContent.Title,
			Contents:  foundWebsiteContent.Contents,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetByWebsiteID websiteContentをWebsiteIDで取得するときのハンドラー
func (wch *websiteContentHandler) GetByWebsiteID() echo.HandlerFunc {
	return func(c echo.Context) error {
		websiteID, err := strconv.Atoi((c.Param("website_id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundWebsiteContent, err := wch.websiteContentUsecase.FindByWebsiteID(websiteID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseWebsiteContent{
			ID:        foundWebsiteContent.ID,
			WebsiteID: foundWebsiteContent.WebsiteID,
			Title:     foundWebsiteContent.Title,
			Contents:  foundWebsiteContent.Contents,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put websiteContentを更新するときのハンドラー
func (wch *websiteContentHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestWebsiteContent
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedWebsiteContent, err := wch.websiteContentUsecase.Update(id, req.Title, req.Contents, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseWebsiteContent{
			ID:        updatedWebsiteContent.ID,
			WebsiteID: updatedWebsiteContent.WebsiteID,
			Title:     updatedWebsiteContent.Title,
			Contents:  updatedWebsiteContent.Contents,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete websiteContentを削除するときのハンドラー
func (wch *websiteContentHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = wch.websiteContentUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
