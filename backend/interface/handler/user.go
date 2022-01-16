package handler

import (
	"net/http"
	"strconv"
	"time"

	"backend/usecase"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler user handlerのinterface
type UserHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Signup() echo.HandlerFunc
	Login() echo.HandlerFunc
	LoginCheck() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// UserHandler user handlerのコンストラクタ
func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type requestUserSession struct {
	Uuid string `json:"uuid"`
}

type requestUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type requestUser struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	RoleId    int       `json:"roleId"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type responseUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleId   int    `json:"roleId"`
}

// Post userを保存するときのハンドラー
func (uh *userHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdUser, err := uh.userUsecase.Create(req.Email, req.Name, req.Password, req.RoleId, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:       createdUser.ID,
			Email:    createdUser.Email,
			Name:     createdUser.Name,
			Password: createdUser.Password,
			RoleId:   createdUser.RoleId,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get userを取得するときのハンドラー
func (uh *userHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundUser, err := uh.userUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:       foundUser.ID,
			Email:    foundUser.Email,
			Name:     foundUser.Name,
			Password: foundUser.Password,
			RoleId:   foundUser.RoleId,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Put userを更新するときのハンドラー
func (uh *userHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedUser, err := uh.userUsecase.Update(id, req.Email, req.Name, req.Password, req.RoleId, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:       updatedUser.ID,
			Email:    updatedUser.Email,
			Name:     updatedUser.Name,
			Password: updatedUser.Password,
			RoleId:   updatedUser.RoleId,
		}

		return c.JSON(http.StatusOK, res)
	}
}

// Delete userを削除するときのハンドラー
func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = uh.userUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}

// Signup userを登録するときのハンドラー
func (uh *userHandler) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUser
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		bytePassword := []byte(req.Password)
		byteHashedPassword, _ := bcrypt.GenerateFromPassword(bytePassword, 10)
		hashedPassword := string(byteHashedPassword)

		createdUser, err := uh.userUsecase.Create(req.Email, req.Name, hashedPassword, req.RoleId, req.UpdatedAt)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:       createdUser.ID,
			Email:    createdUser.Email,
			Name:     createdUser.Name,
			Password: createdUser.Password,
			RoleId:   createdUser.RoleId,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Login userを登録するときのハンドラー
func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUserLogin
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundUser, err := uh.userUsecase.FindByPassword(req.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		dbPassword := foundUser.Password
		byteDbPassword := []byte(dbPassword)
		formPassword := req.Password
		byteFormPassword := []byte(formPassword)

		if err := bcrypt.CompareHashAndPassword(byteDbPassword, byteFormPassword); err != nil {
			return err
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
				return err
			}

			u, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			uu := u.String()

			uh.userUsecase.CreateSession(uu, t)

		}

		res := responseUser{
			ID:       foundUser.ID,
			Email:    foundUser.Email,
			Name:     foundUser.Name,
			Password: foundUser.Password,
			RoleId:   foundUser.RoleId,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// LoginCheck ログイン状態を確認するときのハンドラー
func (uh *userHandler) LoginCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestUserSession
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res, err := uh.userUsecase.LoginCheck(req.Uuid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, res)
	}
}
