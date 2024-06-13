package api

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

const (
	CodeNeedLogin = 2000 + iota
	CodeInvalidToken
)

const (
	CommunityExist = 3000 + iota
	CommunityNotExist
)

var codeMsgMap = map[int]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeInvalidToken: "无效的token",
	CodeNeedLogin:    "需要登录",

	CommunityExist:    "社区已存在",
	CommunityNotExist: "社区不存在",
}

func Code2Msg(code int) string {
	return codeMsgMap[code]
}
