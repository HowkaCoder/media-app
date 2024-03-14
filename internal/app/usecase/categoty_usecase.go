package usecase

import (
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
	GetCategoriesWithPagination(limit, offset int) ([]entity.Category, error)
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
	return us.categoryRepo.CreateCategory(category)
}

func (us *categoryUseCase) UpdateCategory(category *entity.Category, id uint) error {

	return us.categoryRepo.UpdateCategory(id, category)

}

func (us *categoryUseCase) DeleteCategory(id uint) error {
	return us.categoryRepo.DeleteCategory(id)
}

func (us *categoryUseCase) GetCategoriesWithPagination(limit, offset int) ([]entity.Category, error) {
	return us.categoryRepo.GetCategoriesWithPagination(limit, offset)
}
