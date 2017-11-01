package model

import (
	"gops/gops-common"
	"strconv"
)

type CronTag struct {
	TagId   string `json:"tagid"`
	Tagname string `json::tagname`
	Remark  string `json:"remark"`
}

func GetCronTagList(page, size int) (*common.PageDataType, error) {
	rows, ec := common.MysqlQuery("select count(1) as count from cron_tag;")
	if ec != nil {
		return nil, ec
	}
	count, _ := strconv.Atoi(rows[0]["count"])
	pages := count/size + 1

	result, er := common.MysqlQuery("select * from cron_tag;")
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

func (t *CronTag) Create() (bool, error) {
	result, err := common.MysqlExecute(
		"insert into cron_tag(tagname,remark) values(?,?);",
		t.Tagname,
		t.Remark,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (t *CronTag) Delete() (bool, error) {
	result, err := common.MysqlExecute(
		"delete from cron_tag where id=?;",
		t.TagId,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}
