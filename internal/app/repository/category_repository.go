package repository

import (
	"errors"
	"gorm.io/gorm"
	"media-app/internal/app/entity"
)

type CategoryRepository interface {
	GetAllCategories() ([]entity.Category, error)
	GetSingleCategory(id uint) (*entity.Category, error)
	CreateCategory(category *entity.Category) error
	UpdateCategory(id uint, category *entity.Category) error
	DeleteCategory(id uint) error
	GetExsistCategory(id *uint) error
	GetCategoriesWithPagination(limit, offset int) ([]entity.Category, error)
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

func (r *categoryRepository) GetExsistCategory(id *uint) error {
	var category *entity.Category

	if err := r.db.Where("id = ?", uint(*id)).First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}

	return nil
}

func (r *categoryRepository) GetSingleCategory(id uint) (*entity.Category, error) {

	var category *entity.Category

	if err := r.db.Preload("ParentCategory").Preload("ChildrenCategories").First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) CreateCategory(category *entity.Category) error {

	// Проверка на наличие родительсокй категории
	if category.ParentCategoryID != nil {

		parentCategory := &entity.Category{ID: *category.ParentCategoryID}
		if err := r.db.First(parentCategory).Error; err != nil {
			return errors.New("Parent category does not exist")
		}
		category.Level = parentCategory.Level + 1
	}

	if category.ParentCategoryID == nil {
		category.Level = 1
	}

	if err := r.db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) UpdateCategory(id uint, category *entity.Category) error {

	// поиск обновляемой категории
	var eCategory *entity.Category
	if err := r.db.First(&eCategory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errors.New("Record not found")
		}
	}

	// проверка , функция динамичности
	if category.NameEN != "" {
		eCategory.NameEN = category.NameEN
	}
	if category.NameKK != "" {
		eCategory.NameKK = category.NameKK
	}
	if category.NameRU != "" {
		eCategory.NameRU = category.NameRU
	}
	if category.NameUZ != "" {
		eCategory.NameUZ = category.NameUZ
	}
	if category.ParentCategoryID != nil {
		if *category.ParentCategoryID != eCategory.ID {
			eCategory.ParentCategoryID = category.ParentCategoryID
		} else {
			return errors.New("You just cant change it bro!")
		}
	}

	return r.db.Save(&eCategory).Error
}

func (r *categoryRepository) DeleteCategory(id uint) error {

	// поиск удаляемой категории
	var category *entity.Category
	if err := r.db.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errors.New("Record not found")
		}
	}

	return r.db.Delete(&category).Error
}

func (r *categoryRepository) GetCategoriesWithPagination(limit, offset int) ([]entity.Category, error) {
	var categories []entity.Category
	if err := r.db.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
