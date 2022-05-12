package Like

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

	var like Like
	db.Where("Article = ?", article).Where("User = ?", userid).First(&like)
	if like.ID == 0 {
		//	没操作过
		like = Like{
			Model:   gorm.Model{},
			Article: uint(article),
			User:    userid,
			Status:  1,
		}
		db.Create(&like) // 通过数据的指针来创建
		Response(ctx, 0, gin.H{"status": 1}, "点赞成功")
		return
	} else {
		like.Status = 1 - like.Status
		db.Save(&like)
		var msg = ""
		if like.Status == 1 {
			msg = "点赞成功"
		} else {
			msg = "取消点赞成功"
		}
		Response(ctx, 0, gin.H{"status": like.Status}, msg)
		return
	}

}
