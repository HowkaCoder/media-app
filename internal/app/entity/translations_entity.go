package entity

import "gorm.io/gorm"

type ProductTranslations struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	ProductID  uint   `gorm:"not null"`
	LanguageID uint   `gorm:"not null"`
	Name       string `gorm:"not null"`
}

type CharacteristicTranslation struct {
	gorm.Model
	ID               uint   `gorm:"primaryKey"`
	CharacteristicID uint   `gorm:"not null"`
	LanguageID       uint   `gorm:"not null"`
	Value            string `gorm:"not null"`
	Description      string `gorm:"not null"`
}
