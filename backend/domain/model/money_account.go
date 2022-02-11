package model

import (
	"errors"
	"time"
)

// https://echo.labstack.com/guide/binding/
type MoneyAccount struct {
	ID                  int    `gorm:"primaryKey" json:"id"`
	UserID              int    `json:"user_id"`
	MoneyAccountLabelID int    `json:"money_account_label_id"`
	Amount              int    `json:"amount"`
	Title               string `json:"title"`
	Contents            string `json:"contents"`
	Year                int    `json:"year"`
	Month               int    `json:"month"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func NewMoneyAccount(userID int, moneyAccountLabelID int, amount int, title string, contents string, year int, month int, updatedAt time.Time) (*MoneyAccount, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	moneyAccount := &MoneyAccount{
		UserID:              userID,
		MoneyAccountLabelID: moneyAccountLabelID,
		Amount:              amount,
		Title:               title,
		Year:                year,
		Month:               month,
		Contents:            contents,
		UpdatedAt:           updatedAt,
	}

	return moneyAccount, nil
}

// Set moneyAccountのセッター
func (c *MoneyAccount) Set(moneyAccountLabelID int, amount int, title string, contents string, year int, month int, updatedAt time.Time) error {
	if title == "" {
		return errors.New("titleを入力してください")
	}

	c.MoneyAccountLabelID = moneyAccountLabelID
	c.Amount = amount
	c.Title = title
	c.Contents = contents
	c.Year = year
	c.Month = month
	c.UpdatedAt = updatedAt

	return nil
}
