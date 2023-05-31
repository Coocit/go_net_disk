package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/net_disk_db?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		panic("数据库连接错误!")
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql,err:" + err.Error())
	}
}

func DBConn() *sql.DB {
	return db
}
