package main

import (
	"os"
	"time"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gops/gops-common"
	"gops/gops-server"
	"flag"
	"fmt"
	"runtime"
	"gops/gops-server/module"
)

const HTTP_TIMEOUT = 60

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configFile := flag.String(
		"c",
		"/Users/peng/Desktop/golang/src/gops/gops-doc/gops-server.conf",
		"config file path",
	)
	flag.Parse()
	common.ParseConfigFile(configFile)
	common.InitMysql()
	common.InitEtcd3()
	module.InitRPCPool()
	module.InitCrontab()
}

func main() {
	config := common.GetConfig().Server

	gen, _ := os.OpenFile(config.ErrorLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	acc, _ := os.OpenFile(config.AccessLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	e := echo.New()
	e.Debug = false
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetPrefix("general")
	e.Logger.SetOutput(gen)
	e.Use(
		middleware.Recover(),
		middleware.Secure(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://gops.com:8080","http://localhost:8080", "http://127.0.0.1:8080"},
			AllowCredentials: true,
			AllowHeaders: []string{"Authorization", "Content-type"},
			AllowMethods: []string{echo.GET, echo.PATCH,echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		}),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}",` +
				`"host":"${host}","method":"${method}","uri":"${uri}","status":${status},` +
				`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
				`"bytes_out":${bytes_out}}` + "\n",
			Output: acc,
		}))

	gops_server.Routes(e)

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	httpServer := &http.Server{
		Addr:         addr,
		ReadTimeout:  HTTP_TIMEOUT * time.Second,
		WriteTimeout: HTTP_TIMEOUT * time.Second,
	}
	e.Logger.Fatal(e.StartServer(httpServer))
}