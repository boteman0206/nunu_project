package params

// 登陆的参数
type LoginParams struct {
	ID int64 `form:"id" bind:"required"`
}
