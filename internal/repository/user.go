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
