package model

import (
	"gops/gops-common"
	"strings"
	"fmt"
)

type ConfFile struct {
	ConfId   string `json:"confid"`
	ConfName string `json:"confname"`
	Path     string `json:"path"`
	Content  string `json:"content"`
}

func GetConfFileList(env string, page, size int) (*common.PageDataType, error) {
	prefix := fmt.Sprintf("%s/%s/%s", "/conf", env, "file")
	rows, ec := common.EtcdGetPrefix(prefix)
	if ec != nil {
		return nil, ec
	}
	count := len(rows) / 3

	fileMap := make(map[string]*ConfFile)
	for _, row := range rows {
		s := strings.Split(row["key"], "/")
		v := row["value"]
		id := s[4]
		key := s[5]

		if fileMap[id] == nil {
			fileMap[id] = &ConfFile{ConfId: id}
		}

		switch key {
		case "path":
			fileMap[id].Path = v
		case "name":
			fileMap[id].ConfName = v
		case "content":
			fileMap[id].Content = v
		}
	}

	fileSlice := make([]*ConfFile, count)
	var fi int = 0
	for _, f := range fileMap {
		fileSlice[fi] = f
		fi ++
	}

	pages := count/size + 1
	data := &common.PageDataType{
		Total:    count,
		Page:     page,
		PageSize: size,
		Pages:    pages,
		Items:    fileSlice,
	}
	return data, nil
}

func (f *ConfFile) Create(env string) (bool, error) {
	confid := common.StampNanoString()
	name := fmt.Sprintf("%s/%s/%s/%s/%s", "/conf", env, "file", confid, "name")
	path := fmt.Sprintf("%s/%s/%s/%s/%s", "/conf", env, "file", confid, "path")
	content := fmt.Sprintf("%s/%s/%s/%s/%s", "/conf", env, "file", confid, "content")
	if _, err := common.EtcdPut(name, f.ConfName); err != nil {
		return false, err
	}
	if _, err := common.EtcdPut(path, f.Path); err != nil {
		return false, err
	}
	if _, err := common.EtcdPut(content, f.Content); err != nil {
		return false, err
	}
	return true, nil
}

func (f *ConfFile) Update(env string) (bool, error) {

	name := fmt.Sprintf("%s/%s/%s/%s/%s", "/conf", env, "file", f.ConfId, "name")
	//path := fmt.Sprintf("%s/%s/%s/%s", "/conf", env, f.ConfId, "path")
	content := fmt.Sprintf("%s/%s/%s/%s/%s", "/conf", env, "file", f.ConfId, "content")
	if _, err := common.EtcdPut(name, f.ConfName); err != nil {
		return false, err
	}
	//if _, err := common.EtcdPut(path, f.Path); err != nil {
	//	return false, err
	//}
	if _, err := common.EtcdPut(content, f.Content); err != nil {
		return false, err
	}
	return true, nil
}

func (f *ConfFile) Delete(env string) (bool, error) {
	prefix := fmt.Sprintf("%s/%s/%s/%s", "/conf", env, "file", f.ConfId)
	if _, err := common.EtcdDeletePrefix(prefix); err != nil {
		return false, err
	}
	return true, nil
}
