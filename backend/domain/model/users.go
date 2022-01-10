package model

import (
	"errors"
	"time"
)

// https://echo.labstack.com/guide/binding/
type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"content"`
	RoleId    int    `json:"role_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(email string, name string, password string, roleid int, updatedAt time.Time) (*User, error) {
	if name == "" {
		return nil, errors.New("nameを入力してください")
	}

	user := &User{
		Email:     email,
		Name:      name,
		Password:  password,
		RoleId:    roleid,
		UpdatedAt: updatedAt,
	}

	return user, nil
}

// Set userのセッター
func (c *User) Set(email string, name string, password string, roleid int, updatedAt time.Time) error {
	if name == "" {
		return errors.New("nameを入力してください")
	}
	c.Email = email
	c.Name = name
	c.Password = password
	c.RoleId = roleid
	c.UpdatedAt = updatedAt

	return nil
}
