package view

import (
	"deptechdigital/controller"
	"deptechdigital/entity"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type kategoriHandler struct {
	srv controller.KategoriControll
}

func NewKategori(e *echo.Echo, srv controller.KategoriControll) {
	handler := kategoriHandler{srv: srv}
	e.POST("/kategori", handler.InsertKategori())   // REGISTER USER
	e.PUT("/kategori", handler.UpdateKategori())    // EDIT DATA USER
	e.DELETE("/kategori", handler.DeleteKategori()) // DELETE USER
	e.GET("/kategori", handler.GetAllKategori())    // LOGIN
}

func (uh *kategoriHandler) InsertKategori() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input entity.KategoriProduct

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

		res, err := uh.srv.AddKategori(input)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, FailResponse("duplicate email on database"))
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot encrypt password"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("success insert kategori", res))
	}
}

// UPDATE DATA USER
func (uh *kategoriHandler) UpdateKategori() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input entity.KategoriProduct

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind update data"))
		}

		res, err := uh.srv.EditKategori(input)
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusAccepted, SuccessResponse("success update kategori", res))
	}
}

// DELETE USER
func (uh *kategoriHandler) DeleteKategori() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("wrong input id"))
		}
		err = uh.srv.RemoveKategori(uint(userID))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success delete data", nil))
	}
}

func (uh *kategoriHandler) GetAllKategori() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := uh.srv.ShowAllKategori()
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusNotFound, FailResponse("data not found."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponse("success get data", res))
	}
}
