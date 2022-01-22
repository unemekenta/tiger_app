package model

import (
	"errors"
	"time"
)

// https://echo.labstack.com/guide/binding/
type Category struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	AncestorID int    `json:"ancestor_id"`
	Name       string `json:"name"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CategoriesWebsite struct {
	ID         int `gorm:"primaryKey" json:"id"`
	CategoryID int `json:"category_id"`
	WebsiteID  int `json:"website_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewCategory(ancestorId int, name string, updatedAt time.Time) (*Category, error) {
	if name == "" {
		return nil, errors.New("nameを入力してください")
	}

	category := &Category{
		AncestorID: ancestorId,
		Name:       name,
		UpdatedAt:  updatedAt,
	}
	return category, nil
}

func NewCategoriesWebsite(categoryId int, websiteId int, updatedAt time.Time) (*CategoriesWebsite, error) {
	if categoryId == 0 {
		return nil, errors.New("CategoryIDを入力してください")
	}

	cw := &CategoriesWebsite{
		CategoryID: categoryId,
		WebsiteID:  websiteId,
		UpdatedAt:  updatedAt,
	}

	return cw, nil
}

// Set categoryのセッター
func (c *Category) Set(ancestorId int, name string, updatedAt time.Time) error {
	if name == "" {
		return errors.New("nameを入力してください")
	}

	c.AncestorID = ancestorId
	c.Name = name
	c.UpdatedAt = updatedAt

	return nil
}
