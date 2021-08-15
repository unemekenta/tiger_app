package models

import (
	"time"
)

// https://echo.labstack.com/guide/binding/
type Users struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Email     string `json:"e-mail"`
	Name      string `json:"name"`
	Password  string `json:"content"`
	RoleId    int    `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
