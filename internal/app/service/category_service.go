// internal/mycrudapp/service/category_service.go
package service

import (
	"errors"
	"media-app/internal/app/entity"
)

type CategoryService interface {
	ValidateCategory(category *entity.Category) error
	ValidateCategoryByUpdate(category *entity.Category, id uint) error
}

type categoryService struct{}

func NewCategoryService() CategoryService {
	return &categoryService{}
}

func (s *categoryService) ValidateCategory(category *entity.Category) error {
	if category.NameUZ == "" {
		return errors.New("NameUZ cannot be empty")
	}
	if category.NameKK == "" {
		return errors.New("NameKK cannot be empty")
	}
	if category.NameRU == "" {
		return errors.New("NameRU cannot be empty")
	}
	if category.NameEN == "" {
		return errors.New("NameEN cannot be empty")
	}

	return nil
}

func (s *categoryService) ValidateCategoryByUpdate(category *entity.Category, id uint) error {
	if id == *category.ParentCategoryID {
		return errors.New("Parent Category iD cannot be as a category ID !!!")
	}

	return nil
}
