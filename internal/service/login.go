package service

import (
	"context"
	"projectName/internal/model/params"
	"projectName/internal/model/response"
	"projectName/internal/repository"
)

type LoginService interface {
	Login(ctx context.Context, prams params.LoginParams) (*response.LoginResponse, error)
	LoginOut(ctx context.Context, prams params.LoginParams) (*response.LoginResponse, error)
}

func NewLoginService(
	service *Service,
	loginRepository repository.LoginRepository,
) LoginService {
	return &loginService{
		Service:         service,
		loginRepository: loginRepository,
	}
}

type loginService struct {
	*Service
	loginRepository repository.LoginRepository
}

func (s *loginService) Login(ctx context.Context, prams params.LoginParams) (*response.LoginResponse, error) {
	return s.loginRepository.Login(ctx, prams)
}

func (s *loginService) LoginOut(ctx context.Context, prams params.LoginParams) (*response.LoginResponse, error) {
	return s.loginRepository.Login(ctx, prams)
}
