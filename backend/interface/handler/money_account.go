package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/now"
	"github.com/labstack/echo"
	"github.com/unemekenta/tiger_app/backend/usecase"
)

// MoneyAccountHandler moneyAccount handlerのinterface
type MoneyAccountHandler interface {
	Post() echo.HandlerFunc
	PostSubscription() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetByUser() echo.HandlerFunc
	GetSubscriptionWithMoneyAccountByUser() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type moneyAccountHandler struct {
	moneyAccountUsecase usecase.MoneyAccountUsecase
}

// NewMoneyAccountHandler moneyAccount handlerのコンストラクタ
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

type requestSubscription struct {
	MoneyAccount   requestMoneyAccount `json:"money_account"`
	MoneyAccountID int                 `json:"money_account_id"`
	StartYear      int                 `json:"start_year"`
	StartMonth     int                 `json:"start_month"`
	EndYear        int                 `json:"end_year"`
	EndMonth       int                 `json:"end_month"`
	UpdatedAt      time.Time           `json:"updatedAt"`
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

type responseSubscription struct {
	ID             int                  `json:"id"`
	MoneyAccount   responseMoneyAccount `gorm:"embedded" json:"moneyAccount"`
	MoneyAccountID int                  `json:"moneyAccountId"`
	StartYear      int                  `json:"startYear"`
	StartMonth     int                  `json:"startMonth"`
	EndYear        int                  `json:"endYear"`
	EndMonth       int                  `json:"endMonth"`
	UpdatedAt      time.Time            `json:"updatedAt"`
}

type responseUserInfo struct {
	MoneyAccountListIncome      []responseMoneyAccount `json:"moneyAccountListIncome"`
	MoneyAccountListExpenses    []responseMoneyAccount `json:"moneyAccountListExpenses"`
	MoneyAccountListIncomeSum   int                    `json:"moneyAccountListIncomeSum"`
	MoneyAccountListExpensesSum int                    `json:"moneyAccountListExpensesSum"`
	Today                       time.Time              `json:"today"`
	EndOfThisMonth              time.Time              `json:"endOfThisMonth"`
	RemainingDateOfThisMonth    int                    `json:"remainingDateOfThisMonth"`
	RemainingMoney              int                    `json:"remainingMoney"`
	RemainingMoneyPerDay        int                    `json:"remainingMoneyPerDay"`
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

// PostSubscription Subscriptionを作成するときのハンドラー
func (mh *moneyAccountHandler) PostSubscription() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestSubscription
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdSubscriptionWithMoneyAccount, err := mh.moneyAccountUsecase.CreateSubscription(req.StartYear, req.StartMonth, req.EndYear, req.EndMonth, req.MoneyAccount.UserID, req.MoneyAccount.MoneyAccountLabelID, req.MoneyAccount.Amount, req.MoneyAccount.Title, req.MoneyAccount.Contents, req.MoneyAccount.Year, req.MoneyAccount.Month, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		rma := responseMoneyAccount{
			ID:                  createdSubscriptionWithMoneyAccount.MoneyAccount.ID,
			MoneyAccountLabelID: createdSubscriptionWithMoneyAccount.MoneyAccount.MoneyAccountLabelID,
			Amount:              createdSubscriptionWithMoneyAccount.MoneyAccount.Amount,
			Title:               createdSubscriptionWithMoneyAccount.MoneyAccount.Title,
			Contents:            createdSubscriptionWithMoneyAccount.MoneyAccount.Contents,
			Year:                createdSubscriptionWithMoneyAccount.MoneyAccount.Year,
			Month:               createdSubscriptionWithMoneyAccount.MoneyAccount.Month,
			UpdatedAt:           createdSubscriptionWithMoneyAccount.MoneyAccount.UpdatedAt,
		}

		res := responseSubscription{
			ID:             createdSubscriptionWithMoneyAccount.Subscription.ID,
			MoneyAccount:   rma,
			MoneyAccountID: createdSubscriptionWithMoneyAccount.Subscription.MoneyAccountID,
			StartYear:      createdSubscriptionWithMoneyAccount.Subscription.StartYear,
			StartMonth:     createdSubscriptionWithMoneyAccount.Subscription.StartMonth,
			EndYear:        createdSubscriptionWithMoneyAccount.Subscription.EndYear,
			EndMonth:       createdSubscriptionWithMoneyAccount.Subscription.EndMonth,
			UpdatedAt:      createdSubscriptionWithMoneyAccount.Subscription.UpdatedAt,
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

		var mai []responseMoneyAccount
		var mae []responseMoneyAccount
		var maiSum int
		var maeSum int

		// github.com/jinzhu/now
		today := time.Now()
		endOfThisMonth := now.EndOfMonth()
		remainingDateOfThisMonthDuration := endOfThisMonth.Sub(today)

		remainingDateOfThisMonth := int(remainingDateOfThisMonthDuration.Hours()) / 24

		fmt.Println(today, endOfThisMonth, remainingDateOfThisMonth)

		for _, fm := range *foundMoneyAccounts {
			switch fm.MoneyAccountLabelID {
			case 1:
				maiSum += fm.Amount
				mai = append(mai, responseMoneyAccount{
					ID:                  fm.ID,
					MoneyAccountLabelID: fm.MoneyAccountLabelID,
					Amount:              fm.Amount,
					Title:               fm.Title,
					Contents:            fm.Contents,
					Year:                fm.Year,
					Month:               fm.Month,
					UpdatedAt:           fm.UpdatedAt,
				})
				break

			case 2:
				maeSum += fm.Amount
				mae = append(mae, responseMoneyAccount{
					ID:                  fm.ID,
					MoneyAccountLabelID: fm.MoneyAccountLabelID,
					Amount:              fm.Amount,
					Title:               fm.Title,
					Contents:            fm.Contents,
					Year:                fm.Year,
					Month:               fm.Month,
					UpdatedAt:           fm.UpdatedAt,
				})
				break

			default:
				break
			}
		}
		remainingMoney := maiSum - maeSum
		remainingMoneyPerDay := remainingMoney / remainingDateOfThisMonth

		res := responseUserInfo{
			MoneyAccountListIncome:      mai,
			MoneyAccountListExpenses:    mae,
			MoneyAccountListIncomeSum:   maiSum,
			MoneyAccountListExpensesSum: maeSum,
			Today:                       today,
			EndOfThisMonth:              endOfThisMonth,
			RemainingDateOfThisMonth:    remainingDateOfThisMonth,
			RemainingMoney:              remainingMoney,
			RemainingMoneyPerDay:        remainingMoneyPerDay,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// GetSubscriptionWithMoneyAccountByUser moneyAccountのうち、Subscriptionを取得するときのハンドラー
func (mh *moneyAccountHandler) GetSubscriptionWithMoneyAccountByUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundSubscriptionWithMoneyAccount, err := mh.moneyAccountUsecase.FindSubscriptionWithMoneyAccountByUser(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseSubscription

		for _, fs := range *foundSubscriptionWithMoneyAccount {
			rma := responseMoneyAccount{
				ID:                  fs.MoneyAccount.ID,
				MoneyAccountLabelID: fs.MoneyAccount.MoneyAccountLabelID,
				Amount:              fs.MoneyAccount.Amount,
				Title:               fs.MoneyAccount.Title,
				Contents:            fs.MoneyAccount.Contents,
				Year:                fs.MoneyAccount.Year,
				Month:               fs.MoneyAccount.Month,
				UpdatedAt:           fs.MoneyAccount.UpdatedAt,
			}

			res = append(res, responseSubscription{
				ID:             fs.Subscription.ID,
				MoneyAccount:   rma,
				MoneyAccountID: fs.Subscription.MoneyAccountID,
				StartYear:      fs.Subscription.StartYear,
				StartMonth:     fs.Subscription.StartMonth,
				EndYear:        fs.Subscription.EndYear,
				EndMonth:       fs.Subscription.EndMonth,
				UpdatedAt:      fs.Subscription.UpdatedAt,
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
