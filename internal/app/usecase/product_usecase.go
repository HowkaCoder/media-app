package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type ProductUseCase interface {

	// PRODUCT - CRUD FUNCTIONS
	GetProductsSortedByPriceAndCategory(sortOrder string, categoryID uint, language string) ([]entity.Product, error)
	GetProductsByPriceRange(minPrice, maxPrice uint, language string) ([]entity.Product, error)
	GetProductsByCharacteristics(value, description, language string) ([]entity.Product, error)
	GetProductsByCategoryID(id uint, language string) ([]entity.Product, error)
	GetAllProducts(language string) ([]entity.Product, error)
	GetProductsWithPagination(limit int, language string) ([]entity.Product, error)
	GetProductByID(id uint, language string) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product, id uint) error
	DeleteProduct(id uint) error
	GetProductsSorterByThreeParams(name , price , discount string ) ([]entity.Product, error)
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
	productRepo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) *productUseCase {
	return &productUseCase{productRepo: repo}
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


func (pu *productUseCase) GetProductsSorterByThreeParams(name , price , discount string) ([]entity.Product , error ) {
	return pu.productRepo.GetProductsSortedByThreeParams(name , price , discount)
}


func (pu *productUseCase) CreateProduct(product *entity.Product) error {
	//pu.productService.ValidateProduct(product)
	return pu.productRepo.CreateProduct(product)
}

func (pu *productUseCase) GetProductsByCharacteristics(value, description, language string) ([]entity.Product, error) {
	return pu.productRepo.GetProductsByCharacteristics(value, description, language)

}

func (pu *productUseCase) GetAllProducts(language string) ([]entity.Product, error) {

	return pu.productRepo.GetAllProducts(language)
}

func (pu *productUseCase) UpdateProduct(product *entity.Product, id uint) error {

	//err := pu.productService.ValidateProduct(product)
	//if err != nil {
	//	return err
	//}

	return pu.productRepo.UpdateProduct(product, id)
}

func (pu *productUseCase) GetProductByID(id uint, language string) (*entity.Product, error) {
	return pu.productRepo.GetProductByID(id, language)
}

func (pu *productUseCase) GetProductsWithPagination(limit int, language string) ([]entity.Product, error) {
	return pu.productRepo.GetProductsWithPagination(limit, language)
}

func (pu *productUseCase) DeleteProduct(id uint) error {
	return pu.productRepo.DeleteProduct(id)
}

func (pu *productUseCase) GetProductsByCategoryID(id uint, language string) ([]entity.Product, error) {
	return pu.productRepo.GetProductsByCategoryID(id, language)
}

func (pu *productUseCase) GetProductsByPriceRange(minPrice, maxPrice uint, language string) ([]entity.Product, error) {
	return pu.productRepo.GetProductsByPriceRange(minPrice, maxPrice, language)
}

func (pu *productUseCase) GetProductsSortedByPriceAndCategory(sortOrder string, categoryID uint, language string) ([]entity.Product, error) {
	return pu.productRepo.GetProductsSortedByPriceAndCategory(sortOrder, categoryID, language)
}
