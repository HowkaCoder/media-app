package repository

import (
	"errors"
	"gorm.io/gorm"
	"media-app/internal/app/entity"
)

type MainCategoryRepository interface {
	GetAllMainCategories(language string) ([]entity.MainCategory, error)
	GetSingleMainCategory(id uint) (*entity.MainCategory, error)
	CreateMainCategory(maincategory *entity.MainCategory) error
	UpdateMainCategory(id uint, maincategory *entity.MainCategory) error
	DeleteMainCategory(id uint) error
	//GetCategoriesWithPagination(limit, offset int) ([]entity.MainCategory, error)
}

type maincategoryRepository struct {
	db *gorm.DB
}

func NewMainCategoryRepository(db *gorm.DB) *maincategoryRepository {
	return &maincategoryRepository{
		db: db,
	}
}

func (r *maincategoryRepository) GetAllMainCategories(language string) ([]entity.MainCategory, error) {
	var categories []entity.MainCategory
	if err := r.db.Where("language = ?", language).Preload("SubCategories").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *maincategoryRepository) GetSingleMainCategory(id uint) (*entity.MainCategory, error) {

	var maincategory *entity.MainCategory

	if err := r.db.Preload("SubCategories").First(&maincategory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return maincategory, nil
}

func (r *maincategoryRepository) CreateMainCategory(maincategory *entity.MainCategory) error {

	if err := r.db.Create(&maincategory).Error; err != nil {
		return err
	}
	return nil
}

func (r *maincategoryRepository) UpdateMainCategory(id uint, maincategory *entity.MainCategory) error {

	// поиск обновляемой категории
	var eMainCategory *entity.MainCategory
	if err := r.db.First(&eMainCategory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Record not found")
		}
	}

	// проверка , функция динамичности
	if maincategory.Language != "" {
		eMainCategory.Language = maincategory.Language
	}
	if maincategory.Name != "" {
		eMainCategory.Name = maincategory.Name
	}

	return r.db.Save(&eMainCategory).Error
}

func (r *maincategoryRepository) DeleteMainCategory(id uint) error {

	// поиск удаляемой категории
	var maincategory *entity.MainCategory
	if err := r.db.First(&maincategory, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Record not found")
		}
	}

	return r.db.Delete(&maincategory).Error
}

//func (r *maincategoryRepository) GetCategoriesWithPagination(limit, offset int) ([]entity.MainCategory, error) {
//	var categories []entity.MainCategory
//	if err := r.db.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
//		return nil, err
//	}
//	return categories, nil
//}
