package server

import (
	"projectName/internal/handler"
	"projectName/internal/middleware"
	"projectName/internal/repository"
	"projectName/pkg/helper/resp"
	"projectName/pkg/log"

	"github.com/gin-gonic/gin"
)

func NewServerHTTP(
	logger *log.Logger,
	repository *repository.Repository,
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
	// 注册变更密码
	r.POST("/register", userHandler.Register)

	// 用户登陆验证
	r.Use(middleware.TokenMiddleware(repository))

	r.POST("/loginOut", userHandler.LoginOut)

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
