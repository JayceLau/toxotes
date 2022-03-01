package terror

// 全局状态

const OK uint32 = 0
const ErrCommon = 100001
const ErrUnknown uint32 = 999999

// 用户模块

const ErrUsernameEmpty uint32 = 101001
const ErrUsernameExist uint32 = 101002
const ErrUsernameNotExist uint32 = 101003
const ErrUserNotExist uint32 = 101004
const ErrUsernameOrPasswordWrong uint32 = 101005
const ErrPhoneExist uint32 = 101006

//ErrUsernameEmpty           = "用户名为空"
//ErrUsernameExist           = "用户名已存在"
//ErrUsernameOrPasswordWrong = "用户名或密码错误"
//ErrPhoneExist              = "该手机号已存在"
//ErrUnknown                 = "未知错误"
