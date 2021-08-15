package models

import "time"

// https://echo.labstack.com/guide/binding/
type Categories struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	AncestorID int    `json:"ancestor_id"`
	Name       string `json:"name"`
	CreatedAt  time.Time
	UpdateAt   time.Time
}
