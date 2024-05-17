package entity

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID        uint      ` gorm:"primaryKey ; column:id;autoIncrement;autoIncrement:12345678" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	ProductID uint      `gorm:"column:product_id" json:"product_id"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	Status    string    `gorm:"column:status" json:"status"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	User      User      `gorm:"foreignKey:UserID" json:"users"`
}
