package internal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"media-app/internal/app/entity"
)

var DB *gorm.DB

func Init() *gorm.DB {
	var err error
dsn := "root:AvmOCFLHdwIkOcWYyXzGhuDvuTToYjsM@tcp(viaduct.proxy.rlwy.net:38909)/railway?charset=utf8mb4&parseTime=True&loc=Local"
DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

//	DB, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&entity.Image{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(&entity.OrderProduct{})
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

	err = DB.AutoMigrate(&entity.User{} , &entity.Metric{})
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

	DB.Exec("DELETE FROM sqlite_sequence WHERE name='products'")
	DB.Exec("INSERT INTO sqlite_sequence (name, seq) VALUES ('products', 12345678)")

	DB.Exec("DELETE FROM sqlite_sequence WHERE name='main_categories'")
	DB.Exec("INSERT INTO sqlite_sequence (name, seq) VALUES ('main_categories, 12345678)")

	mn := entity.MainCategory{ID: 12345678, Name: "32432d2d23d23"}
	DB.Create(&mn)
	DB.Delete(&mn)

	DB.Exec("DELETE FROM sqlite_sequence WHERE name='sub_categories'")
	DB.Exec("INSERT INTO sqlite_sequence (name, seq) VALUES ('sub_categories', 12345678)")

	DB.Exec("DELETE FROM sqlite_sequence WHERE name='users'")
	DB.Exec("INSERT INTO sqlite_sequence (name, seq) VALUES ('users', 12345678)")

	return DB
}
