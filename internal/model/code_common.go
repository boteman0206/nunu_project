package model

// 业务错误码定义
const (

	// 网络错误
	CodeNetError = 1 // 网络错误
	CodeParamErr = 2 // 参数错误

	CodePasswordErr          = 1001 // 密码错误
	CodeUserNameALreadyExist = 1002 // 用户名已经存在

	CodeUserNameNotExist = 1003 // 用户名不存在

)

// 业务错误信息定义
const (
	MsgNetError = "网络错误"
	MsgParamErr = "参数错误"

	MsgPasswordErr          = "密码错误"
	MsgUserNameAlreadyExist = "用户名已经存在"

	MsgNetErrorNotExist = "用户名不存在"
)

var (

	// 业务错误信息映射
	ErrMsgMap = map[int]string{
		CodeNetError:             MsgNetError,
		CodeParamErr:             MsgParamErr,
		CodePasswordErr:          MsgPasswordErr,
		CodeUserNameALreadyExist: MsgUserNameAlreadyExist,
		CodeUserNameNotExist:     MsgNetErrorNotExist,
	}
)
