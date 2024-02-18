package entity

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
