package Articles

import (
	"AICommunity/Authorization"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"gorm.io/gorm"
)

func RegisterAll(r *gin.Engine) *gin.Engine {
	db := database.InitDB()
	err := db.AutoMigrate(&Articles{})
	if err != nil {
		panic(err)
	}
	article := r.Group("/article")
	{
		article.POST("/publish", Authorization.AuthMiddleware(), CreateArticle)
		article.GET("/info", Info)
		article.POST("/info", Authorization.AuthMiddleware(), ChangeInfo)
		//auth.POST("/login", Login)
		//auth.DELETE("/info", AuthMiddleware(), DeleteUser)
	}
	logger.Trace("成功加载认证 Article 组件")
	return r
}

type Articles struct {
	gorm.Model
	Creator uint   `json:"creator"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	Like    int    `json:"like"`
}

func ToArticleDto(article Articles) gin.H {
	return gin.H{
		"articleId": article.ID,
		"title":     article.Title,
		"body":      article.Body,
		"like":      article.Like,
	}
}
