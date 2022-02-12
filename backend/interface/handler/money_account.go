package handler

import (
	"net/http"
	"strconv"
	"time"

	"backend/usecase"

	"github.com/labstack/echo"
)

// MoneyAccountHandler moneyAccount handlerのinterface
type MoneyAccountHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetByUser() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type moneyAccountHandler struct {
	moneyAccountUsecase usecase.MoneyAccountUsecase
}

// MoneyAccountHandler moneyAccount handlerのコンストラクタ
func NewMoneyAccountHandler(moneyAccountUsecase usecase.MoneyAccountUsecase) MoneyAccountHandler {
	return &moneyAccountHandler{moneyAccountUsecase: moneyAccountUsecase}
}

type requestMoneyAccount struct {
	UserID              int       `json:"userId"`
	MoneyAccountLabelID int       `json:"moneyAccountLabelId"`
	Amount              int       `json:"amount"`
	Title               string    `json:"title"`
	Contents            string    `json:"contents"`
	Year                int       `json:"year"`
	Month               int       `json:"month"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

type responseMoneyAccount struct {
	ID                  int       `json:"id"`
	MoneyAccountLabelID int       `json:"moneyAccountLabelId"`
	Amount              int       `json:"amount"`
	Title               string    `json:"title"`
	Contents            string    `json:"contents"`
	Year                int       `json:"year"`
	Month               int       `json:"month"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

// Post moneyAccountを保存するときのハンドラー
func (mh *moneyAccountHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestMoneyAccount
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdMoneyAccount, err := mh.moneyAccountUsecase.Create(req.UserID, req.MoneyAccountLabelID, req.Amount, req.Title, req.Contents, req.Year, req.Month, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseMoneyAccount{
			ID:                  createdMoneyAccount.ID,
			MoneyAccountLabelID: createdMoneyAccount.MoneyAccountLabelID,
			Amount:              createdMoneyAccount.Amount,
			Title:               createdMoneyAccount.Title,
			Contents:            createdMoneyAccount.Contents,
			Year:                createdMoneyAccount.Year,
			Month:               createdMoneyAccount.Month,
			UpdatedAt:           createdMoneyAccount.UpdatedAt,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get moneyAccountを取得するときのハンドラー
func (mh *moneyAccountHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundMoneyAccount, err := mh.moneyAccountUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseMoneyAccount{
			ID:                  foundMoneyAccount.ID,
			MoneyAccountLabelID: foundMoneyAccount.MoneyAccountLabelID,
			Amount:              foundMoneyAccount.Amount,
			Title:               foundMoneyAccount.Title,
			Contents:            foundMoneyAccount.Contents,
			Year:                foundMoneyAccount.Year,
			Month:               foundMoneyAccount.Month,
			UpdatedAt:           foundMoneyAccount.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetByUser moneyAccountをUserて取得するときのハンドラー
func (mh *moneyAccountHandler) GetByUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		year, err := strconv.Atoi((c.Param("year")))
		month, err := strconv.Atoi((c.Param("month")))
		foundMoneyAccounts, err := mh.moneyAccountUsecase.FindByUser(id, year, month)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseMoneyAccount

		for _, fm := range *foundMoneyAccounts {
			res = append(res, responseMoneyAccount{
				ID:                  fm.ID,
				MoneyAccountLabelID: fm.MoneyAccountLabelID,
				Amount:              fm.Amount,
				Title:               fm.Title,
				Contents:            fm.Contents,
				Year:                fm.Year,
				Month:               fm.Month,
				UpdatedAt:           fm.UpdatedAt,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put moneyAccountを更新するときのハンドラー
func (mh *moneyAccountHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestMoneyAccount
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedMoneyAccount, err := mh.moneyAccountUsecase.Update(id, req.MoneyAccountLabelID, req.Amount, req.Title, req.Contents, req.Year, req.Month, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseMoneyAccount{
			ID:                  updatedMoneyAccount.ID,
			MoneyAccountLabelID: updatedMoneyAccount.MoneyAccountLabelID,
			Amount:              updatedMoneyAccount.Amount,
			Title:               updatedMoneyAccount.Title,
			Contents:            updatedMoneyAccount.Contents,
			Year:                updatedMoneyAccount.Year,
			Month:               updatedMoneyAccount.Month,
			UpdatedAt:           updatedMoneyAccount.UpdatedAt,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete moneyAccountを削除するときのハンドラー
func (mh *moneyAccountHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = mh.moneyAccountUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
