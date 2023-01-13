package view

import (
	"deptechdigital/controller"
	"deptechdigital/entity"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	srv controller.ProductControll
}

func NewProduct(e *echo.Echo, srv controller.ProductControll) {
	handler := productHandler{srv: srv}
	e.POST("/product", handler.InsertProduct())   // REGISTER USER
	e.PUT("/product", handler.UpdateProduct())    // EDIT DATA USER
	e.DELETE("/product", handler.DeleteProduct()) // DELETE USER
	e.GET("/product", handler.GetAllProduct())
}

func (uh *productHandler) InsertProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input entity.Product

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

		res, err := uh.srv.AddProduct(input)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, FailResponse("duplicate email on database"))
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot encrypt password"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("success register product", res))
	}
}

// UPDATE DATA USER
func (uh *productHandler) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input entity.Product

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind update data"))
		}

		res, err := uh.srv.EditProduct(input)
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusAccepted, SuccessResponse("success update product", res))
	}
}

// DELETE USER
func (uh *productHandler) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("wrong input id"))
		}
		err = uh.srv.RemoveProduct(uint(userID))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success delete data", nil))
	}
}

func (uh *productHandler) GetAllProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := uh.srv.ShowAllProduct()
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get data", res))
	}
}
