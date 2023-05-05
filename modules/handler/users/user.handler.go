package users

import (
	"fmt"
	"net/http"
	"strconv"

	eu "github.com/berrylradianh/go_berryl-radian-hamesha/modules/entity/users"
	"github.com/berrylradianh/go_berryl-radian-hamesha/modules/services"
	uu "github.com/berrylradianh/go_berryl-radian-hamesha/modules/usecase/users"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase uu.Usecase
}

func (handler Handler) GetAllUsers() echo.HandlerFunc {
	return func(e echo.Context) error {
		var users []eu.User

		users, err := handler.Usecase.GetAllUsers()
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all users",
			"users":   users,
		})
	}
}

func (handler Handler) GetUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user eu.User
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		user, err = handler.Usecase.GetUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get user",
			"user":    user,
		})
	}
}

func (handler Handler) CreateUser() echo.HandlerFunc {
	var user eu.User

	return func(e echo.Context) error {
		if err := e.Bind(&user); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		if err := e.Validate(user); err != nil {
			message := ""
			for _, e := range err.(validator.ValidationErrors) {
				message += fmt.Sprintf("%s is required ", e.Field())
			}
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": message,
				"errors":  err.Error(),
			})
		}

		HashedPassword, err := services.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(HashedPassword)

		err = handler.Usecase.CreateUser(user)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new user",
			"user":    &user,
		})
	}
}

func (handler Handler) UpdateUser() echo.HandlerFunc {
	var user eu.User

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&user); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err = handler.Usecase.UpdateUser(id, user)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update user",
		})
	}
}

func (handler Handler) DeleteUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = handler.Usecase.DeleteUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete User`",
		})
	}
}
