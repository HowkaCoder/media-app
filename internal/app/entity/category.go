package entity

import (
	"gorm.io/gorm"
	"time"
)

type MainCategory struct {
	gorm.Model
	ID            uint          ` gorm:"primaryKey ; column:id;autoIncrement;autoIncrement:12345678" json:"id"`
	CreatedAt     time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time     `gorm:"column:updated_at" json:"updated_at"`
	Name          string        `gorm:"not null" json:"name"`
	Language      string        `gorm:"not null" json:"language"`
	SubCategories []SubCategory `gorm:"foreignKey:MainCategoryID" json:"sub_categories"`
}

type SubCategory struct {
	gorm.Model
	ID             uint      ` gorm:"primaryKey ; column:id;autoIncrement;autoIncrement:12345678" json:"id"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	MainCategoryID uint      `gorm:"not null" json:"category_id"`
	Value          string    `gorm:"not null" json:"value"`
	Description    string    `gorm:"not null" json:"description"`
}
