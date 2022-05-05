package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// 因为是本地应用，所以不需要对数据库进行加密，但是云端的多人数据库需要
	// Changed to Using SQLite
	location := viper.GetString("datasource.location")
	database := viper.GetString("datasource.database")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open(location+database), &gorm.Config{
		Logger: newLogger,
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
