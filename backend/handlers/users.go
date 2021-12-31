package handlers

import (
	"backend/models"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c echo.Context) error {
	user := new(models.Users)
	Db := models.OpenDBConn()
	db, err := Db.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	user_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	Db.Where("id = ?", user_id).First(&user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// 権限チェック
func GetAuthCheck(c echo.Context) error {
	return c.JSON(http.StatusAccepted, "")
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
		claims["id"] = user.ID
		claims["name"] = formEmail
		claims["iat"] = time.Now()
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		u, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		uu := u.String()

		CreateSession(c, uu, t)

		return c.JSON(http.StatusOK, uu)

	}
}

func UserLoginCheck(c echo.Context) error {
	r := models.OpenRedisConn()
	key := c.QueryParam("key")
	val, err := r.Get(key).Result()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, val)
}

func CreateSession(c echo.Context, uuid string, value string) error {
	r := models.OpenRedisConn()
	// TODO セッション時間変更
	err := r.Set(uuid, value, time.Minute*10).Err()
	if err != nil {
		return err
	}
	return nil
}
