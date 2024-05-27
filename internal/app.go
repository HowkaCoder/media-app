package internal

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"media-app/internal/app/entity"
)

var DB *gorm.DB

func Init() *gorm.DB {
	var err error
	//dsn := "root:pBaYLMKHIVQFHPBbbRKAfphLmzReYKSu@tcp(roundhouse.proxy.rlwy.net:39674)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	//DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.Image{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.Characteristic{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.Product{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.Language{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.ProductTranslations{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.CharacteristicTranslation{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.JWTCredentials{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.SubCategory{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.MainCategory{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.Order{})
	if err != nil {
		log.Fatal(err)
	}

	var tables []string
	DB.Raw("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'").Scan(&tables)

	for _, table := range tables {
		var autoIncrementField string
		DB.Raw(fmt.Sprintf("PRAGMA table_info(%s)", table)).Scan(&autoIncrementField)
  
		if autoIncrementField != "" {
			sql := fmt.Sprintf("UPDATE sqlite_sequence SET seq = 12345677 WHERE name = '%s'", table)
			DB.Exec(sql)
			DB.Exec(fmt.Sprintf("INSERT INTO sqlite_sequence (name, seq) SELECT '%s', 12345677 WHERE NOT EXISTS (SELECT 1 FROM sqlite_sequence WHERE name = '%s')", table, table))
		}
	}
	return DB
}
