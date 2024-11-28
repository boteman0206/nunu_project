package params

type CommonParam struct {
	Token string `form:"token" json:"token" binding:"required"`
}

// 登陆的参数
type LoginParams struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginOutParams struct {
	CommonParam
}

// 注册的参数
type RegisterParams struct {
	Username string `form:"username" bindinging:"required"`
	Password string `form:"password" binding:"required"`
}

// 变更密码的参数
type ChangeParams struct {
	CommonParam
	Username    string `form:"username" binding:"required"`
	OldPassword string `form:"old_password" binding:"required"`
	NewPassword string `form:"new_password" binding:"required"`
}
