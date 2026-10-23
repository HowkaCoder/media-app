package internal

import (
	"fmt"
	"media-app/internal/app/entity"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const idOffset = 12345678

func Init(path string) (*gorm.DB, error) {
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&entity.Image{},
		&entity.OrderProduct{},
		&entity.Characteristic{},
		&entity.Product{},
		&entity.Language{},
		&entity.ProductTranslations{},
		&entity.CharacteristicTranslation{},
		&entity.User{},
		&entity.Metric{},
		&entity.SubCategory{},
		&entity.MainCategory{},
		&entity.Order{},
	)
	if err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	for _, table := range []string{"products", "main_categories", "sub_categories", "users"} {
		if err := startIDsFrom(db, table, idOffset); err != nil {
			return nil, fmt.Errorf("set sequence for %s: %w", table, err)
		}
	}

	return db, nil
}

func startIDsFrom(db *gorm.DB, table string, start int) error {
	if !db.Migrator().HasTable("sqlite_sequence") {
		return nil
	}

	var count int64
	if err := db.Table("sqlite_sequence").Where("name = ?", table).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	return db.Exec("INSERT INTO sqlite_sequence (name, seq) VALUES (?, ?)", table, start).Error
}
