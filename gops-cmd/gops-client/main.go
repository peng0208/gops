package main

import (
	"flag"
	"runtime"
	"gops/gops-common"
	"net"
	"net/rpc"
	"fmt"
	"gops/gops-client"
	"errors"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configFile := flag.String(
		"c",
		"/Users/peng/Desktop/golang/src/gops/gops-doc/gops-client.conf",
		"config file path",
	)
	flag.Parse()
	common.ParseConfigFile(configFile)
	if common.GetConfig().Server.Env == "" {
		common.Logger().Fatal("Env error")
	}
	common.InitEtcd3()
}

func main() {
	ConfServer()
	if common.GetConfig().Server.Cron {
		CronServer()
	}

	select {}
}

type Server int

func (h *Server) Ping(arg string, reply *bool) error {
	*reply = true
	return errors.New(arg)
}


func CronServer() {
	common.Logger().Println("启动任务调度服务")

	rpcServer := rpc.NewServer()
	rpcServer.Register(new(Server))
	rpcServer.Register(new(gops_client.CronTask))

	addr := fmt.Sprintf("%s:%d", common.GetConfig().Server.Host, common.GetConfig().Server.Port)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		common.Logger().Fatal("Listen error:", e)
	}
	go rpcServer.Accept(l)

	// 服务注册
	go gops_client.NodePortInfo()
	go gops_client.NodeLoadInfo()
}

func ConfServer() {
	common.Logger().Println("启动配置监听服务")

	filePrefix := fmt.Sprintf("%s/%s/%s", "/conf", common.GetConfig().Server.Env, "file")
	keyPrefix := fmt.Sprintf("%s/%s/%s", "/conf", common.GetConfig().Server.Env, "key")

	go gops_client.Watcher(filePrefix, keyPrefix)
}
