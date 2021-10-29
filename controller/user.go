package controller

import (
	"Go-User-System/config"
	"Go-User-System/model"
	"Go-User-System/util"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type paramUserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func UserRegister(c echo.Context) (err error) {
	var param paramUserRegister
	if err := c.Bind(&param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数错误")
	}
	if param.Username == "" || param.Password == "" || param.Email == "" {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数不足")
	}

	is, err := model.IsUserExistByName(param.Username)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if is {
		return util.ErrorResponse(c, http.StatusBadRequest, "用户名已存在")
	}

	is, err = model.IsUserExistByEmail(param.Email)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if is {
		return util.ErrorResponse(c, http.StatusBadRequest, "Email已存在")
	}

	user := model.User{
		Username: param.Username,
		Password: util.MD5(param.Password),
		Email:    param.Email,
		Verified: false,
		Role:     "default",
	}

	user.ID, err = model.AddUser(user)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	verifyCode := util.GetRandomString(32)
	err = model.AddVerifyCode(verifyCode, user)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	verifyURL := config.Config.App.Address + "/api/v1/user/verify?id=" + strconv.Itoa(user.ID) + "&code=" + verifyCode
	err = util.SendEmail(user.Email, "注册邮箱验证", "你好！"+user.Username+"，请打开以下链接验证你的邮箱：<a href="+verifyURL+">"+verifyURL+"</a>")
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "发送验证邮件失败")
	}

	return c.String(http.StatusOK, "")
}

type paramUserVerify struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

func UserVerify(c echo.Context) (err error) {
	var param paramUserVerify
	if err := c.Bind(&param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数错误")
	}
	if param.ID == 0 || param.Code == "" {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数不足")
	}

	code, is, err := model.GetVerifyCode(param.ID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if !is {
		return util.ErrorResponse(c, http.StatusBadRequest, "验证失败！")
	}
	if code != param.Code {
		return util.ErrorResponse(c, http.StatusBadRequest, "验证码错误！")
	}

	err = model.DeleteVerifyCode(param.Code)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	err = model.UpdateUserVerified(param.ID, true)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	return c.String(http.StatusOK, "")
}
