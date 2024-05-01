package repository

import (
	"gorm.io/gorm"
	"media-app/internal/app/entity"
)

type SubCategoryRepository interface {
	GetAllSubCategories() ([]entity.SubCategory, error)
	GetSubCategoryById(id uint) (*entity.SubCategory, error)
	CreateSubCategory(subCategory *entity.SubCategory) error
	UpdateSubCategory(subCategory *entity.SubCategory, id uint) error
	DeleteSubCategory(id uint) error
}

type subCategoryRepository struct {
	db *gorm.DB
}

func NewSubCategoryRepository(db *gorm.DB) *subCategoryRepository {
	return &subCategoryRepository{db: db}
}

func (r *subCategoryRepository) GetAllSubCategories() ([]entity.SubCategory, error) {
	var subCategory []entity.SubCategory
	if err := r.db.Find(&subCategory).Error; err != nil {
		return nil, err
	}
	return subCategory, nil
}

func (r *subCategoryRepository) GetSubCategoryById(id uint) (*entity.SubCategory, error) {
	var subCategory *entity.SubCategory
	if err := r.db.First(&subCategory, id).Error; err != nil {
		return nil, err
	}
	return subCategory, nil
}

func (r *subCategoryRepository) CreateSubCategory(subCategory *entity.SubCategory) error {
	if err := r.db.Create(subCategory).Error; err != nil {
		return err
	}
	return nil
}

func (r *subCategoryRepository) UpdateSubCategory(subCategory *entity.SubCategory, id uint) error {
	var eSubCategory *entity.SubCategory
	if err := r.db.First(&eSubCategory, id).Error; err != nil {
		return err
	}
	if subCategory.MainCategoryID != 0 {
		eSubCategory.MainCategoryID = subCategory.MainCategoryID
	}
	if subCategory.Value != "" {
		eSubCategory.Value = subCategory.Value
	}
	if subCategory.Description != "" {
		eSubCategory.Description = subCategory.Description
	}

	if err := r.db.Save(eSubCategory).Error; err != nil {
		return err
	}
	return nil
}

func (r *subCategoryRepository) DeleteSubCategory(id uint) error {
	var eSubCategory *entity.SubCategory
	if err := r.db.First(&eSubCategory, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(eSubCategory).Error; err != nil {
		return err
	}
	return nil
}
