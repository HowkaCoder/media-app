package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type MainCategoryUseCase interface {
	GetAllMainCategories(language string) ([]entity.MainCategory, error)
	GetSingleMainCategory(id uint) (*entity.MainCategory, error)
	CreateMainCategory(maincategory *entity.MainCategory) error
	UpdateMainCategory(id uint, maincategory *entity.MainCategory) error
	DeleteMainCategory(id uint) error
}

type mainCategoryUseCase struct {
	MainCategoryRepo repository.MainCategoryRepository
}

func NewMainCategoryUseCase(categoryRepo repository.MainCategoryRepository) *mainCategoryUseCase {
	return &mainCategoryUseCase{MainCategoryRepo: categoryRepo}
}

func (mu *mainCategoryUseCase) GetAllMainCategories(language string) ([]entity.MainCategory, error) {
	return mu.MainCategoryRepo.GetAllMainCategories(language)
}

func (mu *mainCategoryUseCase) GetSingleMainCategory(id uint) (*entity.MainCategory, error) {
	return mu.MainCategoryRepo.GetSingleMainCategory(id)
}
func (mu *mainCategoryUseCase) CreateMainCategory(maincategory *entity.MainCategory) error {
	return mu.MainCategoryRepo.CreateMainCategory(maincategory)
}
func (mu *mainCategoryUseCase) UpdateMainCategory(id uint, maincategory *entity.MainCategory) error {
	return mu.MainCategoryRepo.UpdateMainCategory(id, maincategory)
}
func (mu *mainCategoryUseCase) DeleteMainCategory(id uint) error {
	return mu.MainCategoryRepo.DeleteMainCategory(id)
}
