package handler

import (
	"net/http"
	"strconv"

	"hacktiv/model"
	"hacktiv/usecase"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) userHandler {
	return userHandler{
		userUsecase: userUsecase,
	}
}

func (u userHandler) RegisterUserRoutes(e *echo.Echo) {
	e.GET("/users", u.GetAllUserHandler)
	e.POST("/users", u.CreateUserHandler)
	e.PUT("/users/:id", u.UpdateUserHandler)
	e.DELETE("/users/:id", u.DeleteUserHandler)

	e.POST("/users/transaction-example", u.TransactionExampleHandler)
}

// ShowAccount godoc
// @Summary      Create Account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.User
// @Failure      400  {object}  utils.HTTPError
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router       /accounts/{id} [get]
func (u userHandler) CreateUserHandler(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// fmt.Println("ACTIVE USER: ", user.ActiveUser)

	err = u.userUsecase.CreateUser(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (u userHandler) GetAllUserHandler(c echo.Context) error {
	users, err := u.userUsecase.GetAllUsers(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func (u userHandler) UpdateUserHandler(c echo.Context) error {
	return nil
}

func (u userHandler) DeleteUserHandler(c echo.Context) error {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = u.userUsecase.DeleteUser(c.Request().Context(), int64(idNum))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return nil
}

func (u userHandler) TransactionExampleHandler(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = u.userUsecase.TransactionExample(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return nil
}
