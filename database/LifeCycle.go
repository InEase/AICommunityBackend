package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// 因为是本地应用，所以不需要对数据库进行加密，但是云端的多人数据库需要
	// Changed to Using SQLite
	location := viper.GetString("datasource.location")
	database := viper.GetString("datasource.database")
	db, err := gorm.Open(sqlite.Open(location+database), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database , err:" + err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
