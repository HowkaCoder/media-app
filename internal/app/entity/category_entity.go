package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID               uint `gorm:"primaryKey"`
	ParentCategoryID *uint
	Level            uint   `gorm:"not null"`
	NameUZ           string `gorm:"not null"`
	NameKK           string `gorm:"not null"`
	NameRU           string `gorm:"not null"`
	NameEN           string `gorm:"not null"`

	ParentCategory     *Category   `gorm:"foreignKey:ParentCategoryID"`
	ChildrenCategories []*Category `gorm:"foreignKey:ParentCategoryID"`
}
