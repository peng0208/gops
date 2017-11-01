package module

import (
	"gops/gops-common"
	"fmt"
	"strings"
	"net/rpc"
	"time"
	"sync"
	"errors"
)

type RPCPool struct {
	Pool map[string]*rpc.Client
	Lock sync.RWMutex
}

var (
	rpcConnPool        RPCPool
	rpcConnStringCache map[string]string
)

const INTERVAL = 2 * time.Second

func InitRPCPool() {
	common.Logger().Println("初始化RPC连接池")

	rpcConnPool.Pool = make(map[string]*rpc.Client)
	rpcConnStringCache = make(map[string]string)
	go RPCConnPool()
}

func RPCConnAddrs() map[string]string {
	prefix := "/register/nodes/"
	values, err := common.EtcdGetPrefix(prefix)
	if err != nil {
		panic(err)
	}

	ports := make(map[string]string, len(values))
	addrs := make(map[string]string, len(values))
	for _, value := range values {
		k := strings.Split(value["key"], "/")[3]
		v := value["value"]
		ports[k] = v
	}

	for i, p := range ports {
		addrs[i] = fmt.Sprintf("%s:%s", i, p)
	}
	return addrs
}

func RPCConnect(host, addr string) {
	rpcConnPool.Lock.Lock()
	defer rpcConnPool.Lock.Unlock()

	conn := rpcConnPool.Pool[host]
	if conn == nil {
		clt, err := rpc.Dial("tcp", addr)
		if err != nil {
			common.Logger().Printf("远程RPC连接失败: %s", err)
			return
		}
		rpcConnPool.Pool[host] = clt
		return
	}

	var online bool
	err := conn.Call("Server.Ping", "", &online)
	if err != nil && !online {
		delete(rpcConnPool.Pool, host)
	}
}

func RPCConnPool() {
	t := time.NewTicker(INTERVAL)
	for {
		select {
		case <-t.C:
			rpcConnStringCache = RPCConnAddrs()
			for host, addr := range rpcConnStringCache {
				go RPCConnect(host, addr)
			}
		}
	}
}

func RPCCall(ip string, function string, args string) (interface{}, error) {
	var rev interface{}

	clt := rpcConnPool.Pool[ip]
	if clt == nil {
		return nil, errors.New(fmt.Sprintf("远程RPC节点: %s连接已关闭", ip))
	}
	err := clt.Call(function, args, &rev)
	return rev, err
}

func RPCAsyncCall(ip string, function string, args string) (interface{}, error) {
	var rev interface{}

	clt := rpcConnPool.Pool[ip]
	if clt == nil {
		return nil, errors.New(fmt.Sprintf("远程RPC节点: %s连接已关闭", ip))
	}
	req := clt.Go(function, args, &rev, nil)
	result := <-req.Done
	return rev, result.Error
}
