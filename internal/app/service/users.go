package service

import (
	errors "errors"
	"github.com/golang-jwt/jwt"
	"media-app/internal/app/entity"
	"time"
)

type UserService interface {
	GenerateAccessToken(user *entity.User) (string, error)
	GenerateRefreshToken(user *entity.User) (string, error)
	ValidateCreateUser(user *entity.User) error
}

type userService struct{}

func NewUserService() UserService { return &userService{} }

func (s *userService) GenerateAccessToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.JWTCredentials{
		UserID:    user.ID,
		Username:  user.Firstname,
		Role:      user.Role,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Age:       user.Age,
		Address:   user.Address,
		Password:  user.Password,
		Ava:       user.Ava,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(), // Token expires in 10 Hours
		},
	})

	return token.SignedString(entity.SecretKey)
}

func (s *userService) ValidateCreateUser(user *entity.User) error {
	if user.Role != "admin" && user.Role != "user" {
		return errors.New("role not allowed")
	}
	return nil
}

func (s *userService) GenerateRefreshToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.JWTCredentials{
		UserID:    user.ID,
		Username:  user.Firstname,
		Role:      user.Role,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Age:       user.Age,
		Address:   user.Address,
		Password:  user.Password,
		Ava:       user.Ava,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 20).Unix(), // Token expires in 7 days
		},
	})

	return token.SignedString(entity.SecretKey)
}
