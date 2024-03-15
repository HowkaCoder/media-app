package entity

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	ID                 uint           ` gorm:"primaryKey ; column:id" json:"id"`
	CreatedAt          time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ParentCategoryID   *uint          `json:"parent_category_id"`
	Level              uint           `gorm:"not null" json:"level"`
	NameUZ             string         `gorm:"not null" json:"name_uz"`
	NameKK             string         `gorm:"not null" json:"name_kk"`
	NameRU             string         `gorm:"not null" json:"name_ru"`
	NameEN             string         `gorm:"not null" json:"name_en"`
	ParentCategory     *Category      `gorm:"foreignKey:ParentCategoryID" json:"parent_category"`
	ChildrenCategories []*Category    `gorm:"foreignKey:ParentCategoryID" json:"children_categories"`
}
