package models

import "time"

// https://echo.labstack.com/guide/binding/
type WebsiteContents struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	WebsiteID int       `json:"website_id"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
