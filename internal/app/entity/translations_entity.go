package entity

import (
	"gorm.io/gorm"
	"time"
)

type ProductTranslations struct {
	gorm.Model
	ID          uint           ` gorm:"primaryKey ; column:id;autoIncrement;autoIncrement:12345678" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ProductID   uint           `gorm:"not null" json:"product_id"`
	LanguageID  uint           `gorm:"not null" json:"language_id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
}

type CharacteristicTranslation struct {
	gorm.Model
	ID               uint           ` gorm:"primaryKey ; column:id" json:"id"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CharacteristicID uint           `gorm:"not null" json:"characteristic_id"`
	LanguageID       uint           `gorm:"not null" json:"language_id"`
	Value            string         `gorm:"not null" json:"value"`
	Description      string         `gorm:"not null" json:"description"`
}
