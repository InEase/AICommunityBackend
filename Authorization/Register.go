package Authorization

import (
	. "AICommunity/Responses"
	"AICommunity/database"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	db := database.GetDB()
	var user Users
	// 获取参数
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	gender := ctx.PostForm("gender")
	school := ctx.PostForm("school")
	avatar := ctx.DefaultPostForm("avatar", AvatarGenerator())

	logger.Info("用户注册:", name, password, gender, school, avatar)
	if password == "" {
		ResponseWithNoData(ctx, 1002)
	}
	//TODO: 此处需要对数据进行合法性校验 性别 学校 以及恶意avatar网址注入

	// 查询用户名是否重复
	db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		ResponseWithNoData(ctx, 1006)
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ResponseWithNoData(ctx, 1007)
		return
	}
	// 创建用户
	newUser := Users{
		Name:     name,
		Password: string(hashedPassword),
		Status:   "正常",
		// 保证非负
		Gender: gender,
		Avatar: avatar,
		Desc:   "",
		School: school,
	}
	db.Create(&newUser)

	Response(ctx, 0, gin.H{
		"name":   name,
		"userid": newUser.ID,
		"status": newUser.Status,
		"gender": gender,
		"avatar": avatar,
		"desc":   "",
		"school": school,
	}, StatusMsg(0))
}

func AvatarGenerator() string {
	// TODO: Avatar Generator
	return "Not Implemented"
}
