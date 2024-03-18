package service

import (
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

	err := ps.categoryRepo.GetExsistCategory(product.CategoryID)
	if err != nil {
		return err
	}
	return nil
}
