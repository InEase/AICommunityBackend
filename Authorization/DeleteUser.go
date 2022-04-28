package Authorization

import (
	"AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	db := database.GetDB()
	db.Delete(&Users{}, user.(Users).ID)
	Responses.Response(
		ctx, 0, gin.H{
			"userid": user.(Users).ID,
			"name":   user.(Users).Name,
		}, Responses.StatusMsg(0),
	)
}
