package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type T1 struct { // T1 -> t1
	// Id -> id  驼峰-》蛇形
	Id   string `gorm:"primaryKey"`
	Name string
}

var db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:Gzdx123456@@tcp(47.109.189.135:3306)/test?charset=utf8"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	println(db)
}
func main() {
	var t1 T1
	tx := db.Take(&t1)
	fmt.Println(t1) // result
	fmt.Println(tx.RowsAffected, "====", tx.Error, "====", tx.Row())
}
