package router

import (
	"Go-User-System/config"
	"Go-User-System/controller"
	"Go-User-System/model"
	"Go-User-System/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func initUserGroup(group *echo.Group) {

	group.POST("", controller.UserRegister)
	group.GET("/token", controller.UserGetToken)

	group.POST("/email", controller.UserVerify)

	group.GET("", controller.UserGetAllInfo, middleware.JWT([]byte(config.Config.JWT.Secret)), JWTCheck)
	group.GET("/:id", controller.UserGetInfo, middleware.JWT([]byte(config.Config.JWT.Secret)), JWTCheck)
	group.PUT("/:id", controller.UserUpdateInfo, middleware.JWT([]byte(config.Config.JWT.Secret)), JWTCheck)
	group.DELETE("/:id", controller.UserDelete, middleware.JWT([]byte(config.Config.JWT.Secret)), JWTCheck)

	group.GET("/WX/token", controller.UserGetTokenWX)
	group.GET("/WX/bind", controller.UserBindWX, middleware.JWT([]byte(config.Config.JWT.Secret)), JWTCheck)
	group.POST("/WX", controller.UserGetWXInfo, middleware.JWT([]byte(config.Config.JWT.Secret)), JWTCheck)
}

func JWTCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		is, err := model.IsTokenBanned(c.Get("user").(*jwt.Token).Raw)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, "")
		}
		if is {
			return util.ErrorResponse(c, http.StatusBadRequest, "请重新登录")
		}
		return next(c)
	}
}
