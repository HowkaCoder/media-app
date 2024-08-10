package entity

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"time"
)

var (
	SecretKey = []byte("3278yd&8327dh32*(@#$E(2")
)

type User struct {
	gorm.Model
	ID        uint           ` gorm:"primaryKey ; column:id;autoIncrement;autoIncrement:12345678" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Username  string         `gorm:"type:varchar(100);uniqueIndex" json:"username"`
	Firstname string         `gorm:"not null" json:"firstname"`
	Lastname  string         `gorm:"not null" json:"lastname"`
	Age       uint           `gorm:"not null" json:"age"`
	Phone     uint           `gorm:"uniqueIndex" json:"phone"`
	Address   string         `gorm:"not null" json:"address"`
	Password  string         `gorm:"not null" json:"password"`
	Role      string         `gorm:"not null" json:"role"`
	Ava       string         `gorm:"null" json:"ava"`
}

type JWTCredentials struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       uint   `json:"age"`
	Phone 	  uint   `json:"phone"`
	Address   string `json:"address"`
	Password  string `json:"password"`
	Ava       string `json:"ava"`
	jwt.StandardClaims
}
