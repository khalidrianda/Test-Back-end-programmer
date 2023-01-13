package view

import (
	"deptechdigital/controller"
	"deptechdigital/entity"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	validate = validator.New()
)

type userHandler struct {
	srv controller.AdminControll
}

func New(e *echo.Echo, srv controller.AdminControll) {
	handler := userHandler{srv: srv}
	e.POST("/user", handler.Insert())   // REGISTER USER
	e.PUT("/user", handler.Update())    // EDIT DATA USER
	e.DELETE("/user", handler.Delete()) // DELETE USER
	e.POST("/login", handler.Login())   // LOGIN
	e.GET("/user", handler.GetAll())
}

func (uh *userHandler) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input entity.Admin

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		er := validate.Struct(input)
		if er != nil {
			if strings.Contains(er.Error(), "min") {
				return c.JSON(http.StatusBadRequest, FailResponse("min. 4 character"))
			} else if strings.Contains(er.Error(), "max") {
				return c.JSON(http.StatusBadRequest, FailResponse("max. 30 character"))
			} else if strings.Contains(er.Error(), "email") {
				return c.JSON(http.StatusBadRequest, FailResponse("must input valid email"))
			}
			return c.JSON(http.StatusBadRequest, FailResponse(er.Error()))
		}

		res, err := uh.srv.Add(input)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, FailResponse("duplicate email on database"))
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot encrypt password"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("success register user", res))
	}
}

// UPDATE DATA USER
func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input entity.Admin

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind update data"))
		}

		res, err := uh.srv.Edit(input)
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusAccepted, SuccessResponse("success update user", res))
	}
}

// DELETE USER
func (uh *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("wrong input id"))
		}
		err = uh.srv.Remove(uint(userID))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success delete data", nil))
	}
}

// LOGIN
func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input entity.Admin

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		res, err := uh.srv.LogIn(input)
		if err != nil {
			if strings.Contains(err.Error(), "an invalid client request") {
				return c.JSON(http.StatusBadRequest, FailResponse("email doesn't exist."))
			} else if strings.Contains(err.Error(), "password not match") {
				return c.JSON(http.StatusBadRequest, FailResponse("password not match."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusAccepted, SuccessResponse("success login", res))
	}
}

func (uh *userHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := uh.srv.ShowAll()
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get data", res))
	}
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
