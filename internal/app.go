package internal

import (
	"log"
	"media-app/internal/app/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Category{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Image{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Characteristic{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.Language{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.ProductTranslations{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.CharacteristicTranslation{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
