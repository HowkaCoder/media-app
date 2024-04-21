package entity

import (
	"gorm.io/gorm"
	"time"
)

type Language struct {
	gorm.Model
	ID         uint           ` gorm:"primaryKey ; column:id" json:"id"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Name       string         `gorm:"not null" json:"name"`
	MainStatus bool           `gorm:"not null"   json:"main_status" default:"false"`
}
