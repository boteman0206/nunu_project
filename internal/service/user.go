package service

import (
	"encoding/json"
	"fmt"

	"projectName/internal/model"
	"projectName/internal/model/params"
	"projectName/internal/model/response"
	"projectName/internal/repository"
	"projectName/pkg/helper/md5"
	"projectName/pkg/helper/uuid"

	"go.uber.org/zap"
)

type UserService interface {
	Register(params *params.RegisterParams) (int, error)
	Login(params *params.LoginParams) (response.LoginResponse, int, error)
	LoginOut(params *params.LoginOutParams) error
	ChangePassword(params *params.ChangeParams) (int, error)

	GetUserById(id int64) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
	GetUserIDByToken(token string) (*model.User, int, error)
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

func (s *userService) GetUserIDByToken(token string) (*model.User, int, error) {

	res := &model.User{}
	exist, err := s.userRepository.GetData(token, res)
	if err != nil {
		s.logger.Error("ChangePassword", zap.Any("err", err))
		return res, 0, err
	}

	if !exist || res.ID <= 0 {
		return res, model.TokenExpErr, nil
	}
	return res, 0, nil
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
		return model.CodeUserNameALreadyExist, nil
	}
	// todo 加密密码处理
	// 账号的特殊字符和密码的特殊字符校验 长度校验等

	data := &model.User{
		Username:    params.Username,
		Password:    params.Password,
		PortraitUrl: params.PortraitUrl,
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

	uuid := uuid.GenUUID()
	token := fmt.Sprintf("%s_%d", uuid, user.ID)
	token, err = md5.HashPassword(token)
	if err != nil {
		s.logger.Error("HashPassword", zap.Any("err", err))
		return res, model.CodeNetError, err
	}

	res.Token = token

	userBytes, err := json.Marshal(user)
	if err != nil {
		s.logger.Error("json Marshal", zap.Any("err", err))
		return res, 0, err
	}

	err = s.userRepository.SetData(token, string(userBytes), model.TokenExp)
	if err != nil {
		s.logger.Error("SetLoginToken", zap.Any("err", err))
		return res, 0, err
	}

	return res, 0, nil
}

// LoginOut implements UserService.
func (s *userService) LoginOut(params *params.LoginOutParams) error {

	err := s.userRepository.DelData([]string{params.Token})
	if err != nil {
		s.logger.Error("LoginOut", zap.Any("err", err))
		return err
	}
	return nil
}

// UpdateUser implements UserService.
func (s *userService) ChangePassword(params *params.ChangeParams) (int, error) {

	res := &model.User{}
	exist, err := s.userRepository.GetData(params.Token, res)
	if err != nil {
		s.logger.Error("ChangePassword", zap.Any("err", err))
		return 0, err
	}

	if !exist || res.ID <= 0 {
		return model.TokenExpErr, nil
	}

	if res.Password != params.OldPassword ||
		params.Username != res.Username {
		return model.CodeOldPasswordErr, nil
	}

	data := map[string]interface{}{
		"password": params.NewPassword,
	}

	row, err := s.userRepository.UpdateUserByID(res.ID, data)
	if err != nil {
		s.logger.Error("ChangePassword", zap.Any("err", err))
		return 0, err
	}
	if row != 1 {
		return model.CodeNetError, nil
	}
	return 0, nil
}
