package handlers

import (
	"fmt"
	"net/http"
	"tiger_app/backend_src/models"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// todo mypage
func GetMypage(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	fmt.Println("hi")
	return c.JSON(http.StatusAccepted, "mypage")
}

func UserSignup(c echo.Context) error {
	user := new(models.Users)
	Db := models.OpenDBConn()
	db, err := Db.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	email := c.FormValue("email")
	name := c.FormValue("name")
	bytePassword := []byte(c.FormValue("password"))
	role_id := 1

	byteHashedPassword, _ := bcrypt.GenerateFromPassword(bytePassword, 10)
	hashedPassword := string(byteHashedPassword)

	user.Email = email
	user.Name = name
	user.Password = hashedPassword
	user.RoleId = role_id

	res := Db.Create(&user)
	if res.Error != nil {
		return res.Error
	} else {
		return c.JSON(http.StatusOK, "Welcome! "+user.Name)
	}
}

func UserLogin(c echo.Context) error {
	user := new(models.Users)
	Db := models.OpenDBConn()
	db, err := Db.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	formEmail := c.FormValue("email")
	Db.First(user, "email = ?", formEmail)
	dbPassword := user.Password
	byteDbPassword := []byte(dbPassword)
	formPassword := c.FormValue("password")
	byteFormPassword := []byte(formPassword)

	if err := bcrypt.CompareHashAndPassword(byteDbPassword, byteFormPassword); err != nil {
		return err
	} else {
		cookie := new(http.Cookie)
		cookie.Name = "username"
		cookie.Value = string(byteDbPassword)
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.Path = "/"
		fmt.Println(cookie)
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, "Hi! "+user.Name)
	}
}
