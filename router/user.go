package router

import (
	"Go-User-System/controller"
	"github.com/labstack/echo"
)

func initUserGroup(group *echo.Group) {

	group.POST("/", controller.UserRegister)

}
