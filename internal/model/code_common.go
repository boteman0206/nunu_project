package model

// 业务错误码定义
const (

	// 网络错误
	CodeNetError = 1 // 网络错误
	CodeParamErr = 2 // 参数错误

	CodePasswordErr          = 1001 // 密码错误
	CodeUserNameALreadyExist = 1002 // 用户名已经存在
	CodeUserNameNotExist     = 1003 // 用户名不存在
	CodeOldPasswordErr       = 1004 // 旧密码错误

	TokenExpErr = 2001 // token过期

)

// 业务错误信息定义
var (

	// 业务错误信息映射
	ErrMsgMap = map[int]string{
		CodeNetError:             "网络错误",
		CodeParamErr:             "参数错误",
		CodePasswordErr:          "密码错误",
		CodeUserNameALreadyExist: "用户名已经存在",
		CodeUserNameNotExist:     "用户名不存在",
		TokenExpErr:              "token过期",
		CodeOldPasswordErr:       "旧密码错误",
	}
)
