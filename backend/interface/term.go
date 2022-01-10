package handler

import (
	"net/http"
	"strconv"
	"time"

	"backend/usecase"

	"github.com/labstack/echo"
)

// TermHandler term handlerのinterface
type TermHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type termHandler struct {
	termUsecase usecase.TermUsecase
}

// TermHandler term handlerのコンストラクタ
func NewTermHandler(termUsecase usecase.TermUsecase) TermHandler {
	return &termHandler{termUsecase: termUsecase}
}

type requestTerm struct {
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

type responseTerm struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// Post termを保存するときのハンドラー
func (th *termHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestTerm
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdTerm, err := th.termUsecase.Create(req.Name, req.Content, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTerm{
			ID:      createdTerm.ID,
			Name:    createdTerm.Name,
			Content: createdTerm.Content,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get termを取得するときのハンドラー
func (th *termHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundTerm, err := th.termUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTerm{
			ID:      foundTerm.ID,
			Name:    foundTerm.Name,
			Content: foundTerm.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put termを更新するときのハンドラー
func (th *termHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestTerm
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedTerm, err := th.termUsecase.Update(id, req.Name, req.Content, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTerm{
			ID:      updatedTerm.ID,
			Name:    updatedTerm.Name,
			Content: updatedTerm.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete termを削除するときのハンドラー
func (th *termHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.termUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
