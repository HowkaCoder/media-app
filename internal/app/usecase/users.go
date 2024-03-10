package usecase

import (
	"errors"
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
)

type UsersUseCase interface {
	//   USERS
	GetAllUsers() ([]entity.User, error)
	FindUserByUsername(username string) (*entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User, id uint) error
	DeleteUser(id uint) error
}

type usersUseCase struct {
	usersRepo   repository.UserRepository
	userService service.UserService
}

func NewUsersUseCase(userRepository repository.UserRepository) *usersUseCase {
	return &usersUseCase{usersRepo: userRepository}
}

func (uu *usersUseCase) FindUserByUsername(username string) (*entity.User, error) {
	if username == "" {
		return nil, errors.New("Empty User!!")
	}
	return uu.usersRepo.FindUserByUsername(username)
}

func (uu *usersUseCase) GetAllUsers() ([]entity.User, error) {
	return uu.usersRepo.GetAllUsers()
}

func (uu *usersUseCase) GetUserByID(id uint) (*entity.User, error) {
	return uu.usersRepo.GetUserByID(id)
}

func (uu *usersUseCase) CreateUser(user *entity.User) error {

	if err := uu.userService.ValidateCreateUser(user); err != nil {
		return err
	}

	return uu.usersRepo.CreateUser(user)
}

func (uu *usersUseCase) UpdateUser(user *entity.User, id uint) error {
	return uu.usersRepo.UpdateUser(user, id)
}

func (uu *usersUseCase) DeleteUser(id uint) error {
	return uu.usersRepo.DeleteUser(id)
}
