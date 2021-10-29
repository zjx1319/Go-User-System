package main

import (
	"log"

	"Go-User-System/config"
	"Go-User-System/model"
	"Go-User-System/router"
	"Go-User-System/util"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.InitConfig()
	model.InitModel()
	util.InitUtil()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	apiGroup := e.Group("/api/v1")
	router.InitRouter(apiGroup)

	log.Fatal(e.Start(config.Config.App.Address))
}
