package repository

import (
	"encoding/json"
	"fmt"
	"projectName/internal/model"
	"strings"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FirstById(id int64) (*model.User, error)
	FirstByName(name string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUserByID(id int64, data map[string]interface{}) (int64, error)

	SetData(key string, data string, exp int64) error
	DelData(key []string) error
	GetData(key string, data interface{}) (bool, error)
}
type userRepository struct {
	*Repository
}

func NewUserRepository(repository *Repository) UserRepository {
	return &userRepository{
		Repository: repository,
	}
}

func (r *userRepository) FirstById(id int64) (*model.User, error) {
	var user model.User
	// TODO: query db

	err := r.Repository.db.Where("id=?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &user, nil
	}
	if err != nil {
		return &user, err
	}
	return &user, nil
}

// FirstByName implements UserRepository.
func (r *userRepository) FirstByName(name string) (*model.User, error) {
	var user model.User

	err := r.Repository.db.Where("username=?", name).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &user, nil
	}
	if err != nil {
		return &user, err
	}
	return &user, nil
}

// CreateUser implements UserRepository.
func (r *userRepository) CreateUser(user *model.User) error {

	t := time.Now().Unix()
	user.UpdatedAt = t
	user.CreatedAt = t
	err := r.Repository.db.Create(user).Error
	return err
}

func (r *userRepository) SetData(token string, data string, exp int64) error {

	return r.rdb.Set(r.rdb.Context(), token, data, time.Duration(exp)*time.Second).Err()
}

// DelData implements UserRepository.
func (r *userRepository) DelData(key []string) error {
	return r.rdb.Del(r.rdb.Context(), key...).Err()

}

// GetData implements UserRepository.
func (r *userRepository) GetData(key string, data interface{}) (bool, error) {
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

// UpdateUserByID implements UserRepository.
func (r *userRepository) UpdateUserByID(id int64, data map[string]interface{}) (int64, error) {
	result := r.Repository.db.Model(&model.User{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}
