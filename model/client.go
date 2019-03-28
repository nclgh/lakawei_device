package model

import (
	"github.com/jinzhu/gorm"
	"github.com/nclgh/lakawei_scaffold/mysql"
)

var (
	lakaweiDb *gorm.DB
)

func Init() {
	lakaweiDb = mysql.GetMysqlDB("lakawei")
	//lakaweiDb = mysql.GetMysqlDB("lakawei").Debug()
	lakaweiDb.SingularTable(true)
}

func GetLakaweiDb() *gorm.DB {
	return lakaweiDb
}