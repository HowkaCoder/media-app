package usecase

import (
	"errors"
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
)

type CategoryUseCase interface {
	GetAllCategories() ([]entity.Category, error)
	GetCategoryByID(id uint) (*entity.Category, error)
	CreateCategory(category *entity.Category) error
	UpdateCategory(category *entity.Category, id uint) error
	DeleteCategory(id uint) error
}

type categoryUseCase struct {
	categoryRepo    repository.CategoryRepository
	categoryService service.CategoryService
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepository, categoryService service.CategoryService) *categoryUseCase {
	return &categoryUseCase{
		categoryRepo:    categoryRepo,
		categoryService: categoryService,
	}
}

func (us *categoryUseCase) GetAllCategories() ([]entity.Category, error) {
	return us.categoryRepo.GetAllCategories()
}

func (us *categoryUseCase) GetCategoryByID(id uint) (*entity.Category, error) {
	return us.categoryRepo.GetSingleCategory(id)
}

func (us *categoryUseCase) CreateCategory(category *entity.Category) error {
	if err := us.categoryService.ValidateCategory(category); err != nil {
		return err
	}
	if category.ParentCategoryID != 0 {

		parentCategory, err := us.categoryRepo.GetSingleCategory(category.ParentCategoryID)
		if err != nil {
			return err
		}
		if parentCategory == nil {
			return errors.New("Parent category does not exist")
		}
	}
	return us.categoryRepo.CreateCategory(category)
}

func (us *categoryUseCase) UpdateCategory(category *entity.Category, id uint) error {
	if err := us.categoryService.ValidateCategoryByUpdate(category, id); err != nil {
		return err
	}
	return us.categoryRepo.UpdateCategory(id, category)
}

func (us *categoryUseCase) DeleteCategory(id uint) error {
	return us.categoryRepo.DeleteCategory(id)
}
