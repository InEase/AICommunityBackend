package Favorite

import (
	"AICommunity/Authorization"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"gorm.io/gorm"
)

func RegisterAll(r *gin.Engine) *gin.Engine {
	db := database.InitDB()
	err := db.AutoMigrate(&Favorite{})
	if err != nil {
		panic(err)
	}
	r.POST("/api/article/Favorite", Authorization.AuthMiddleware(), HandleLike)
	logger.Trace("成功加载认证 Favorite 组件")
	return r
}

type Favorite struct {
	gorm.Model
	Article uint `json:"article"`
	User    uint `json:"user"`
	Status  uint `json:"status"`
}

func GetFavoriteCountRPC(article uint) int64 {
	db := database.GetDB()

	var favorite []Favorite
	result := db.Where("Article = ?", article).Find(&favorite)
	return result.RowsAffected
}

func IsFavoriteRPC(article uint, user uint) uint {
	db := database.GetDB()

	var favorite Favorite
	db.Where("Article = ?", article).Where("User = ?", user).First(&favorite)
	if favorite.ID == 0 {
		return 0
	} else {
		return favorite.Status
	}

}
