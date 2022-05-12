package Like

import (
	"AICommunity/Authorization"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"gorm.io/gorm"
)

func RegisterAll(r *gin.Engine) *gin.Engine {
	db := database.InitDB()
	err := db.AutoMigrate(&Like{})
	if err != nil {
		panic(err)
	}
	r.POST("/api/article/like", Authorization.AuthMiddleware(), HandleLike)
	logger.Trace("成功加载认证 Like 组件")
	return r
}

type Like struct {
	gorm.Model
	Article uint `json:"article"`
	User    uint `json:"user"`
	Status  uint `json:"status"`
}

func GetLikeCountRPC(article uint) int64 {
	db := database.GetDB()

	var like []Like
	result := db.Where("Article = ?", article).Find(&like)
	return result.RowsAffected
}

func IsLikedRPC(article uint, user uint) uint {
	db := database.GetDB()

	var like Like
	db.Where("Article = ?", article).Where("User = ?", user).First(&like)
	if like.ID == 0 {
		return 0
	} else {
		return like.Status
	}

}
