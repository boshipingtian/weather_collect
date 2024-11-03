package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"weather-colly/global"
)

func InitGorm() {
	global.DB = mySQLConnector()
}

func mySQLConnector() *gorm.DB {
	if global.Config.DataBase.Host == "" {
		global.Logger.Infoln("database host is not configured")
		return nil
	}
	connectUrl := global.Config.DataBase.GenerateUrl()
	db, err := gorm.Open(mysql.Open(connectUrl), &gorm.Config{})
	if err != nil {
		global.Logger.Error(err.Error())
	}
	global.Logger.Infoln("connect database success")
	return db
}
