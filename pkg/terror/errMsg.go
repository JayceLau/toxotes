package terror

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "成功"
	message[ErrCommon] = "服务器开小差啦,稍后再来试一试"
	message[ErrUnknown] = "未知错误"
	message[ErrUsernameEmpty] = "用户名为空"
	message[ErrUsernameExist] = "用户名已存在"
	message[ErrUsernameNotExist] = "用户名不存在"
	message[ErrUserNotExist] = "用户不存在"
	message[ErrUsernameOrPasswordWrong] = "用户名或密码错误"
	message[ErrPhoneExist] = "该手机号已存在"
	message[ErrTokenExpire] = "token失效，请重新登录"
	message[ErrTokenGenerate] = "生产token失败"
	message[ErrDatabase] = "数据库错误，请稍后重试"
}

func MapErrMsg(code uint32) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return message[ErrUnknown]
	}
}

func IsCodeErr(code uint32) bool {
	if _, ok := message[code]; ok {
		return true
	} else {
		return false
	}
}
