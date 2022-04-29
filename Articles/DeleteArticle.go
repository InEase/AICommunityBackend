package Articles

import (
	"AICommunity/Authorization"
	"AICommunity/Responses"
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
)

func DeleteArticle(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	db := database.GetDB()

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

	if user.(Authorization.Users).ID != article.Creator {
		ResponseWithNoData(ctx, "你不是这篇文字的作者哦")
		return
	}

	db.Delete(&Articles{}, article.ID)
	Responses.Response(
		ctx, 0, nil, "完成",
	)
}
