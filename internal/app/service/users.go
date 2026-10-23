package service

import (
	"errors"
	"fmt"
	"media-app/internal/app/entity"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessTTL  = 10 * time.Hour
	refreshTTL = 20 * 24 * time.Hour
)

type UserService interface {
	GenerateAccessToken(user *entity.User) (string, error)
	GenerateRefreshToken(user *entity.User) (string, error)
	ParseToken(token string) (*entity.JWTCredentials, error)
	HashPassword(password string) (string, error)
	CheckPassword(hash, password string) bool
}

type userService struct {
	secret []byte
}

func NewUserService(secret []byte) UserService {
	return &userService{secret: secret}
}

func (s *userService) GenerateAccessToken(user *entity.User) (string, error) {
	return s.sign(user, accessTTL)
}

func (s *userService) GenerateRefreshToken(user *entity.User) (string, error) {
	return s.sign(user, refreshTTL)
}

func (s *userService) sign(user *entity.User, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.JWTCredentials{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(ttl).Unix(),
		},
	})
	return token.SignedString(s.secret)
}

func (s *userService) ParseToken(tokenString string) (*entity.JWTCredentials, error) {
	token, err := jwt.ParseWithClaims(tokenString, &entity.JWTCredentials{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entity.JWTCredentials)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func (s *userService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *userService) CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
