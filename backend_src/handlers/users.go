package handlers

import (
	"net/http"
	"tiger_app/backend_src/models"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// todo mypage
func GetMypage(c echo.Context) error {
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

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["admin"] = true
		claims["name"] = formEmail
		claims["iat"] = time.Now()
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})

	}
}
