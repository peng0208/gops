package model

import (
	"gops/gops-common"
	"strconv"
)

type ConfTag struct {
	TagId   string `json:"tagid"`
	Tagname string `json::tagname`
	Remark  string `json:"remark"`
}

func GetConfTagList(page, size int) (*common.PageDataType, error) {
	rows, ec := common.MysqlQuery("select count(1) as count from conf_tag;")
	if ec != nil {
		return nil, ec
	}
	count, _ := strconv.Atoi(rows[0]["count"])
	pages := count/size + 1

	result, er := common.MysqlQuery("select * from conf_tag;")
	if er != nil {
		return nil, er
	}

	data := &common.PageDataType{
		Total:    count,
		Page:     page,
		PageSize: size,
		Pages:    pages,
		Items:    result,
	}
	return data, nil
}

func (t *ConfTag) Create() (bool, error) {
	result, err := common.MysqlExecute(
		"insert into conf_tag(tagname,remark) values(?,?);",
		t.Tagname,
		t.Remark,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (t *ConfTag) Delete() (bool, error) {
	result, err := common.MysqlExecute(
		"delete from conf_tag where id=?;",
		t.TagId,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}
