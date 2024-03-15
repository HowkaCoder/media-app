package repository

import (
	"gorm.io/gorm"
	"media-app/internal/app/entity"
)

type UserRepository interface {
	//   USERS
	GetAllUsers() ([]entity.User, error)
	FindUserByUsername(username string) (*entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User, id uint) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindUserByUsername(username string) (*entity.User, error) {
	var user *entity.User
	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUserByID(id uint) (*entity.User, error) {
	var user *entity.User
	if err := ur.db.Preload("Ava").First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) CreateUser(user *entity.User) error {
	return ur.db.Create(&user).Error
}

func (ur *userRepository) UpdateUser(user *entity.User, id uint) error {
	var eUser *entity.User
	if err := ur.db.First(&eUser, id).Error; err != nil {
		return err
	}
	if user.Username != "" {
		eUser.Username = user.Username
	}
	if user.Firstname != "" {
		eUser.Firstname = user.Firstname
	}
	if user.Lastname != "" {
		eUser.Lastname = user.Lastname
	}
	if user.Age != 0 {
		eUser.Age = user.Age
	}
	if user.Phone != 0 {
		eUser.Phone = user.Phone
	}
	if user.Address != "" {
		eUser.Address = user.Address
	}
	if user.Password != "" {
		eUser.Password = user.Password
	}
	if user.Role != "" {
		eUser.Role = user.Role
	}

	return ur.db.Save(&eUser).Error
}

func (ur *userRepository) DeleteUser(id uint) error {
	var user *entity.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return err
	}
	return ur.db.Delete(&user).Error
}
