package common

import (
	"github.com/coreos/etcd/clientv3"
	"gops/gops-common/util"

)

var clt *clientv3.Client

// Init a global Etcd rpcclient
func InitEtcd3() {
	Logger().Println("初始化Etcd连接")

	var err error
	config := GetConfig().Etcd
	conn := &util.EtcdConn{
		Host: config.Host,
		Port: config.Port,
	}
	clt, err = conn.Conn()
	if err != nil {
		panic(err)
	}
}

func EtcdGet(keyString string) (string, error) {
	return util.EtcdGet(clt, keyString)
}

func EtcdGetPrefix(prefixString string) ([]map[string]string, error) {
	return util.EtcdGetPrefix(clt, prefixString)
}

func EtcdGetKeysOnly(prefixString string) ([]string, error) {
	return util.EtcdGetKeysOnly(clt, prefixString)
}

func EtcdPut(keyString string, valueString string) (bool, error) {
	return util.EtcdPut(clt, keyString, valueString)
}

func EtcdPutWithLeast(keyString string, valueString string, t int64) (bool, error) {
	return util.EtcdPutWithLeast(clt, keyString, valueString, t)
}

func EtcdWatchOnce(keyString string) (string, string, string) {
	return util.EtcdWatchOnce(clt, keyString)
}

func EtcdWatchPrefixOnce(prefixString string) (string, string, string) {
	return util.EtcdWatchPrefixOnce(clt, prefixString)
}

func EtcdDelete(keyString string) (bool, error) {
	return util.EtcdDelete(clt, keyString)
}

func EtcdDeletePrefix(prefixString string) (bool, error) {
	return util.EtcdDeletePrefix(clt, prefixString)
}
