package Articles

import (
	"AICommunity/Authorization"
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	db := database.GetDB()
	// 获取参数
	articleid, exists := ctx.GetQuery("articleid")
	if !exists {
		ResponseWithNoData(ctx, "缺失article id参数")
		return
	}

	// Start Process
	var article Articles
	db.Where("id = ?", articleid).First(&article)
	if article.ID == 0 {
		ResponseWithNoData(ctx, "找不到文章")
		return
	}

	Response(ctx, 0, ToArticleDto(article), "完成")
}

func ChangeInfo(ctx *gin.Context) {
	db := database.GetDB()
	// 获取参数
	articleid, exists := ctx.GetPostForm("articleid")
	if !exists {
		ResponseWithNoData(ctx, "缺失article id参数")
		return
	}

	var article Articles
	db.Where("id = ?", articleid).First(&article)
	if article.ID == 0 {
		ResponseWithNoData(ctx, "找不到文章")
		return
	}

	// 鉴权
	user, _ := ctx.Get("user")
	if user.(Authorization.Users).ID != article.Creator {
		ResponseWithNoData(ctx, "你不是这篇文字的作者哦")
		return
	}

	var updated = false
	title, _ := ctx.GetPostForm("title")
	if title != "" {
		article.Title = title
		updated = true
	}

	body, _ := ctx.GetPostForm("body")
	if body != "" {
		article.Body = body
		updated = true
	}

	if updated {
		db.Save(&article)
	}
	Response(ctx, 0, ToArticleDto(article), "完成")
}
