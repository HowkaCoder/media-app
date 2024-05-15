package entity

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ID            uint           ` gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	SubCategoryID uint           `gorm:"not null" json:"sub_category_id"`
	Name          string         `gorm:"not null" json:"name"`
	Description   string         `gorm:"not null" json:"description"`
	Price         uint           `gorm:"not null" json:"price"`
	Discount      uint           `gorm:"null" json:"discount"`
	Quantity      uint           `gorm:"not null" json:"quantity"`
	Language      string         `gorm:"not null" json:"language"`
	Category      SubCategory    `gorm:"foreignKey:SubCategoryID" json:"category"`

	Images          []Image               `json:"images" gorm:"foreignKey:ProductID"`
	Characteristics []Characteristic      `json:"characteristics" gorm:"foreignKey:ProductID"`
	Translations    []ProductTranslations `json:"translations" gorm:"foreignKey:ProductID"`
}

type Characteristic struct {
	gorm.Model
	ID          uint           ` gorm:"primaryKey ; column:id" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ProductID   uint           `gorm:"not null" json:"product_id"`
	Value       string         `gorm:"not null" json:"value"`
	Description string         `gorm:"not null" json:"description"`

	Translations []CharacteristicTranslation `gorm:"foreignKey:CharacteristicID"`
}

type Image struct {
	gorm.Model
	ID        uint           ` gorm:"primaryKey ; column:id" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ProductID uint           `gorm:"not null" json:"product_id"`

	Path string `gorm:"not null" json:"path"`
}
