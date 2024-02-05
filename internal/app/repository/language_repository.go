package repository

import (
	"errors"
	"gorm.io/gorm"
	"media-app/internal/app/entity"
)

type LanguageRepository interface {
	GetAllLanguages() ([]entity.Language, error)
	GetLanguageByID(id uint) (*entity.Language, error)
	CreateLanguage(language *entity.Language) error
	UpdateLanguage(language *entity.Language, id uint) error
	DeleteLanguage(id uint) error
}

type languageRepository struct {
	db *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *languageRepository {
	return &languageRepository{db: db}
}

func (lr *languageRepository) GetAllLanguages() ([]entity.Language, error) {
	var languages []entity.Language
	if err := lr.db.Find(&languages).Error; err != nil {
		return nil, err
	}
	return languages, nil
}

func (lr *languageRepository) GetLanguageByID(id uint) (*entity.Language, error) {
	var language *entity.Language
	if err := lr.db.First(&language, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Record not found")
		} else {
			return nil, err
		}
	}
	return language, nil
}

func (lr *languageRepository) CreateLanguage(language *entity.Language) error {
	return lr.db.Create(&language).Error
}

func (lr *languageRepository) UpdateLanguage(language *entity.Language, id uint) error {
	var eLanguage *entity.Language
	if err := lr.db.First(&language, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Record not found")
		} else {
			return err
		}
	}

	if language.Name != "" {
		eLanguage.Name = language.Name
	}
	return lr.db.Save(&eLanguage).Error
}

func (lr *languageRepository) DeleteLanguage(id uint) error {
	var language *entity.Language
	if err := lr.db.First(&language, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Record not found")
		} else {
			return err
		}
	}
	return lr.db.Delete(&language).Error
}
