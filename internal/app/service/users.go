package service

import (
	"github.com/golang-jwt/jwt"
	"media-app/internal/app/entity"
	"time"
)

type UserService interface {
	GenerateAccessToken(user *entity.User) (string, error)
	GenerateRefreshToken(user *entity.User) (string, error)
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
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 50).Unix(), // Token expires in 15 minutes
		},
	})

	return token.SignedString(entity.SecretKey)
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
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Token expires in 7 days
		},
	})

	return token.SignedString(entity.SecretKey)
}
