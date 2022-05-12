package Articles

import (
	"AICommunity/Authorization"
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreateArticle POST 创建文章 需要认证
func CreateArticle(ctx *gin.Context) {
	db := database.GetDB()
	title := ctx.PostForm("title")
	body := ctx.PostForm("body")
	strCategory := ctx.PostForm("category")
	category, err := strconv.Atoi(strCategory)
	if err != nil {
		ResponseWithNoData(ctx, "未传category 或category不是一个整数")
		return
	}

	user, _ := ctx.Get("user")
	userid := user.(Authorization.Users).ID

	//	 这里好像不需要做验证
	var article Articles
	// title和body的非空校验
	if title == "" && body == "" {
		ResponseWithNoData(ctx, "Title 或 body 为空")
		return
	}

	//	一个简单的标题查重验证
	db.Where("title = ?", title).First(&article)
	if article.ID != 0 {
		ResponseWithNoData(ctx, "标题重复")
		return
	}

	// 确认category id在范围内
	if !IfInSlice(category, CategoryKeys) {
		ResponseWithNoData(ctx, "未知category")
		return
	}

	newArticle := Articles{
		Creator:  userid,
		Title:    title,
		Body:     body,
		Like:     0,
		Category: category,
	}
	db.Create(&newArticle)

	Response(ctx, 0, ToArticleDto(newArticle, userid), StatusMsg(0))
}
