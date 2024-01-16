package repository

import "media-app/model"

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetUser(id uint) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(id uint, user *model.User) (*model.User, error)
	DeleteUser(id uint) error
}

type userRepository struct {
	db model.Database
}

func NewUserRepository(db model.Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepository) GetUser(id uint) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *userRepository) CreateUser(user *model.User) (*model.User, error) {
	result := r.db.Create(user)
	return user, result.Error
}

func (r *userRepository) UpdateUser(id uint, user *model.User) (*model.User, error) {
	result := r.db.Model(&model.User{}).Where("id = ?", id).Updates(user)
	return user, result.Error
}

func (r *userRepository) DeleteUser(id uint) error {
	result := r.db.Delete(&model.User{}, id)
	return result.Error
}
