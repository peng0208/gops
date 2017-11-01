package handler

import (
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"gops/gops-common"
	"gops/gops-server/model"
)

type LoginInfo struct {
	Username string
	Password string
}

func CreateToken(c echo.Context) error {
	user := new(LoginInfo)
	if err := c.Bind(user); err != nil {
		return err
	}

	passwordInfo, eg := model.GetUserPassword(user.Username)

	if eg != nil {
		return eg
	} else if passwordInfo == nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "用户名不存在", nil})
	}

	result := model.CheckPassword(
		user.Username,
		user.Password,
		passwordInfo.Ctime,
		passwordInfo.Password,
	)

	if !result {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "用户名密码不匹配", nil})
	}

	secret := []byte("NDQ4MjUzNEMtRTI4==")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 10000).Unix()

	t, es := token.SignedString(secret)
	if es != nil {
		return es
	}

	ct := new(http.Cookie)
	ct.Name = "_token"
	ct.Value = t
	ct.Path = "/"

	cu := new(http.Cookie)
	cu.Name = "_cu"
	cu.Value = user.Username
	cu.Path = "/"

	c.SetCookie(ct)
	c.SetCookie(cu)

	return c.JSON(http.StatusOK, &common.ResponseType{1, "登录成功", nil})
}

func ClearToken(c echo.Context) error {
	ct, err := c.Cookie("_token")
	if err != nil {
		goto RETURN
	}
	ct.MaxAge = -1
	ct.Path = "/"
	c.SetCookie(ct)

RETURN:
	return c.JSON(http.StatusOK, &common.ResponseType{1, "已退出", nil})
}
