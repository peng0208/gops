package model

import (
	"gops/gops-common"
	"strconv"
)

type CronTask struct {
	CronID   string `json:"cronid"`
	CronName string `json:"cronname"`
	Env      string `json:"env"`
	Args     string `json:"args"`
	Schedule string `json:"schedule"`
	Enable   string `json:"enable"`
	Ctime    int64  `json:"ctime"`
	Mtime    int64  `json:"mtime"`
	Remark   string `json:"remark"`
}

func GetCronTaskList(env string, page, size int) (*common.PageDataType, error) {
	rows, ec := common.MysqlQuery("select count(1) as count from cron_task where env=?;",
		env)
	if ec != nil {
		return nil, ec
	}
	count, _ := strconv.Atoi(rows[0]["count"])
	pages := count/size + 1

	result, er := common.MysqlQuery(
		"select id,name,args,schedule,enable,ctime,mtime,remark from cron_task where env=? limit ?,?;",
		env,
		(page-1)*size,
		size,
	)
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

func (t *CronTask) Create() (bool, error) {
	result, err := common.MysqlExecute(
		"insert into cron_task (name,env,args,schedule,ctime,mtime,remark) values(?,?,?,?,?,?,?);",
		t.CronName,
		t.Env,
		t.Args,
		t.Schedule,
		t.Ctime,
		t.Mtime,
		t.Remark,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (t *CronTask) Update() (bool, error) {
	result, err := common.MysqlExecute(
		"update cron_task set args=?,schedule=?,ctime=?,mtime=?,remark=? where id=?;",
		t.Args,
		t.Schedule,
		t.Ctime,
		t.Mtime,
		t.Remark,
		t.CronID,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (t *CronTask) UpdateStatus() (bool, error) {
	result, err := common.MysqlExecute(
		"update cron_task set enable=? where id=?;",
		t.Enable,
		t.CronID,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (t *CronTask) Status() (bool, error) {
	result, err := common.MysqlQuery(
		"select enable from cron_task where id=?;",
		t.CronID,
	)
	if err != nil {
		return false, err
	}

	var run bool
	for _, row := range result {
		if row["enable"] == "1" {
			run = true
		}
	}

	return run, nil
}

func (t *CronTask) Delete() (bool, error) {
	result, err := common.MysqlExecute(
		"delete from cron_task where id=?;",
		t.CronID,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}
