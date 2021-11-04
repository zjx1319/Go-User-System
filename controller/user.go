package controller

import (
	"Go-User-System/config"
	"Go-User-System/model"
	"Go-User-System/util"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type paramUserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func UserRegister(c echo.Context) error {
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
	err = model.AddVerifyCode(verifyCode, user.ID, user.Email)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	verifyURL := fmt.Sprintf("%s?id=%d&code=%s", config.Config.SMTP.VerifyURL, user.ID, verifyCode)
	err = util.SendEmail(user.Email, "注册邮箱验证", fmt.Sprintf("你好！%s，请打开以下链接验证你的邮箱：<a href=%s>%s</a>", user.Username, verifyURL, verifyURL))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "发送验证邮件失败")
	}

	return c.String(http.StatusOK, "")
}

type paramUserVerify struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

func UserVerify(c echo.Context) error {
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

type paramUserGetToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type responseUserGetToken struct {
	ID     int    `json:"id"`
	Token  string `json:"token"`
	Expire int64  `json:"expire_time"`
}

func UserGetToken(c echo.Context) error {
	var param paramUserGetToken
	if err := c.Bind(&param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数错误")
	}
	if param.Username == "" || param.Password == "" {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数不足")
	}

	user, is, err := model.GetUserByName(param.Username)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if !is {
		return util.ErrorResponse(c, http.StatusBadRequest, "用户未注册！")
	}

	if user.Password != util.MD5(param.Password) {
		return util.ErrorResponse(c, http.StatusBadRequest, "用户名或密码错误！")
	}
	if !user.Verified {
		return util.ErrorResponse(c, http.StatusBadRequest, "邮箱未验证！")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	expireTime := time.Now().Add(time.Duration(config.Config.JWT.Expire) * time.Minute).Unix()
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = user.ID
	claims["exp"] = expireTime
	t, err := token.SignedString([]byte(config.Config.JWT.Secret))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, responseUserGetToken{
		ID:     user.ID,
		Token:  t,
		Expire: expireTime,
	})
}

type responseUserGetInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func UserGetInfo(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if ID == 0 || err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数错误")
	}
	userID := int(c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["ID"].(float64))

	if !(userID == ID) {
		isAdmin, err := model.IsUserAdmin(userID)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, "")
		}
		if !isAdmin {
			return util.ErrorResponse(c, http.StatusBadRequest, "权限不足")
		}
	}

	user, is, err := model.GetUserByID(ID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if !is {
		return util.ErrorResponse(c, http.StatusBadRequest, "用户不存在")
	}

	return c.JSON(http.StatusOK, responseUserGetInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	})
}

func UserGetAllInfo(c echo.Context) error {
	userID := int(c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["ID"].(float64))
	isAdmin, err := model.IsUserAdmin(userID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if !isAdmin {
		return util.ErrorResponse(c, http.StatusBadRequest, "权限不足")
	}

	users, err := model.GetAllUser()
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	var response []responseUserGetInfo
	for _, user := range users {
		response = append(response, responseUserGetInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		})
	}
	return c.JSON(http.StatusOK, response)
}

type paramUserUpdateInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func UserUpdateInfo(c echo.Context) error {
	var param paramUserUpdateInfo
	err := c.Bind(&param)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数错误")
	}
	ID, err := strconv.Atoi(c.Param("id"))
	if ID == 0 || err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数错误")
	}
	userID := int(c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["ID"].(float64))
	if !(userID == ID) || param.Role != "" {
		isAdmin, err := model.IsUserAdmin(userID)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, "")
		}
		if !isAdmin {
			return util.ErrorResponse(c, http.StatusBadRequest, "权限不足")
		}
	}

	is, err := model.IsUserExistByID(ID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if !is {
		return util.ErrorResponse(c, http.StatusBadRequest, "用户不存在")
	}

	if param.Username != "" {
		is, err = model.IsUserExistByName(param.Username)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, "")
		}
		if is {
			return util.ErrorResponse(c, http.StatusBadRequest, "用户名已存在")
		}
	}

	if param.Email != "" {
		is, err = model.IsUserExistByEmail(param.Email)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, "")
		}
		if is {
			return util.ErrorResponse(c, http.StatusBadRequest, "邮箱已使用")
		}
	}

	if param.Password != "" {
		param.Password = util.MD5(param.Password)
		err = model.BanToken(c.Get("user").(*jwt.Token).Raw)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, "")
		}
	}

	err = model.UpdateUser(ID, model.User{
		Username: param.Username,
		Password: param.Password,
		Email:    param.Email,
		Role:     param.Role,
	})

	if param.Email != "" {
		verifyCode := util.GetRandomString(32)
		err = model.AddVerifyCode(verifyCode, ID, param.Email)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, "")
		}

		verifyURL := fmt.Sprintf("%s?id=%d&code=%s", config.Config.SMTP.VerifyURL, ID, verifyCode)
		err = util.SendEmail(param.Email, "更改邮箱验证", fmt.Sprintf("你好！你的邮箱已更改，请打开以下链接验证你的邮箱：<a href=%s>%s</a>", verifyURL, verifyURL))
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, "发送验证邮件失败")
		}

		err = model.UpdateUserVerified(ID, false)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, "")
		}
	}

	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	user, is, err := model.GetUserByID(ID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if !is {
		return util.ErrorResponse(c, http.StatusBadRequest, "用户不存在")
	}

	return c.JSON(http.StatusOK, responseUserGetInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	})
}

