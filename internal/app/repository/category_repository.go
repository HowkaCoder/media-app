package repository

import (
	"errors"
	"media-app/internal/app/entity"
	"strconv"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories() ([]entity.Category, error)
	GetSingleCategory(id uint) (*entity.Category, error)
	CreateCategory(category *entity.Category) error
	UpdateCategory(id uint, category *entity.Category) error
	DeleteCategory(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category
	if err := r.db.Preload("ParentCategory").Preload("ChildrenCategories").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) GetSingleCategory(id uint) (*entity.Category, error) {
	var category *entity.Category
	if err := r.db.Preload("ParentCategory").First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) CreateCategory(category *entity.Category) error {
	if category.ParentCategoryID != 0 {

		parentCategory, err := r.GetSingleCategory(category.ParentCategoryID)
		if err != nil {
			return err
		}
		if parentCategory == nil {
			return errors.New("Parent category does not exist")
		}

		category.Level = parentCategory.Level + 1
	}
	if category.ParentCategoryID == 0 {
		category.Level = 1
	}
	if err := r.db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) UpdateCategory(id uint, category *entity.Category) error {

	var eCategory *entity.Category
	if err := r.db.First(&eCategory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errors.New("Record not found")
		}
	}

	if eCategory.ParentCategoryID != 0 {

		parentCategory, err := r.GetSingleCategory(category.ParentCategoryID)
		if err != nil {
			return err
		}
		if parentCategory == nil {
			return errors.New("Parent category does not exist")
		}

		eCategory.Level = parentCategory.Level + 1
	}

	eCategory.ParentCategoryID = category.ParentCategoryID
	eCategory.NameEN = category.NameEN
	eCategory.NameKK = category.NameKK
	eCategory.NameRU = category.NameRU
	eCategory.NameUZ = category.NameUZ

	return r.db.Save(&eCategory).Error
}

func (r *categoryRepository) DeleteCategory(id uint) error {
	var category *entity.Category
	if err := r.db.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errors.New("Record not found")
		}
	}
	return r.db.Delete(&category).Error
}

func uintToString(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

func stringToUint(value string) uint {
	result, _ := strconv.ParseUint(value, 10, 64)
	return uint(result)
}
