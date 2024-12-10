package server

import (
	"projectName/internal/handler"
	"projectName/internal/middleware"
	"projectName/pkg/helper/resp"
	"projectName/pkg/log"

	"github.com/gin-gonic/gin"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler *handler.UserHandler,
	feedHchander *handler.FeedHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Nunu!",
		})
	})
	r.GET("/ping", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, "successful")
	})

	// 登陆登出
	r.POST("/login", userHandler.Login)
	r.POST("/loginOut", userHandler.LoginOut)

	// 注册变更密码
	r.POST("/register", userHandler.Register)
	r.POST("/changePassword", userHandler.ChangePassword)

	// 用户模块相关
	userController := r.Group("/user")
	userController.GET("/userInfoCenter", userHandler.UserInfoCenter)
	userController.POST("/updateUserInfo", userHandler.UpdateUserInfo)

	// feed帖子相关

	feedController := r.Group("/feed")
	feedController.GET("/feedInfo", feedHchander.GetFeed)
	feedController.POST("/createFeed", feedHchander.CreateFeed)

	return r
}
