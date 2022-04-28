package main

import (
	"AICommunity/Articles"
	"AICommunity/Authorization"
	"AICommunity/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))
	r = RegisterAll(r)
	panic(r.Run(":" + viper.GetString("server.port")))
}

func RegisterAll(r *gin.Engine) *gin.Engine {
	database.InitDB()
	Authorization.RegisterAll(r)
	Articles.RegisterAll(r)
	//TaskList.RegisterAll(db, r)
	//Layout.RegisterAll(r)
	return r
}

// InitConfig 初始化全局设置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
