package models

import "time"

// https://echo.labstack.com/guide/binding/
type Websites struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	CompanyName string `json:"company_name"`
	CreatedAt   time.Time
	UpdateAt    time.Time
}
