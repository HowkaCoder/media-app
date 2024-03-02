package entity

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var (
	SecretKey = []byte("3278yd&8327dh32*(@#$E(2")
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null, unique"`
	Firstname string `gorm:"not null"`
	Lastname  string `gorm:"not null"`
	Age       uint   `gorm:"not null"`
	Phone     uint   `gorm:"not null , unique"`
	Address   string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	Ava       string `gorm:"not null"`
}

type JWTCredentials struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       uint   `json:"age"`
	Address   string `json:"address"`
	jwt.StandardClaims
}
