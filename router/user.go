package router

import (
	"Go-User-System/controller"
	"github.com/labstack/echo"
)

func initUserGroup(group *echo.Group) {

	group.POST("", controller.UserRegister)
	//group.GET("/token", controller.UserGetToken)

	group.GET("/verify", controller.UserVerify)

	//group.GET("", controller.UserGetAllInfo)
	//group.GET("/:id", controller.UserGetInfo)
	//group.PUT("/:id", controller.UserUpdateInfo)
	//group.DELETE("/:id", controller.UserDelete)

	//group.GET("/tokenWX", controller.UserGetTokenWX)
	//group.GET("/bindWX", controller.UserBindWX)

	//TODO:JWT
}
