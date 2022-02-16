package dbops

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)


var (
	dbConn *sql.dbops
	err error
)

func init() {
	dbConn, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}