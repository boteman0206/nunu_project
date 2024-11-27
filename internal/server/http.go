package server

import (
	"github.com/gin-gonic/gin"
	"projectName/internal/handler"
	"projectName/internal/middleware"
	"projectName/pkg/helper/resp"
	"projectName/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler *handler.UserHandler,
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

	// 登陆登出
	r.POST("/login", userHandler.Login)
	r.POST("/loginOut", userHandler.Login)

	// 注册变更密码
	r.POST("/register", userHandler.Register)
	r.POST("/changePassword", userHandler.Login)

	// 用户模块相关
	userController := r.Group("/user")
	userController.GET("/getUserById", userHandler.GetUserById)

	// feed帖子相关

	return r
}
