package params

// 登陆的参数
type LoginParams struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 注册的参数
type RegisterParams struct {
	Username string `form:"username" bindinging:"required"`
	Password string `form:"password" binding:"required"`
}

// 变更密码的参数
type ChangePasswordParams struct {
	Username    string `form:"username" binding:"required"`
	OldPassword string `form:"old_password" binding:"required"`
	NewPassword string `form:"new_password" binding:"required"`
}
