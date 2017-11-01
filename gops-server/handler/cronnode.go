package handler

import (
	"gops/gops-server/model"
	"github.com/labstack/echo"
	"net/http"
	"gops/gops-common"
)

func GetCronNodeList(c echo.Context) error {
	env := c.Param("env")
	nodes, err := model.GetCronNodeList(env)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &common.ResponseType{0, "", nodes})
}
