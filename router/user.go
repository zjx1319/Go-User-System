package router

import (
	"Go-User-System/config"
	"Go-User-System/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func initUserGroup(group *echo.Group) {

	group.POST("", controller.UserRegister)
	group.GET("/token", controller.UserGetToken)

	group.POST("/email", controller.UserVerify)

	group.GET("", controller.UserGetAllInfo, middleware.JWT([]byte(config.Config.JWT.Secret)))
	group.GET("/:id", controller.UserGetInfo, middleware.JWT([]byte(config.Config.JWT.Secret)))
	group.PUT("/:id", controller.UserUpdateInfo, middleware.JWT([]byte(config.Config.JWT.Secret)))
	group.DELETE("/:id", controller.UserDelete, middleware.JWT([]byte(config.Config.JWT.Secret)))

	group.GET("/WX/token", controller.UserGetTokenWX)
	group.POST("/WX", controller.UserBindWX, middleware.JWT([]byte(config.Config.JWT.Secret)))
	group.GET("/WX", controller.UserGetWXInfo, middleware.JWT([]byte(config.Config.JWT.Secret)))
}
