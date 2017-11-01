package common

import (
	"time"
	"strconv"
	"github.com/labstack/echo"
)

type ResponseType struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageDataType struct {
	Total    int                 `json:"total"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
	Pages    int                 `json:"pages"`
	Items    interface{} `json:"items"`
}

func StampInt64() int64 {
	return time.Now().Unix()
}

func StampNanoInt64() int64 {
	return time.Now().UnixNano() / 1000
}

func StampString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func StampNanoString() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000, 10)
}

func StampToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToStamp(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func GetPageParams(c echo.Context) (int, int) {
	var page int = 1
	var size int = 20

	p := c.QueryParam("page")
	s := c.QueryParam("page_size")
	if p != "" {
		page, _ = strconv.Atoi(p)
	}
	if s != "" {
		size, _ = strconv.Atoi(s)
	}
	return page, size
}
