package main

import (
	"deptechdigital/config"
	"deptechdigital/controller"
	"deptechdigital/model"
	"deptechdigital/utils/database"
	"deptechdigital/view"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	adminModel := model.AdminModel{db}
	adminControll := controller.AdminControll{adminModel}
	kategoriModel := model.KategoriModel{db}
	kategoriControll := controller.KategoriControll{kategoriModel}
	productModel := model.ProductModel{db}
	productControll := controller.ProductControll{productModel}
	transactionModel := model.TransactionModel{db}
	transactionControll := controller.TransactionControll{transactionModel}
	view.New(e, adminControll)
	view.NewKategori(e, kategoriControll)
	view.NewProduct(e, productControll)
	view.NewTransaction(e, transactionControll)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServerPort)))
}
