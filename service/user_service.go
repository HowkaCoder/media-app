package service

import (
	"media-app/model"
	"media-app/repository"
)

type UserService interface {
	GetUsers() ([]model.User, error)
	GetUser(id uint) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(id uint, user *model.User) (*model.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) GetUsers() ([]model.User, error) {
	return s.userRepository.GetUsers()
}

func (s *userService) GetUser(id uint) (*model.User, error) {
	return s.userRepository.GetUser(id)
}

func (s *userService) CreateUser(user *model.User) (*model.User, error) {
	return s.userRepository.CreateUser(user)
}

func (s *userService) UpdateUser(id uint, user *model.User) (*model.User, error) {
	return s.userRepository.UpdateUser(id, user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepository.DeleteUser(id)
}
