package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type CategoryUseCase interface {
	GetAllCategories() ([]entity.Category, error)
	GetCategoryByID(id uint) (*entity.Category, error)
	CreateCategory(category *entity.Category) error
	UpdateCategory(category *entity.Category, id uint) error
	DeleteCategory(id uint) error
}

type categoryUseCase struct {
	categoryRepo repository.CategoryRepository
	// there soon willbe services
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepository) *categoryUseCase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (us *categoryUseCase) GetAllCategories() ([]entity.Category, error) {
	return us.categoryRepo.GetAllCategories()
}

func (us *categoryUseCase) GetCategoryByID(id uint) (*entity.Category, error) {
	return us.categoryRepo.GetSingleCategory(id)
}

func (us *categoryUseCase) CreateCategory(category *entity.Category) error {
	return us.categoryRepo.CreateCategory(category)
}

func (us *categoryUseCase) UpdateCategory(category *entity.Category, id uint) error {
	return us.categoryRepo.UpdateCategory(id, category)
}

func (us *categoryUseCase) DeleteCategory(id uint) error {
	return us.categoryRepo.DeleteCategory(id)
}
