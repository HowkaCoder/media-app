package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID              uint                  `gorm:"primaryKey"`
	CategoryID      uint                  `gorm:"not null"`
	Name            string                `gorm:"not null"`
	Price           uint                  `gorm:"not null"`
	Discount        uint                  `gorm:"null"`
	Quantity        uint                  `gorm:"not null"`
	Category        Category              `gorm:"foreignKey:CategoryID"`
	Images          []Image               `gorm:"foreignKey:ProductID"`
	Characteristics []Characteristic      `gorm:"foreignKey:ProductID"`
	Translations    []ProductTranslations `gorm:"foreignKey:ProductID"`
}

type Characteristic struct {
	gorm.Model
	ID           uint                        `gorm:"primaryKey"`
	ProductID    uint                        `gorm:"not null"`
	Value        string                      `gorm:"not null"`
	Description  string                      `gorm:"not null"`
	Translations []CharacteristicTranslation `gorm:"foreignKey:CharacteristicID"`
}

type Image struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	ProductID uint   `gorm:"not null"`
	Path      string `gorm:"not null"`
}
