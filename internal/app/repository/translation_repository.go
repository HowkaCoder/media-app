package repository

import (
	"gorm.io/gorm"
	"media-app/internal/app/entity"
)

type TranslationRepository interface {
	CreateProductTranslation(translation *entity.ProductTranslations) error
	GetProductTranslationsByProductID(productID uint) ([]entity.ProductTranslations, error)
	UpdateProductTranslation(translation *entity.ProductTranslations, id uint) error
	DeleteProductTranslation(id uint) error

	CreateCharacteristicTranslation(translation *entity.CharacteristicTranslation) error
	GetCharacteristicTranslationsByCharacteristicID(characteristicID uint) ([]entity.CharacteristicTranslation, error)
	UpdateCharacteristicTranslation(translation *entity.CharacteristicTranslation, id uint) error
	DeleteCharacteristicTranslation(id uint) error
}

type translationRepository struct {
	db *gorm.DB
}

func NewTranslationRepository(db *gorm.DB) TranslationRepository {
	return &translationRepository{db: db}
}

func (tr *translationRepository) CreateProductTranslation(translation *entity.ProductTranslations) error {
	return tr.db.Create(translation).Error
}

func (tr *translationRepository) GetProductTranslationsByProductID(productID uint) ([]entity.ProductTranslations, error) {
	var translations []entity.ProductTranslations
	if err := tr.db.Where("product_id = ?", productID).Find(&translations).Error; err != nil {
		return nil, err
	}
	return translations, nil
}

func (tr *translationRepository) UpdateProductTranslation(translation *entity.ProductTranslations, id uint) error {
	var ePT *entity.ProductTranslations
	if err := tr.db.First(&ePT, id).Error; err != nil {
		return err
	}
	ePT.Name = translation.Name
	if translation.ProductID != 0 {
		if translation.ProductID != ePT.ProductID {
			ePT.ProductID = translation.ProductID
		}
	}
	return tr.db.Save(&ePT).Error
}

func (tr *translationRepository) DeleteProductTranslation(id uint) error {
	return tr.db.Delete(&entity.ProductTranslations{}, id).Error
}

func (tr *translationRepository) CreateCharacteristicTranslation(translation *entity.CharacteristicTranslation) error {
	return tr.db.Create(translation).Error
}

func (tr *translationRepository) GetCharacteristicTranslationsByCharacteristicID(characteristicID uint) ([]entity.CharacteristicTranslation, error) {
	var translations []entity.CharacteristicTranslation
	if err := tr.db.Where("characteristic_id = ?", characteristicID).Find(&translations).Error; err != nil {
		return nil, err
	}
	return translations, nil
}

func (tr *translationRepository) UpdateCharacteristicTranslation(translation *entity.CharacteristicTranslation, id uint) error {
	var Etranslation *entity.CharacteristicTranslation
	if err := tr.db.First(&Etranslation, id).Error; err != nil {
		return err
	}
	if translation.CharacteristicID != 0 {
		if translation.CharacteristicID != Etranslation.CharacteristicID {
			Etranslation.CharacteristicID = translation.CharacteristicID
		}
	}
	if translation.Value != "" {
		Etranslation.Value = translation.Value
	}
	if translation.Description != "" {
		Etranslation.Description = translation.Description
	}
	return tr.db.Save(&Etranslation).Error
}

func (tr *translationRepository) DeleteCharacteristicTranslation(id uint) error {
	return tr.db.Delete(&entity.CharacteristicTranslation{}, id).Error
}
