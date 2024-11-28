package repository

import (
	"context"
	"projectName/pkg/log"

	"gorm.io/driver/mysql"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(logger *log.Logger, db *gorm.DB, rdb *redis.Client) *Repository {
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}
func NewDb(conf *viper.Viper) *gorm.DB {
	// TODO: init db
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db

}

func NewRedisDb(conf *viper.Viper) *redis.Client {
	// TODO: init db
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("data.redis.addr"),
		Password: conf.GetString("data.redis.password"), // no password set
		DB:       conf.GetInt("data.redis.db"),          // use default DB
	})

	// 测试连接
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	return rdb

}
