package handler

import (
	"gops/gops-server/model"
	"github.com/labstack/echo"
	"net/http"
	"gops/gops-common"
)

func GetUserList(c echo.Context) error {
	page, size := common.GetPageParams(c)

	users, err := model.GetUserList(page, size)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &common.ResponseType{0, "", users})
}

func AddUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	ctime := common.StampInt64()
	cpwd, en := model.EncryptPassword(
		user.Username,
		user.Password,
		common.StampToString(ctime),
	)
	if en != nil {
		return en
	}
	user.Password = string(cpwd)
	user.Ctime = ctime
	result, err := user.Create()
	if err != nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "用户名已存在", nil})
	}
	if result {
		return c.JSON(http.StatusOK, &common.ResponseType{1, "新建用户成功", nil})
	}
	return c.JSON(http.StatusOK, &common.ResponseType{2, "新建用户失败", nil})
}

func ChangeUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	user.UserId = c.Param("userid")
	user.Mtime = common.StampInt64()

	result, err := user.Update()
	if err != nil {
		return err
	}
	if result {
		return c.JSON(http.StatusOK, &common.ResponseType{1, "修改用户成功", nil})
	}
	return c.JSON(http.StatusOK, &common.ResponseType{2, "修改用户失败", nil})
}

func RemoveUser(c echo.Context) error {
	user := new(model.User)
	user.UserId = c.Param("userid")

	result, err := user.Delete()
	if err != nil {
		return err
	}
	if result {
		return c.JSON(http.StatusOK, &common.ResponseType{1, "删除用户成功", nil})
	}
	return c.JSON(http.StatusOK, &common.ResponseType{2, "删除用户失败", nil})
}
