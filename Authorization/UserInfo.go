package Authorization

import (
	. "AICommunity/Responses"
	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	Response(ctx, 0, ToUserDto(user.(Users)), StatusMsg(0))
}
