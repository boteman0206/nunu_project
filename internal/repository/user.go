package repository

import (
	"projectName/internal/model"

	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FirstById(id int64) (*model.User, error)
	FirstByName(name string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUserByID(id int64, data map[string]interface{}) (int64, error)
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

// UpdateUserByID implements UserRepository.
func (r *userRepository) UpdateUserByID(id int64, data map[string]interface{}) (int64, error) {

	data["updated_at"] = time.Now().Unix()
	result := r.Repository.db.Model(&model.User{}).Where("id = ?", id).Updates(data)
	return result.RowsAffected, result.Error
}
