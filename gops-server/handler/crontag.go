package handler

import (
	"gops/gops-server/model"
	"github.com/labstack/echo"
	"net/http"
	"gops/gops-common"
)

func GetCronTagList(c echo.Context) error {
	page, size := common.GetPageParams(c)

	tags, err := model.GetCronTagList(page, size)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &common.ResponseType{0, "", tags})
}

func AddCronTag(c echo.Context) error {
	tag := new(model.CronTag)
	if err := c.Bind(tag); err != nil {
		return err
	}

	result, err := tag.Create()
	if err != nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "项目已存在", nil})
	}
	if result {
		return c.JSON(http.StatusOK, &common.ResponseType{1, "新建项目成功", nil})
	}
	return c.JSON(http.StatusOK, &common.ResponseType{2, "新建项目失败", nil})
}

func RemoveCronTag(c echo.Context) error {
	tag := new(model.CronTag)
	tag.TagId = c.Param("tagid")

	result, err := tag.Delete()
	if err != nil {
		return err
	}
	if result {
		return c.JSON(http.StatusOK, &common.ResponseType{1, "删除项目成功", nil})
	}
	return c.JSON(http.StatusOK, &common.ResponseType{2, "删除项目失败", nil})
}