package model

import (
	"gops/gops-common"
	"strconv"
)

type UserPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ctime    string `json:"ctime"`
}

type User struct {
	UserId   string `json:"userid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Ctime    int64  `json:"ctime"`
	Mtime    int64  `json:"mtime"`
}

func GetUserPassword(user string) (*UserPassword, error) {
	rows, err := common.MysqlQuery(
		"select username, password, ctime from user where username =?;", user)
	if len(rows) == 0 {
		return nil, nil
	}
	row := rows[0]
	passwordInfo := &UserPassword{
		row["username"],
		row["password"],
		row["ctime"],
	}
	return passwordInfo, err
}

func GetUserList(page, size int) (*common.PageDataType, error) {
	rows, ec := common.MysqlQuery("select count(1) as count from user;")
	if ec != nil {
		return nil, ec
	}
	count, _ := strconv.Atoi(rows[0]["count"])
	pages := count/size + 1

	result, er := common.MysqlQuery(
		"select id,username,nickname,ctime from user limit ?,?;",
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

func (u *User) Create() (bool, error) {
	result, err := common.MysqlExecute(
		"insert into user(username,nickname,password,ctime) values(?,?,?,?);",
		u.Username,
		u.Nickname,
		u.Password,
		u.Ctime,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (u *User) Update() (bool, error) {
	result, err := common.MysqlExecute(
		"update user set nickname=?,mtime=? where id =?;",
		u.Nickname,
		u.Mtime,
		u.UserId,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (u *User) Delete() (bool, error) {
	result, err := common.MysqlExecute(
		"delete from user where id=?;",
		u.UserId,
	)
	if err != nil {
		return false, err
	}
	return result, nil
}