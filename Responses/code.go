package Responses

var dict = map[int]string{
	0:    "成功",
	1000: "无权限",
	1001: "请求方法错误",
	1002: "缺少参数",
	1003: "参数错误",
	1004: "服务端错误 未找到所要求的数据",
	1005: "权限验证错误",
	1006: "注册：用户已存在 无法重复注册",
	1007: "内部服务器错误：加密错误",
	1008: "登录：无效用户名",
	1009: "文章标题或body为空",
}

func StatusMsg(code int) string {
	return dict[code]
}
