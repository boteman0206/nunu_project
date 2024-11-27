package repository

import (
	"context"
	"projectName/internal/model/params"
	"projectName/internal/model/response"
)

type LoginRepository interface {
	Login(ctx context.Context, params params.LoginParams) (*response.LoginResponse, error)
	LoginOut(ctx context.Context, params params.LoginParams) (*response.LoginResponse, error)
}

func NewLoginRepository(
	repository *Repository,
) LoginRepository {
	return &loginRepository{
		Repository: repository,
	}
}

type loginRepository struct {
	*Repository
}

func (r *loginRepository) Login(ctx context.Context, params params.LoginParams) (*response.LoginResponse, error) {
	var login response.LoginResponse

	return &login, nil
}

func (r *loginRepository) LoginOut(ctx context.Context, params params.LoginParams) (*response.LoginResponse, error) {
	var login response.LoginResponse

	return &login, nil
}
