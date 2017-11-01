package gops_client

import (
	"gops/gops-common"
	"fmt"
	"github.com/shirou/gopsutil/load"
	"strconv"
	"time"
	"runtime"
	"github.com/shirou/gopsutil/cpu"
)

const (
	INTERVAL       = 1 * time.Second
	TTL      int64 = 3
)

func NodePortInfo() {
	t := time.NewTicker(INTERVAL)
	for {
		select {
		case <-t.C:
			prefix := "/register/nodes/" + common.GetConfig().Server.Host
			port := strconv.Itoa(common.GetConfig().Server.Port)
			if _, err := common.EtcdPutWithLeast(prefix, port, TTL); err != nil {
				common.Logger().Print(err)
			}
		}
	}
}

func NodeLoadInfo() {
	cpus := runtime.NumCPU()

	t := time.NewTicker(INTERVAL)
	prefix := fmt.Sprintf(
		"%s/%s/%s",
		"/cron",
		common.GetConfig().Server.Env,
		common.GetConfig().Server.Host,
	)

	var (
		loadStat *load.AvgStat
		percent  []float64
		percent1 float64
		loadP1  float64
		loadP    string
	)

	for {
		select {
		case <-t.C:
			loadStat, _ = load.Avg()
			percent, _ = cpu.Percent(1*time.Second, false)
			percent1 = percent[0]
			loadP1 = loadStat.Load1 / float64(cpus) * percent1

			loadP = strconv.FormatFloat(loadP1, 'f', 2, 64)
			if _, err := common.EtcdPutWithLeast(prefix, loadP, TTL); err != nil {
				common.Logger().Print(err)
			}
		}
	}
}
