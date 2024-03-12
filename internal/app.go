package internal

import (
	"gorm.io/gorm"
	"log"
	"media-app/internal/app/entity"

	"gorm.io/driver/postgres"
)

func Init() *gorm.DB {
	dsn := "host=monorail.proxy.rlwy.net user=postgres password=xlsCRByWFYqwvqBIrVHkkAvFBWkaqHLJ dbname=railway port=31045"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.JWTCredentials{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
