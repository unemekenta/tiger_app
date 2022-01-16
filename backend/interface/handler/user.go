package handler

import (
	"net/http"
	"strconv"
	"time"

	"backend/usecase"

	"github.com/labstack/echo"
)

// UserHandler user handlerのinterface
type UserHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// UserHandler user handlerのコンストラクタ
func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type requestUser struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"content"`
	RoleId    int       `json:"role_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type responseUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"content"`
	RoleId   int    `json:"role_id"`
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
