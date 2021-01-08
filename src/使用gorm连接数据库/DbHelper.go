package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DbHelper *gorm.DB

func init() {
	DbHelper, _ = gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	DbHelper.LogMode(true)
	DbHelper.DB().SetMaxIdleConns(10)
	DbHelper.DB().SetMaxOpenConns(100)
	DbHelper.DB().SetConnMaxLifetime(time.Hour)
}
