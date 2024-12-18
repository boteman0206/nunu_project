//go:build wireinject
// +build wireinject

package wire

import (
	"projectName/internal/handler"
	"projectName/internal/repository"
	"projectName/internal/server"
	"projectName/internal/service"
	"projectName/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRedisDb,
	repository.NewRepository,
	repository.NewUserRepository,
	repository.NewFeedRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewFeedService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewFeedHandler,
)

func NewWire(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
