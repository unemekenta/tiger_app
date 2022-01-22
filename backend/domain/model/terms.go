package model

import (
	"errors"
	"time"
)

// https://echo.labstack.com/guide/binding/
type Term struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTerm(name string, content string, updatedAt time.Time) (*Term, error) {
	if name == "" {
		return nil, errors.New("nameを入力してください")
	}

	term := &Term{
		Name:      name,
		Content:   content,
		UpdatedAt: updatedAt,
	}

	return term, nil
}

// Set termのセッター
func (c *Term) Set(name string, content string, updatedAt time.Time) error {
	if name == "" {
		return errors.New("nameを入力してください")
	}

	c.Name = name
	c.Content = content
	c.UpdatedAt = updatedAt

	return nil
}
