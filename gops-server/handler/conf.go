package handler

import (
	"gops/gops-server/model"
	"github.com/labstack/echo"
	"net/http"
	"gops/gops-common"
)

func GetConfFileList(c echo.Context) error {
	page, size := common.GetPageParams(c)
	env := c.Param("env")
	files, err := model.GetConfFileList(env, page, size)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &common.ResponseType{0, "", files})
}

func AddFile(c echo.Context) error {
	env := c.Param("env")
	file := new(model.ConfFile)
	if err := c.Bind(file); err != nil {
		return err
	}

	_, err := file.Create(env)
	if err != nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "新建配置失败", nil})
	}
	return c.JSON(http.StatusOK, &common.ResponseType{1, "新建配置成功", nil})

}

func ChangeFile(c echo.Context) error {
	env := c.Param("env")
	confid := c.Param("confid")
	file := new(model.ConfFile)
	if err := c.Bind(file); err != nil {
		return err
	}
	file.ConfId = confid
	_, err := file.Update(env)
	if err != nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "修改配置失败", nil})
	}
	return c.JSON(http.StatusOK, &common.ResponseType{1, "修改配置成功", nil})

}

func RemoveFile(c echo.Context) error {
	env := c.Param("env")
	confid := c.Param("confid")
	file := new(model.ConfFile)
	file.ConfId = confid

	_, err := file.Delete(env)
	if err != nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "删除项目失败", nil})
	}
	return c.JSON(http.StatusOK, &common.ResponseType{1, "删除项目成功", nil})

}
