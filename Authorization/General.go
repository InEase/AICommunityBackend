package Authorization

import (
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"gorm.io/gorm"
)

func RegisterAll(r *gin.Engine) *gin.Engine {
	db := database.InitDB()
	err := db.AutoMigrate(&Users{})
	if err != nil {
		panic(err)
	}
	auth := r.Group("/user")
	{
		auth.POST("/register", Register)
		auth.POST("/login", Login)
		auth.GET("/info", AuthMiddleware(), Info)
		auth.DELETE("/info", AuthMiddleware(), DeleteUser)
	}
	logger.Trace("成功加载认证 Authorization & 用户 Users 组件")
	return r
}

type Users struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   string `json:"status"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	Desc     string `json:"desc"`
	School   string `json:"school"`
}

func ToUserDto(user Users) gin.H {
	return gin.H{
		"name": user.Name,
	}
}
