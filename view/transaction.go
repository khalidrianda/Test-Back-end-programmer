package view

import (
	"deptechdigital/controller"
	"deptechdigital/entity"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	srv controller.TransactionControll
}

func NewTransaction(e *echo.Echo, srv controller.TransactionControll) {
	handler := transactionHandler{srv: srv}
	e.POST("/transaction", handler.InsertTransaction())
	e.GET("/transaction", handler.GetAllTransaction())
	e.GET("/transaction/:id", handler.GetTransactionDetail())
}

func (uh *transactionHandler) InsertTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input entity.TransactionFormat

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

		res, err := uh.srv.AddTransaction(input)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, FailResponse("duplicate email on database"))
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot encrypt password"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("success post transaction", res))
	}
}

func (uh *transactionHandler) GetAllTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := uh.srv.ShowAllTransaction()
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get data", res))
	}
}

func (uh *transactionHandler) GetTransactionDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("wrong input id"))
		}

		res, err := uh.srv.ShowTransactionDetail(uint(transactionID))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get data", res))
	}
}
