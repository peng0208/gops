package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gops/gops-common/util"
)

var db *sql.DB

// Init a global DB interface, create a connection pool
func InitMysql() {
	Logger().Println("初始化Mysql连接")

	config := GetConfig().Database
	conn := &util.MysqlConn{
		User:    config.User,
		Pwd:     config.Password,
		Host:    config.Host,
		Port:    config.Port,
		Db:      config.Db,
		MaxOpen: config.MaxOpen,
		MaxIdle: config.MaxIdle,
	}
	db, _ = conn.Conn()
}

func MysqlQuery(sqlString string, args ...interface{}) ([]map[string]string, error) {
	return util.MysqlQueryDb(db, sqlString, args...)
}

func MysqlExecute(sqlString string, args ...interface{}) (bool, error) {
	return util.MysqlExecuteDb(db, sqlString, args...)
}