package models

import "time"

// https://echo.labstack.com/guide/binding/
type Terms struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time `gorm:"index"`
}
