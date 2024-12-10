package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"projectName/pkg/log"
	"strings"
	"time"

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

func (r *Repository) SetData(token string, data string, exp int64) error {

	return r.rdb.Set(r.rdb.Context(), token, data, time.Duration(exp)*time.Second).Err()
}

// DelData implements UserRepository.
func (r *Repository) DelData(key []string) error {
	return r.rdb.Del(r.rdb.Context(), key...).Err()

}

// GetData implements UserRepository.
func (r *Repository) GetData(key string, data interface{}) (bool, error) {
	key = strings.TrimSpace(key)
	if key == "" {

		return false, fmt.Errorf("[library:cache]:getDataFromCache -> key is empty")
	}

	// key不存在，返回空字符串
	res, _ := r.rdb.Get(r.rdb.Context(), key).Result()
	res = strings.TrimSpace(res)
	// key不存在时, res == "" && 返回的error不为空
	if res == "" {
		// 缓存为空，没有命中缓存
		return false, nil
	}
	// 命中缓存
	err := json.Unmarshal([]byte(res), data)
	if err != nil {
		return true, fmt.Errorf("[library:cache]:getDataFromCache -> json.Unmarshal failed with error: %s, data: %s", err.Error(), res)
	}
	return true, nil
}
