package Authorization

import (
	"AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
)

// TODO:改进APIToken 验证方式

func AuthMiddleware() gin.HandlerFunc {
	// 进行权限验证
	return func(ctx *gin.Context) {
		// 获取 Authorization Header
		tokenString := ctx.GetHeader("Authorization")

		// validate token format
		if tokenString == "" {
			cookie, err := ctx.Cookie("token")

			if err != nil || cookie == "" {
				println("验证：无Token，返回")
				Responses.NotAuthorized(ctx)
				return
			} else {
				tokenString = cookie
			}
		} else {
			tokenString = tokenString[7:]
		}

		var userId uint
		// 分离判断条件，用来做API认证接口
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			println("验证：Token无效/解析错误，返回")
			Responses.NotAuthorized(ctx)
			return
		}
		// 通过验证
		// 获取token中的userid
		userId = claims.UserId

		DB := database.GetDB()
		var user Users
		DB.First(&user, userId)
		// 验证用户是否存在
		if userId == 0 {
			Responses.NotAuthorized(ctx)
			return
		}
		// 用户存在 把user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
