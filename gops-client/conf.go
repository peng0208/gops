package gops_client

import (
	"strings"
	"os"
	"gops/gops-common"
)

func Watcher(prefix ...string) {
	for _, p := range prefix {
		go func(p string) {
			for {
				e, k, v := common.EtcdWatchPrefixOnce(p)
				go WatchHandler(e, k, v)
			}
		}(p)
	}
}

func WatchHandler(event, key, value string) {
	prefix := strings.Join(strings.Split(key, "/")[:5], "/")
	suffix := strings.Join(strings.Split(key, "/")[5:], "")
	if event == "PUT" {
		if suffix == "content" {
			filePath, err := common.EtcdGet(prefix + "/path")
			if err != nil {
				common.Logger().Print(err)
				return
			}
			f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			defer f.Close()
			if err != nil {
				common.Logger().Print(err)
				return
			}

			f.WriteString(value + "\n")
		}
	}

	if event == "DELETE" {
		if suffix == "name" {
			filePath, err := common.EtcdGet(prefix + "/path")
			if err != nil {
				common.Logger().Print(err)
				return
			}
			if err := os.Remove(filePath); err != nil {
				common.Logger().Print(err)
				return
			}
			if _, err := common.EtcdDeletePrefix(prefix); err != nil {
				common.Logger().Print(err)
				return
			}
		}
	}
}
