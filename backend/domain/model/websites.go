package model

import (
	"errors"
	"time"
)

// https://echo.labstack.com/guide/binding/
type Website struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	CompanyName string `json:"company_name"`
	Image       string `json:"image"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewWebsite(name string, url string, companyName string, image string, updatedAt time.Time) (*Website, error) {
	if name == "" {
		return nil, errors.New("nameを入力してください")
	}

	website := &Website{
		Name:        name,
		Url:         url,
		CompanyName: companyName,
		Image:       image,
		UpdatedAt:   updatedAt,
	}

	return website, nil
}

// Set websiteのセッター
func (c *Website) Set(name string, url string, companyName string, image string, updatedAt time.Time) error {
	if name == "" {
		return errors.New("nameを入力してください")
	}

	c.Name = name
	c.Url = url
	c.CompanyName = companyName
	c.Image = image
	c.UpdatedAt = updatedAt

	return nil
}
