package Favorite

import (
	"AICommunity/Authorization"
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func HandleLike(ctx *gin.Context) {
	db := database.GetDB()

	articleId := ctx.DefaultPostForm("articleid", "")
	if articleId == "" {
		ResponseWithNoData(ctx, "未传articleId")
		return
	}

	article, err := strconv.Atoi(articleId)
	if err != nil {
		ResponseWithNoData(ctx, "articleId不是整数")
		return
	}

	user, _ := ctx.Get("user")
	userid := user.(Authorization.Users).ID

	var favorite Favorite
	db.Where("Article = ?", article).Where("User = ?", userid).First(&favorite)
	if favorite.ID == 0 {
		//	没操作过
		favorite = Favorite{
			Model:   gorm.Model{},
			Article: uint(article),
			User:    userid,
			Status:  1,
		}
		db.Create(&favorite) // 通过数据的指针来创建
		Response(ctx, 0, gin.H{"status": 1}, "收藏成功")
		return
	} else {
		favorite.Status = 1 - favorite.Status
		db.Save(&favorite)
		var msg = ""
		if favorite.Status == 1 {
			msg = "收藏成功"
		} else {
			msg = "取消收藏成功"
		}
		Response(ctx, 0, gin.H{"status": favorite.Status}, msg)
		return
	}

}
