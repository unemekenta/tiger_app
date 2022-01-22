package model

import (
	"errors"
	"time"
)

// https://echo.labstack.com/guide/binding/
type WebsiteContent struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	WebsiteID int       `json:"website_id"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewWebsiteContent(websiteId int, title string, contents string, updatedAt time.Time) (*WebsiteContent, error) {
	if title == "" {
		return nil, errors.New("titleを入力してください")
	}

	wc := &WebsiteContent{
		WebsiteID: websiteId,
		Title:     title,
		Contents:  contents,
		UpdatedAt: updatedAt,
	}

	return wc, nil
}

// Set websiteContentのセッター
func (c *WebsiteContent) Set(title string, contents string, updatedAt time.Time) error {
	if title == "" {
		return errors.New("titleを入力してください")
	}

	c.Title = title
	c.Contents = contents
	c.UpdatedAt = updatedAt

	return nil
}
