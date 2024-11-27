package service

import (
	"projectName/internal/model"
	"projectName/internal/model/params"
	"projectName/internal/model/response"
	"projectName/internal/repository"

	"go.uber.org/zap"
)

type UserService interface {
	GetUserById(id int64) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
	Register(params *params.RegisterParams) (int, error)
	Login(params *params.LoginParams) (response.LoginResponse, int, error)
}

type userService struct {
	*Service
	userRepository repository.UserRepository
}

func NewUserService(service *Service, userRepository repository.UserRepository) UserService {
	return &userService{
		Service:        service,
		userRepository: userRepository,
	}
}

func (s *userService) GetUserById(id int64) (*model.User, error) {
	return s.userRepository.FirstById(id)
}

// GetUserByName implements UserService.
func (s *userService) GetUserByName(name string) (*model.User, error) {
	return s.userRepository.FirstByName(name)
}

// CreateUser implements UserService.
func (s *userService) Register(params *params.RegisterParams) (int, error) {

	user, err := s.GetUserByName(params.Username)
	s.logger.Info("GetUserByName", zap.Any("user", user))
	if err != nil {
		s.logger.Error("GetUserByName", zap.Any("err", err))
		return 0, err
	}
	// 判断用户名是否已经存在
	if user.ID > 0 {
		return model.CodeUserNameNotExist, nil
	}
	// todo 加密密码处理
	// 账号的特殊字符和密码的特殊字符校验 长度校验等

	data := &model.User{
		Username: params.Username,
		Password: params.Password,
	}
	err = s.userRepository.CreateUser(data)
	if err != nil {
		s.logger.Error("CreateUser", zap.Any("err", err))
		return 0, err
	}

	return 0, nil

}

// Login implements UserService.
func (s *userService) Login(params *params.LoginParams) (response.LoginResponse, int, error) {

	res := response.LoginResponse{}
	user, err := s.GetUserByName(params.Username)
	s.logger.Info("GetUserByName", zap.Any("user", user))
	if err != nil {
		s.logger.Error("GetUserByName", zap.Any("err", err))
		return res, 0, err
	}
	// 判断用户名是否已经存在
	if user.ID <= 0 {
		return res, model.CodeUserNameNotExist, nil
	}
	if user.Password != params.Password {
		return res, model.CodePasswordErr, nil
	}

	res.Username = user.Username
	res.ID = user.ID
	return res, 0, nil
}
