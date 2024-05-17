package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type SubCategoryUseCase interface {
	GetAllSubCategories() ([]entity.SubCategory, error)
	GetSubCategoryById(id uint) (*entity.SubCategory, error)
	CreateSubCategory(subCategory *entity.SubCategory) error
	UpdateSubCategory(subCategory *entity.SubCategory, id uint) error
	DeleteSubCategory(id uint) error
}

type subCategoryUseCase struct {
	repo repository.SubCategoryRepository
}

func NewSubCategoryUseCase(repo repository.SubCategoryRepository) *subCategoryUseCase {
	return &subCategoryUseCase{repo: repo}
}

func (uc *subCategoryUseCase) GetAllSubCategories() ([]entity.SubCategory, error) {
	return uc.repo.GetAllSubCategories()
}
func (uc *subCategoryUseCase) GetSubCategoryById(id uint) (*entity.SubCategory, error) {
	return uc.repo.GetSubCategoryById(id)
}
func (uc *subCategoryUseCase) CreateSubCategory(subCategory *entity.SubCategory) error {
	return uc.repo.CreateSubCategory(subCategory)
}
func (uc *subCategoryUseCase) DeleteSubCategory(id uint) error {
	return uc.repo.DeleteSubCategory(id)
}
func (uc *subCategoryUseCase) UpdateSubCategory(subCategory *entity.SubCategory, id uint) error {
	return uc.repo.UpdateSubCategory(subCategory, id)
}
