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
		article.GET("/list", DefaultList)
		article.DELETE("/info", Authorization.AuthMiddleware(), DeleteArticle)
		article.GET("/search", SearchList)
	}
	logger.Trace("成功加载认证 Article 组件")
	return r
}

type Articles struct {
	gorm.Model
	Creator  uint   `json:"creator"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Category int    `json:"category"`
	Like     int    `json:"like"`
}

func ToArticleDto(article Articles) gin.H {
	return gin.H{
		"articleId":     article.ID,
		"title":         article.Title,
		"body":          article.Body,
		"category":      article.Category,
		"like":          article.Like,
		"category_name": CATEGORY[article.Category],
	}
}

func ToListArticleDto(articles []Articles) []gin.H {
	var data []gin.H
	for _, article := range articles {
		data = append(data, gin.H{
			"articleId":     article.ID,
			"title":         article.Title,
			"body":          article.Body,
			"category":      article.Category,
			"like":          article.Like,
			"category_name": CATEGORY[article.Category],
		})
	}

	return data
}
