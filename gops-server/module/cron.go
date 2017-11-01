package module

import (
	"github.com/jakecoffman/cron"
	"gops/gops-common"
	_ "gops/gops-server/model"
	"strconv"
	"strings"
	"time"
)

var Crontab *cron.Cron

func leastCPUUsed(env string) string {
	prefix := "/cron/" + env
	rows, err := common.EtcdGetPrefix(prefix)
	if err != nil {
		panic(err)
	}

	var minLoad float64
	var minKey string
	var minNode string
	for i, row := range rows {
		n, _ := strconv.ParseFloat(row["value"], 64)
		if i == 0 || n <= minLoad {
			minLoad = n
			minKey = row["key"]
		}
	}
	if minKey != "" {
		minNode = strings.Split(minKey, "/")[3]
	}
	return minNode

}

func InitCrontab() {
	common.Logger().Println("初始化Cron服务")
	time.Sleep(2 * time.Second)

	Crontab = cron.New()
	Crontab.Start()
	rows, err := common.MysqlQuery(
		"select name,env,args,schedule from cron_task where enable=1;")
	if err != nil {
		common.Logger().Printf("初始化Crontab失败: %s", err)
		panic(err)
	}

	for _, row := range rows {
		schedule := row["schedule"]
		name := row["name"]
		args := row["args"]
		env := row["env"]

		AddCrontabFunc(name, env, schedule, args)
	}
}

func AddCrontabFunc(name, env, schedule, args string) {
	Crontab.AddFunc(schedule, func() { go execRemoteCommand(env, args) }, name)
}

func DeleteCrontabFunc(name string) {
	Crontab.RemoveJob(name)
}

func execRemoteCommand(env, args string) {
	const function = "CronTask.ExecShell"

	next := leastCPUUsed(env)
	if next != "" {
		res, err := RPCAsyncCall(next, function, args)
		if err != nil {
			common.Logger().Printf("%s[%s:%s]: Failed, %s", function, next, args, err)
			return
		}
		common.Logger().Printf("%s[%s:%s]: Success, %s", function, next, args, res)
		return
	}
	common.Logger().Printf("远程RPC节点: %s未找到", env)
}
