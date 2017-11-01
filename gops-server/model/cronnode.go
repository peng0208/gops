package model

import (
	"gops/gops-common"
	"strings"
)

type CronNode struct {
	NodeName string `json:"nodename"`
	LoadStat string `json:"loadstat"`
}

func GetCronNodeList(env string) ([]*CronNode, error) {
	prefix := "/cron/" + env
	rows, err := common.EtcdGetPrefix(prefix)
	nodes := make([]*CronNode, len(rows))

	for i, row := range rows {
		name := strings.Split(row["key"], "/")[3]
		loadStat := row["value"]
		nodes[i] = &CronNode{
			name,
			loadStat,
		}
	}
	return nodes, err
}
