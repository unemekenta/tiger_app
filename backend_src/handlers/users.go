package handlers

import (
	"net/http"
	"tiger_app/backend_src/models"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

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

	if err := Db.Create(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	} else {
		return c.JSON(http.StatusAccepted, "Welcome! "+user.Name)
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
		return c.JSON(http.StatusBadRequest, "BadRequest")
	} else {
		return c.JSON(http.StatusAccepted, "Hi! "+user.Name)
	}
}
