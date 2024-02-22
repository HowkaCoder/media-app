package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type UsersUseCase interface {
	//   USERS
	GetAllUsers() ([]entity.User, error)
	FindUserByUsername(username string) (*entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User, id uint) error
	DeleteUser(id uint) error

	// AVA
	GetAvaByUserID(id uint) (*entity.Ava, error)
	GetAvaByID(id uint) (*entity.Ava, error)
	CreateAva(ava *entity.Ava) error
	DeleteAva(id uint) error
}

type usersUseCase struct {
	usersRepo repository.UserRepository
}

func NewUsersUseCase(userRepository repository.UserRepository) *usersUseCase {
	return &usersUseCase{usersRepo: userRepository}
}

func (uu *usersUseCase) FindUserByUsername(username string) (*entity.User, error) {
	return uu.usersRepo.FindUserByUsername(username)
}

func (uu *usersUseCase) GetAllUsers() ([]entity.User, error) {
	return uu.usersRepo.GetAllUsers()
}

func (uu *usersUseCase) GetUserByID(id uint) (*entity.User, error) {
	return uu.usersRepo.GetUserByID(id)
}

func (uu *usersUseCase) CreateUser(user *entity.User) error {
	return uu.usersRepo.CreateUser(user)
}

func (uu *usersUseCase) UpdateUser(user *entity.User, id uint) error {
	return uu.usersRepo.UpdateUser(user, id)
}

func (uu *usersUseCase) DeleteUser(id uint) error {
	return uu.usersRepo.DeleteUser(id)
}

func (uu *usersUseCase) GetAvaByUserID(id uint) (*entity.Ava, error) {
	return uu.usersRepo.GetAvaByUserID(id)
}

func (uu *usersUseCase) GetAvaByID(id uint) (*entity.Ava, error) {
	return uu.usersRepo.GetAvaByID(id)
}

func (uu *usersUseCase) CreateAva(ava *entity.Ava) error {
	return uu.usersRepo.CreateAva(ava)
}

func (uu *usersUseCase) DeleteAva(id uint) error {
	return uu.usersRepo.DeleteAva(id)
}