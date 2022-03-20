package model

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

type SessionUser struct {
	Val string `json:"val"`
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

func CreateJwt(foundUser *User, byteDbPassword []byte, byteFormPassword []byte) (string, string, error) {
	if err := bcrypt.CompareHashAndPassword(byteDbPassword, byteFormPassword); err != nil {
		return "", "", err
	} else {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["admin"] = true
		claims["id"] = foundUser.ID
		claims["name"] = foundUser.Name
		claims["iat"] = time.Now()
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return "", "", err
		}

		u, err := uuid.NewRandom()
		if err != nil {
			return "", "", err
		}
		uu := u.String()

		return uu, t, nil
	}
}
