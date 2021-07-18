package errcode

var zhCN = map[int]string{
	ErrNotLogin:               "未登录",
	ErrInvalidParam:           "参数无效",
	ErrDB:                     "数据库错误",
	ErrServer:                 "服务器错误",
	ErrNoRight:                "没有权限",
	ErrRegisterParamInvalid:   "注册参数错误",
	ErrRegisterUsernameExists: "用户名已经注册",
	ErrUnknown:                "发生未知错误，请联系管理员",
	ErrInvalidToken:           "登录令牌无效",
	ErrTokenExpired:           "登录令牌已失效",
	ErrUserNotExists:          "用户不存在",
	ErrLoginWrongPassword:     "用户名或者密码错误",
}
