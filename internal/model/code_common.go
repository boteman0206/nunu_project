package model

// 业务错误码定义
const (
	CodePasswordErr      = 1001 // 密码错误
	CodeUserNameNotExist = 1002 // 用户名不存在

)

// 业务错误信息定义
const (
	ErrPasswordErr      = "密码错误"
	ErrUserNameNotExist = "用户名不存在"
)

var (

	// 业务错误信息映射
	ErrMsgMap = map[int]string{
		CodePasswordErr:      ErrPasswordErr,
		CodeUserNameNotExist: ErrUserNameNotExist,
	}
)
