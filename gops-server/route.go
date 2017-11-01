package gops_server

import (
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"gops/gops-server/handler"
	"github.com/labstack/echo/middleware"
)

func Routes(e *echo.Echo) {
	// Login, needn't authentication token
	e.POST("/api/login", handler.CreateToken)

	// Logined, need authentication token
	r := e.Group(
		"/api",
		middleware.JWT([]byte("NDQ4MjUzNEMtRTI4==")),
	)
	r.POST("/logout", handler.ClearToken)

	r.GET("/users", handler.GetUserList)
	r.POST("/users", handler.AddUser)
	r.DELETE("/user/:userid", handler.RemoveUser)
	r.GET("/conftags", handler.GetConfTagList)
	r.POST("/conftags", handler.AddConfTag)
	r.DELETE("/conftag/:tagid", handler.RemoveConfTag)
	r.GET("/conffiles/:env", handler.GetConfFileList)
	r.POST("/conffiles/:env", handler.AddFile)
	r.PUT("/conffile/:env/:confid", handler.ChangeFile)
	r.DELETE("/conffile/:env/:confid", handler.RemoveFile)
	r.GET("/cronnodes/:env", handler.GetCronNodeList)
	r.GET("/crontags", handler.GetCronTagList)
	r.POST("/crontags", handler.AddCronTag)
	r.DELETE("/crontag/:tagid", handler.RemoveCronTag)
	r.GET("/crons/:env", handler.GetCronTaskList)
	r.POST("/crons/:env", handler.AddCronTask)
	r.PUT("/cron/:env/:cronid", handler.ChangeCronTask)
	r.DELETE("/cron/:env/:cronid", handler.RemoveCronTask)
	r.PATCH("/cron/status/:env/:cronid", handler.ChangeCronStatusTask)












}
