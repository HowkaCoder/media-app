package service

import (
	"errors"
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type ProductService interface {
	ValidateProduct(product *entity.Product) error
}

type productService struct {
	categoryRepo repository.CategoryRepository
}

func NewProductService(categoryRepository repository.CategoryRepository) *productService {
	return &productService{categoryRepo: categoryRepository}
}

func (ps *productService) ValidateProduct(product *entity.Product) error {
	category, err := ps.categoryRepo.GetSingleCategory(product.CategoryID)
	if err != nil {
		return err

	}
	if category.ID == 0 {
		return errors.New("Error with categoryID")
	}

	return nil
}