func UserDelete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if ID == 0 || err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数错误")
	}
	userID := int(c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["ID"].(float64))
	isAdmin, err := model.IsUserAdmin(userID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	if !isAdmin {
		return util.ErrorResponse(c, http.StatusBadRequest, "权限不足")
	}

	is, err := model.IsUserExistByID(ID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if !is {
		return util.ErrorResponse(c, http.StatusBadRequest, "用户不存在")
	}

	err = model.DeleteUser(ID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	return c.String(http.StatusOK, "")
}

type WXResponse struct {
	AccessToken string `json:"access_token"`
	Openid      string `json:"openid"`
}

type WXResponseInfo struct {
	Nickname string `json:"nickname"`
}

type paramUserBindWX struct {
	Code string `json:"code"`
}

type responseUserBindWX struct {
	WXName string `json:"wx_name"`
}

func UserBindWX(c echo.Context) error {
	var param paramUserBindWX
	err := c.Bind(&param)
	if err != nil || param.Code == "" {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数错误")
	}
	userID := int(c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["ID"].(float64))

	response, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", config.Config.WX.AppID, config.Config.WX.AppSecret, param.Code))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	var wxResponse WXResponse
	err = json.Unmarshal(body, &wxResponse)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if wxResponse.Openid == "" || wxResponse.AccessToken == "" {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	response, err = http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", wxResponse.AccessToken, wxResponse.Openid))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	var wxResponseInfo WXResponseInfo
	err = json.Unmarshal(body, &wxResponseInfo)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if wxResponseInfo.Nickname == "" {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	err = model.BindWX(userID, wxResponseInfo.Nickname, wxResponse.Openid)
	if err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, "超时或已绑定其他账号")
	}

	return c.JSON(http.StatusOK, responseUserBindWX{
		WXName: wxResponseInfo.Nickname,
	})
}

type paramUserGetTokenWX struct {
	Code string `json:"code"`
}

func UserGetTokenWX(c echo.Context) error {
	var param paramUserGetTokenWX
	err := c.Bind(&param)
	if err != nil || param.Code == "" {
		return util.ErrorResponse(c, http.StatusBadRequest, "参数错误")
	}

	response, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", config.Config.WX.AppID, config.Config.WX.AppSecret, param.Code))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	var wxResponse WXResponse
	err = json.Unmarshal(body, &wxResponse)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if wxResponse.Openid == "" {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	user, is, err := model.GetUserByWX(wxResponse.Openid)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if !is {
		return util.ErrorResponse(c, http.StatusBadRequest, "未绑定账号！")
	}

	if !user.Verified {
		return util.ErrorResponse(c, http.StatusBadRequest, "邮箱未验证！")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	expireTime := time.Now().Add(time.Duration(config.Config.JWT.Expire) * time.Minute).Unix()
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = user.ID
	claims["exp"] = expireTime
	t, err := token.SignedString([]byte(config.Config.JWT.Secret))
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, responseUserGetToken{
		ID:     user.ID,
		Token:  t,
		Expire: expireTime,
	})
}

type responseUserGetWXInfo struct {
	WXName string `json:"wx_name"`
}

func UserGetWXInfo(c echo.Context) error {
	userID := int(c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["ID"].(float64))

	wxName, is, err := model.GetWXName(userID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, "")
	}
	if !is {
		return util.ErrorResponse(c, http.StatusBadRequest, "未绑定账号！")
	}

	return c.JSON(http.StatusOK, responseUserGetWXInfo{
		WXName: wxName,
	})
}
