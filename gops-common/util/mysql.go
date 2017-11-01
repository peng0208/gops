package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const (
	MYSQL_DRIVER       = "mysql"
	NONE         int64 = 0
)

type MysqlConn struct {
	Host    string
	Port    int
	User    string
	Pwd     string
	Db      string
	MaxOpen int
	MaxIdle int
}

func (c *MysqlConn) Conn() (*sql.DB, error) {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.User,
		c.Pwd,
		c.Host,
		c.Port,
		c.Db,
	)
	db, err := sql.Open(MYSQL_DRIVER, connString)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(c.MaxOpen)
	db.SetMaxIdleConns(c.MaxIdle)
	db.SetConnMaxLifetime(30)
	return db, nil
}

// Read from db, return data
func MysqlQueryDb(tdb *sql.DB, sqlString string, args ...interface{}) ([]map[string]string, error) {
	stmt, ep := tdb.Prepare(sqlString)
	if ep != nil {
		return nil, ep
	}
	defer stmt.Close()

	rows, eq := stmt.Query(args...)
	if eq != nil {
		return nil, eq
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	cols := make([]sql.RawBytes, len(columns))
	scans := make([]interface{}, len(columns))
	for v := range cols {
		scans[v] = &cols[v]
	}

	var result []map[string]string
	for rows.Next() {
		_ = rows.Scan(scans...)
		row := make(map[string]string)
		for i, col := range cols {
			row[columns[i]] = string(col)
		}
		result = append(result, row)
	}
	return result, nil
}

// Write to db, return number of rows
func MysqlExecuteDb(tdb *sql.DB, sqlString string, args ...interface{}) (bool, error) {
	stmt, ep := tdb.Prepare(sqlString)
	if ep != nil {
		return false, ep
	}
	defer stmt.Close()

	rows, ee := stmt.Exec(args...)
	if ee != nil {
		return false, ee
	}

	_, er := rows.RowsAffected()
	if er != nil {
		return false, er
	}
	return true, nil
}
