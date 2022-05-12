package Articles

import (
	"AICommunity/Authorization"
	"AICommunity/Favorite"
	"AICommunity/Like"
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
	article := r.Group("/api/article")
	{
		article.POST("/publish", Authorization.AuthMiddleware(), CreateArticle)
		article.GET("/info", Authorization.NoBlockAuth(), Info)
		article.POST("/info", Authorization.AuthMiddleware(), ChangeInfo)
		article.GET("/list", Authorization.NoBlockAuth(), DefaultList)
		article.DELETE("/info", Authorization.AuthMiddleware(), DeleteArticle)
		article.GET("/search", Authorization.NoBlockAuth(), SearchList)
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
	Favorite int    `json:"favorite"`
}

func ToArticleDto(article Articles, user uint) gin.H {
	return gin.H{
		"articleId":     article.ID,
		"title":         article.Title,
		"body":          article.Body,
		"category":      article.Category,
		"like":          Like.GetLikeCountRPC(article.ID),
		"isLike":        Like.IsLikedRPC(article.ID, user),
		"favorite":      Favorite.GetFavoriteCountRPC(article.ID),
		"isFavorite":    Favorite.IsFavoriteRPC(article.ID, user),
		"category_name": CATEGORY[article.Category],
	}
}

func ToListArticleDto(articles []Articles, user uint) []gin.H {
	var data []gin.H
	for _, article := range articles {
		data = append(data, gin.H{
			"articleId":     article.ID,
			"title":         article.Title,
			"body":          article.Body,
			"category":      article.Category,
			"like":          Like.GetLikeCountRPC(article.ID),
			"isLike":        Like.IsLikedRPC(article.ID, user),
			"favorite":      Favorite.GetFavoriteCountRPC(article.ID),
			"isFavorite":    Favorite.IsFavoriteRPC(article.ID, user),
			"category_name": CATEGORY[article.Category],
		})
	}

	return data
}
