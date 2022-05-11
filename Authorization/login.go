package Authorization

import (
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	// TODO:多次重复登录销毁上一次token，刷新存活时间
	db := database.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	// Start Process
	var user Users
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		ResponseWithNoData(ctx, "找不到用户")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ResponseWithNoData(ctx, "密码错误")
		return
	}
	// 发放token
	token, err := ReleaseToken(user)
	if err != nil {
		ResponseWithNoData(ctx, "Token生成失败")
		return
	}
	// Set Cookie
	//cookie, err := ctx.Cookie("token")
	//cookie = "NotSet"
	ctx.SetCookie("token", token, 3600, "/", "localhost", false, false)

	//cookie, err := ctx.Cookie("token")
	//if err != nil {
	//	//创建cookie
	//	cookie = "NotSet"
	//}

	// 返回结果
	Response(ctx, 0, gin.H{"token": token}, StatusMsg(0))
}
