package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type ProductUseCase interface {

	// PRODUCT - CRUD FUNCTIONS

	GetAllProducts() ([]entity.Product, error)
	//GetProductsWithPagination(limit, offset uint) ([]entity.Product, error)
	//GetProductByID(id uint) (*entity.Product, error)
	CreateProduct(product entity.Product) error
	UpdateProduct(product entity.Product, id uint) error
	//DeleteProduct(id uint) error

	// IMAGE - CRUD FUNCTIONS

	CreateImage(image entity.Image) error
	GetImagesByProductID(product_id uint) ([]entity.Image, error)
	GetImageByID(id uint) (*entity.Image, error)
	DeleteImage(id uint) error

	// Characteristics - CRUD FUNCTIONS

	CreateCharacteristic(characteristic entity.Characteristic) error
	GetCharacteristicsByProductID(product_id uint) ([]entity.Characteristic, error)
	DeleteCharacteristic(id uint) error
}

type productUseCase struct {
	productRepo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) *productUseCase {
	return &productUseCase{productRepo: repo}
}

// IMAGE FUNCTIONS

func (pu *productUseCase) CreateImage(image entity.Image) error {
	return pu.productRepo.CreateImage(&image)
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

// CHARACTERISTIC FUNCTIONS

func (pu *productUseCase) CreateCharacteristic(characteristic entity.Characteristic) error {
	return pu.productRepo.CreateCharacteristic(characteristic)
}

func (pu *productUseCase) GetCharacteristicsByProductID(product_id uint) ([]entity.Characteristic, error) {
	return pu.productRepo.GetCharacteristicsByProductID(product_id)
}

func (pu *productUseCase) DeleteCharacteristic(id uint) error {
	return pu.productRepo.DeleteCharacteristic(id)
}

// PRODUCT FUNCTIONS

func (pu *productUseCase) CreateProduct(product entity.Product) error {
	return pu.productRepo.CreateProduct(product)
}

func (pu *productUseCase) GetAllProducts() ([]entity.Product, error) {
	return pu.productRepo.GetAllProducts()
}

func (pu *productUseCase) UpdateProduct(product entity.Product, id uint) error {
	return pu.productRepo.UpdateProduct(product, id)
}
