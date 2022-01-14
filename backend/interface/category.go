package handler

import (
	"net/http"
	"strconv"
	"time"

	"backend/usecase"

	"github.com/labstack/echo"
)

// CategoryHandler category handlerのinterface
type CategoryHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type categoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

// CategoryHandler category handlerのコンストラクタ
func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) CategoryHandler {
	return &categoryHandler{categoryUsecase: categoryUsecase}
}

type requestCategory struct {
	AncestorID int       `json:"ancestor_id"`
	Name       string    `json:"name"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type responseCategory struct {
	ID         int    `json:"id"`
	AncestorID int    `json:"ancestor_id"`
	Name       string `json:"name"`
}

// Post categoryを保存するときのハンドラー
func (ch *categoryHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestCategory
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdCategory, err := ch.categoryUsecase.Create(req.AncestorID, req.Name, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseCategory{
			ID:         createdCategory.ID,
			AncestorID: createdCategory.AncestorID,
			Name:       createdCategory.Name,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get categoryを取得するときのハンドラー
func (ch *categoryHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundCategory, err := ch.categoryUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseCategory{
			ID:         foundCategory.ID,
			AncestorID: foundCategory.AncestorID,
			Name:       foundCategory.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetAllChildCategories 末端categoryを全て取得するときのハンドラー
func (ch *categoryHandler) GetAllChildCategories() echo.HandlerFunc {
	return func(c echo.Context) error {

		foundCategories, err := ch.categoryUsecase.FindAllChildCategories()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseCategory{
			ID:         foundCategories.ID,
			AncestorID: foundCategories.AncestorID,
			Name:       foundCategories.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetCategoriesByAncestor 親カテゴリから子カテゴリを取得するときのハンドラー
func (ch *categoryHandler) GetCategoriesByAncestor() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundCategories, err := ch.categoryUsecase.GetCategoriesByAncestor(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseCategory{
			ID:         foundCategories.ID,
			AncestorID: foundCategories.AncestorID,
			Name:       foundCategories.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put categoryを更新するときのハンドラー
func (ch *categoryHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestCategory
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedCategory, err := ch.categoryUsecase.Update(id, req.AncestorID, req.Name, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseCategory{
			ID:         updatedCategory.ID,
			AncestorID: updatedCategory.AncestorID,
			Name:       updatedCategory.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete categoryを削除するときのハンドラー
func (ch *categoryHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = ch.categoryUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
