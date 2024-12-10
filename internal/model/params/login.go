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
	Username    string `form:"username" bindinging:"required"`
	Password    string `form:"password" binding:"required"`
	PortraitUrl string `form:"portrait_url"`
}

// 变更密码的参数
type ChangeParams struct {
	CommonParam
	Username    string `form:"username" binding:"required"`
	OldPassword string `form:"old_password" binding:"required"`
	NewPassword string `form:"new_password" binding:"required"`
}

type GetFeeInfoParams struct {
	FeedID int64 `form:"feed_id" binding:"required"`
}

type CreateFeedParams struct {
	CommonParam
	Title       string `json:"title" `
	Tags        string `json:"tags"`
	Description string `json:"description" binding:"required"`
	ImageList   string `json:"image_list"`
}
