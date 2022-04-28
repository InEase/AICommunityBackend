package Responses

import (
	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, code int, data gin.H, msg string) {
	ctx.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func NotAuthorized(ctx *gin.Context) {
	ResponseWithNoData(ctx, 1000)
}

func ResponseWithNoData(ctx *gin.Context, code int) {
	ctx.JSON(200, gin.H{
		"code": code,
		"data": nil,
		"msg":  StatusMsg(code),
	})
	ctx.Abort()
}
