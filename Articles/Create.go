package Articles

import (
	"AICommunity/Authorization"
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
)

// CreateArticle POST 创建文章 需要认证
func CreateArticle(ctx *gin.Context) {
	db := database.GetDB()
	title := ctx.PostForm("title")
	body := ctx.PostForm("body")

	user, _ := ctx.Get("user")
	userid := user.(Authorization.Users).ID

	//	 这里好像不需要做验证
	var article Articles
	// title和body的非空校验
	if title == "" && body == "" {
		ResponseWithNoData(ctx, 1009)
		return
	}

	//	一个简单的标题查重验证
	db.Where("title = ?", title).First(&article)
	if article.ID != 0 {
		ResponseWithNoData(ctx, 1008)
		return
	}

	newArticle := Articles{
		Creator: userid,
		Title:   title,
		Body:    body,
		Like:    0,
	}
	db.Create(&newArticle)

	Response(ctx, 0, ToArticleDto(newArticle), StatusMsg(0))
}
