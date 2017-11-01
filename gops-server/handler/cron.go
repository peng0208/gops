package handler

import (
	"github.com/labstack/echo"
	"gops/gops-server/model"
	"gops/gops-common"
	"net/http"
	"gops/gops-server/module"
	"fmt"
)

func GetCronTaskList(c echo.Context) error {
	page, size := common.GetPageParams(c)
	env := c.Param("env")
	crontasks, err := model.GetCronTaskList(env, page, size)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &common.ResponseType{0, "", crontasks})
}

func AddCronTask(c echo.Context) error {
	env := c.Param("env")
	crontask := new(model.CronTask)
	crontask.Env = env
	if err := c.Bind(crontask); err != nil {
		return err
	}
	_, err := crontask.Create()

	if err != nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "任务名已存在", nil})
	}

	if crontask.Enable == "1" {
		module.AddCrontabFunc(crontask.CronName, env, crontask.Schedule, crontask.Args)
	}
	return c.JSON(http.StatusOK, &common.ResponseType{1, "新建任务成功", nil})

}

func ChangeCronTask(c echo.Context) error {
	env := c.Param("env")
	cronid := c.Param("cronid")
	crontask := new(model.CronTask)
	crontask.CronID = cronid
	crontask.Env = env
	if err := c.Bind(crontask); err != nil {
		return err
	}
	run, _ := crontask.Status()
	_, err := crontask.Update()
	if err != nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "修改任务失败", nil})
	}
	if run {
		module.DeleteCrontabFunc(crontask.CronName)
		module.AddCrontabFunc(crontask.CronName, env, crontask.Schedule, crontask.Args)
	}
	return c.JSON(http.StatusOK, &common.ResponseType{1, "修改任务成功", nil})

}

func RemoveCronTask(c echo.Context) error {
	env := c.Param("env")
	cronid := c.Param("cronid")
	crontask := new(model.CronTask)
	crontask.CronID = cronid
	crontask.Env = env

	run, _ := crontask.Status()
	_, err := crontask.Delete()
	if err != nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "删除任务失败", nil})
	}
	if run {
		module.DeleteCrontabFunc(crontask.CronName)
	}
	return c.JSON(http.StatusOK, &common.ResponseType{1, "删除任务成功", nil})

}

func ChangeCronStatusTask(c echo.Context) error {
	env := c.Param("env")
	cronid := c.Param("cronid")
	crontask := new(model.CronTask)
	crontask.CronID = cronid
	crontask.Env = env

	if err := c.Bind(crontask); err != nil {
		return err
	}
	fmt.Println(crontask.CronName)
	run, _ := crontask.Status()
	_, err := crontask.UpdateStatus()
	if err != nil {
		return c.JSON(http.StatusOK, &common.ResponseType{2, "修改任务失败", nil})
	}

	if crontask.Enable == "1" {
		module.AddCrontabFunc(crontask.CronName, env, crontask.Schedule, crontask.Args)
	} else {
		if run {
			module.DeleteCrontabFunc(crontask.CronName)
		}
	}
	return c.JSON(http.StatusOK, &common.ResponseType{1, "修改任务成功", nil})

}
