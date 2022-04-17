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

type Subscription struct {
	ID             int `gorm:"primaryKey" json:"id"`
	MoneyAccountID int `json:"money_account_id"`
	StartYear      int `json:"start_year"`
	StartMonth     int `json:"start_month"`
	EndYear        int `json:"end_year"`
	EndMonth       int `json:"end_month"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SubscriptionWithMoneyAccount struct {
	MoneyAccount MoneyAccount `gorm:"embedded"`
	Subscription Subscription `gorm:"embedded"`
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

func NewSubscription(startYear int, startMonth int, endYear int, endMonth int, userID int, moneyAccountLabelID int, amount int, title string, contents string, year int, month int, updatedAt time.Time) (*Subscription, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	subscription := &Subscription{
		MoneyAccountID: moneyAccountLabelID,
		StartYear:      startYear,
		StartMonth:     startMonth,
		EndYear:        endYear,
		EndMonth:       endMonth,
		UpdatedAt:      updatedAt,
	}

	return subscription, nil
}

func NewSubscriptionWithMoneyAccount(startYear int, startMonth int, endYear int, endMonth int, userID int, moneyAccountLabelID int, amount int, title string, contents string, year int, month int, updatedAt time.Time) (*SubscriptionWithMoneyAccount, error) {
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

	subscription := &Subscription{
		MoneyAccountID: moneyAccountLabelID,
		StartYear:      startYear,
		StartMonth:     startMonth,
		EndYear:        endYear,
		EndMonth:       endMonth,
		UpdatedAt:      updatedAt,
	}

	subscriptionWithMoneyAccount := &SubscriptionWithMoneyAccount{
		MoneyAccount: *moneyAccount,
		Subscription: *subscription,
	}

	return subscriptionWithMoneyAccount, nil
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
