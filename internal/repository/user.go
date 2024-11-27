package repository

import (
	"projectName/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FirstById(id int64) (*model.User, error)
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
