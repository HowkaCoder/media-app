package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type TranslationUseCase interface {
	CreateProductTranslation(translation *entity.ProductTranslations) error
	GetProductTranslationsByProductID(productID uint) ([]entity.ProductTranslations, error)
	UpdateProductTranslation(translation *entity.ProductTranslations, id uint) error
	DeleteProductTranslation(id uint) error

	CreateCharacteristicTranslation(translation *entity.CharacteristicTranslation) error
	GetCharacteristicTranslationsByCharacteristicID(characteristicID uint) ([]entity.CharacteristicTranslation, error)
	UpdateCharacteristicTranslation(translation *entity.CharacteristicTranslation, id uint) error
	DeleteCharacteristicTranslation(id uint) error
}

type translationUseCase struct {
	translationRepo repository.TranslationRepository
}

func NewTranslationUseCase(repo repository.TranslationRepository) TranslationUseCase {
	return &translationUseCase{translationRepo: repo}
}

func (tu *translationUseCase) CreateProductTranslation(translation *entity.ProductTranslations) error {
	return tu.translationRepo.CreateProductTranslation(translation)
}

func (tu *translationUseCase) GetProductTranslationsByProductID(productID uint) ([]entity.ProductTranslations, error) {
	return tu.translationRepo.GetProductTranslationsByProductID(productID)
}

func (tu *translationUseCase) UpdateProductTranslation(translation *entity.ProductTranslations, id uint) error {
	return tu.translationRepo.UpdateProductTranslation(translation, id)
}

func (tu *translationUseCase) DeleteProductTranslation(id uint) error {
	return tu.translationRepo.DeleteProductTranslation(id)
}

func (tu *translationUseCase) CreateCharacteristicTranslation(translation *entity.CharacteristicTranslation) error {
	return tu.translationRepo.CreateCharacteristicTranslation(translation)
}

func (tu *translationUseCase) GetCharacteristicTranslationsByCharacteristicID(characteristicID uint) ([]entity.CharacteristicTranslation, error) {
	return tu.translationRepo.GetCharacteristicTranslationsByCharacteristicID(characteristicID)
}

func (tu *translationUseCase) UpdateCharacteristicTranslation(translation *entity.CharacteristicTranslation, id uint) error {
	return tu.translationRepo.UpdateCharacteristicTranslation(translation, id)
}

func (tu *translationUseCase) DeleteCharacteristicTranslation(id uint) error {
	return tu.translationRepo.DeleteCharacteristicTranslation(id)
}
