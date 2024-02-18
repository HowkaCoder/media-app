package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type LanguageUseCase interface {
	GetAllLanguages() ([]entity.Language, error)
	GetLanguageByID(id uint) (*entity.Language, error)
	CreateLanguage(language *entity.Language) error
	UpdateLanguage(language *entity.Language, id uint) error
	DeleteLanguage(id uint) error
}

type languageUseCase struct {
	repo repository.LanguageRepository
}

func NewLanguageUseCase(languageRepository repository.LanguageRepository) *languageUseCase {
	return &languageUseCase{repo: languageRepository}
}

func (lu *languageUseCase) GetAllLanguages() ([]entity.Language, error) {
	return lu.repo.GetAllLanguages()
}

func (lu *languageUseCase) GetLanguageByID(id uint) (*entity.Language, error) {
	return lu.repo.GetLanguageByID(id)
}

func (lu *languageUseCase) CreateLanguage(language *entity.Language) error {
	return lu.repo.CreateLanguage(language)
}

func (lu *languageUseCase) UpdateLanguage(language *entity.Language, id uint) error {
	return lu.repo.UpdateLanguage(language, id)
}

func (lu *languageUseCase) DeleteLanguage(id uint) error {
	return lu.repo.DeleteLanguage(id)
}
