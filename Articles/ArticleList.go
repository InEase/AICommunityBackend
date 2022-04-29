package Articles

import (
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DefaultList(ctx *gin.Context) {
	db := database.GetDB()

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
	db.Limit(limit).Offset(offset).Select("ID").Find(&articles)

	var ids []uint
	for _, article := range articles {
		ids = append(ids, article.ID)
	}

	Response(ctx, 0, ids, "完成")
}
