package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
)

type ProductUseCase interface {

	// PRODUCT - CRUD FUNCTIONS
	GetProductsSortedByPriceAndCategory(sortOrder string, categoryID uint) ([]entity.Product, error)
	GetProductsByCharacteristics(value, description string) ([]entity.Product, error)
	GetProductsByPriceRange(minPrice, maxPrice uint) ([]entity.Product, error)
	GetProductsByCategoryID(id uint) ([]entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	GetProductsWithPagination(limit int) ([]entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product entity.Product, id uint) error
	DeleteProduct(id uint) error

	// IMAGE - CRUD FUNCTIONS

	CreateImage(image *entity.Image) error
	GetImagesByProductID(product_id uint) ([]entity.Image, error)
	GetImageByID(id uint) (*entity.Image, error)
	DeleteImage(id uint) error
	UpdateImage(image entity.Image, id uint) error

	// Characteristics - CRUD FUNCTIONS

	CreateCharacteristic(characteristic *entity.Characteristic) error
	GetCharacteristicsByProductID(product_id uint) ([]entity.Characteristic, error)
	DeleteCharacteristic(id uint) error
	UpdateCharacteristic(characteristic entity.Characteristic, id uint) error
}

type productUseCase struct {
	productRepo    repository.ProductRepository
	productService service.ProductService
}

func NewProductUseCase(repo repository.ProductRepository, productService service.ProductService) *productUseCase {
	return &productUseCase{productRepo: repo, productService: productService}
}

// IMAGE FUNCTIONS

func (pu *productUseCase) CreateImage(image *entity.Image) error {
	return pu.productRepo.CreateImage(image)
}

func (pu *productUseCase) GetImagesByProductID(product_id uint) ([]entity.Image, error) {
	return pu.productRepo.GetImagesByProductID(product_id)
}

func (pu *productUseCase) GetImageByID(id uint) (*entity.Image, error) {
	return pu.productRepo.GetImageByID(id)
}

func (pu *productUseCase) DeleteImage(id uint) error {
	return pu.productRepo.DeleteImage(id)
}

func (pu *productUseCase) UpdateImage(image entity.Image, id uint) error {
	return pu.productRepo.UpdateImage(image, id)
}

// CHARACTERISTIC FUNCTIONS

func (pu *productUseCase) CreateCharacteristic(characteristic *entity.Characteristic) error {
	return pu.productRepo.CreateCharacteristic(characteristic)
}

func (pu *productUseCase) GetCharacteristicsByProductID(product_id uint) ([]entity.Characteristic, error) {
	return pu.productRepo.GetCharacteristicsByProductID(product_id)
}

func (pu *productUseCase) UpdateCharacteristic(characteristic entity.Characteristic, id uint) error {
	return pu.productRepo.UpdateCharacteristic(characteristic, id)
}

func (pu *productUseCase) DeleteCharacteristic(id uint) error {
	return pu.productRepo.DeleteCharacteristic(id)
}

// PRODUCT FUNCTIONS

func (pu *productUseCase) CreateProduct(product *entity.Product) error {
	//pu.productService.ValidateProduct(product)
	return pu.productRepo.CreateProduct(product)
}

func (pu *productUseCase) GetProductsByCharacteristics(value, description string) ([]entity.Product, error) {
	return pu.productRepo.GetProductsByCharacteristics(value, description)

}

func (pu *productUseCase) GetAllProducts() ([]entity.Product, error) {

	return pu.productRepo.GetAllProducts()
}

func (pu *productUseCase) UpdateProduct(product entity.Product, id uint) error {

	pu.productService.ValidateProduct(&product)

	return pu.productRepo.UpdateProduct(product, id)
}

func (pu *productUseCase) GetProductByID(id uint) (*entity.Product, error) {
	return pu.productRepo.GetProductByID(id)
}

func (pu *productUseCase) GetProductsWithPagination(limit int) ([]entity.Product, error) {
	return pu.productRepo.GetProductsWithPagination(limit)
}

func (pu *productUseCase) DeleteProduct(id uint) error {
	return pu.productRepo.DeleteProduct(id)
}

func (pu *productUseCase) GetProductsByCategoryID(id uint) ([]entity.Product, error) {
	return pu.productRepo.GetProductsByCategoryID(id)
}

func (pu *productUseCase) GetProductsByPriceRange(minPrice, maxPrice uint) ([]entity.Product, error) {
	return pu.productRepo.GetProductsByPriceRange(minPrice, maxPrice)
}

func (pu *productUseCase) GetProductsSortedByPriceAndCategory(sortOrder string, categoryID uint) ([]entity.Product, error) {
	return pu.productRepo.GetProductsSortedByPriceAndCategory(sortOrder, categoryID)
}
