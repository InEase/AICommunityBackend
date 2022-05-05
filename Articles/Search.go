package Articles

import (
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SearchList(ctx *gin.Context) {
	db := database.GetDB()

	title := ctx.DefaultQuery("title", "")
	body := ctx.DefaultQuery("body", "")

	title = "%" + title + "%"
	body = "%" + body + "%"

	temp1, exists := ctx.GetQuery("limit")

	limit := 10
	err := error(nil)
	if exists {
		limit, err = strconv.Atoi(temp1)
		if err != nil {
			ResponseWithNoData(ctx, "limit 请传入数值类型")
			return
		}
	}

	offset := 0
	temp2, exists := ctx.GetQuery("offset")
	if exists {
		offset, err = strconv.Atoi(temp2)
		if err != nil {
			ResponseWithNoData(ctx, "offset 请传入数值类型")
			return
		}
	}

	// Start Process
	var articles []Articles
	usualQuery := db.Limit(limit).Offset(offset).Where("title like ?", title).Where("body like ?", body)
	temp, exists := ctx.GetQuery("category")
	if exists {
		category, parseError := strconv.Atoi(temp)
		if parseError != nil {
			ResponseWithNoData(ctx, "category 请传入数值类型")
			return
		}
		usualQuery.Where("category = ?", category).Find(&articles)
	} else {
		usualQuery.Find(&articles)
	}

	Response(ctx, 0, ToListArticleDto(articles), "完成")
}
