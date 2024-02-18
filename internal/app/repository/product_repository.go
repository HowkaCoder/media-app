package repository

import (
	"errors"
	"gorm.io/gorm"
	"media-app/internal/app/entity"
)

type ProductRepository interface {

	// PRODUCT - CRUD FUNCTIONS
	GetProductsSortedByPriceAndCategory(sortOrder string, categoryID uint) ([]entity.Product, error)
	GetProductsByPriceRange(minPrice, maxPrice uint) ([]entity.Product, error)
	GetProductsByCharacteristics(value, description string) ([]entity.Product, error)
	GetProductsByCategoryID(id uint) ([]entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	GetProductsWithPagination(limit int) ([]entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product entity.Product, id uint) error
	DeleteProduct(id uint) error

	// IMAGE - CRUD FUNCTIONS

	CreateImage(image *entity.Image) error
	GetImageByID(id uint) (*entity.Image, error)
	GetImagesByProductID(product_id uint) ([]entity.Image, error)
	DeleteImage(id uint) error

	// Characteristics - CRUD FUNCTIONS

	CreateCharacteristic(characteristic *entity.Characteristic) error
	GetCharacteristicsByProductID(product_id uint) ([]entity.Characteristic, error)
	DeleteCharacteristic(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

// IMAGE FUNCTIONS

func (pr *productRepository) CreateImage(image *entity.Image) error {
	return pr.db.Create(&image).Error
}

func (pr *productRepository) GetImagesByProductID(product_id uint) ([]entity.Image, error) {
	var images []entity.Image
	if err := pr.db.Where("product_id = ?", product_id).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (pr *productRepository) GetImageByID(id uint) (*entity.Image, error) {
	var image *entity.Image
	if err := pr.db.First(&image, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return image, nil
}

func (pr *productRepository) DeleteImage(id uint) error {
	var eImage entity.Image
	if err := pr.db.First(&eImage, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return pr.db.Delete(&eImage).Error
}

// CHARACTERISTICS FUNCTIONS

func (pr *productRepository) CreateCharacteristic(characteristic *entity.Characteristic) error {
	return pr.db.Create(&characteristic).Error
}

func (pr *productRepository) GetCharacteristicsByProductID(product_id uint) ([]entity.Characteristic, error) {
	var characteristics []entity.Characteristic
	if err := pr.db.Where("product_id = ?", product_id).Find(&characteristics).Error; err != nil {
		return nil, err
	}
	return characteristics, nil
}

func (pr *productRepository) DeleteCharacteristic(id uint) error {
	var characteristic entity.Characteristic
	if err := pr.db.First(&characteristic, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return pr.db.Delete(&characteristic).Error

}

// PRODUCT FUNCTIONS

func (pr *productRepository) CreateProduct(product *entity.Product) error {
	return pr.db.Create(product).Error
}

func (pr *productRepository) UpdateProduct(product entity.Product, id uint) error {
	var eProduct entity.Product
	if err := pr.db.First(&eProduct, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errors.New("Record not found")
		}
	}
	if product.CategoryID != 0 {
		if product.CategoryID != eProduct.CategoryID {
			eProduct.CategoryID = product.CategoryID
		}
	}
	if product.Discount != 0 {
		eProduct.Discount = product.Discount
	}
	if product.Name != "" {
		eProduct.Name = product.Name
	}
	if product.Price != 0 {
		eProduct.Price = product.Price
	}
	if product.Quantity != 0 {
		eProduct.Quantity = product.Quantity
	}
	return pr.db.Save(&eProduct).Error
}

func (pr *productRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.Preload("Images").Preload("Characteristics").Preload("Category").Preload("Translations").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productRepository) GetProductByID(id uint) (*entity.Product, error) {
	var product *entity.Product
	if err := pr.db.Preload("Category").Preload("Images").Preload("Characteristics").Preload("Translations").First(&product, id).Error; err != nil {
		return nil, err
	}

	return product, nil

}

func (pr *productRepository) GetProductsWithPagination(limit int) ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.Preload("Images").Preload("Characteristics").Preload("Category").Preload("Translations").Limit(limit).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productRepository) DeleteProduct(id uint) error {
	var product *entity.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return pr.db.Delete(&product).Error
}

func (pr *productRepository) GetProductsByCategoryID(id uint) ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.Where("category_id = ?", id).Preload("Translations").Find(&products); err != nil {

	}
	return products, nil
}

func (pr *productRepository) GetProductsByPriceRange(minPrice, maxPrice uint) ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.Where("price >= ? AND price <= ?", minPrice, maxPrice).Preload("Translations").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productRepository) GetProductsSortedByPriceAndCategory(sortOrder string, categoryID uint) ([]entity.Product, error) {
	var products []entity.Product
	order := "price ASC"
	if sortOrder == "expensive" {
		order = "price DESC"
	}
	if err := pr.db.Where("category_id = ?", categoryID).Order(order).Preload("Translations").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productRepository) GetProductsByCharacteristics(value, description string) ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.
		Joins("JOIN characteristics ON products.id = characteristics.product_id").
		Where("value = ? AND description = ?", value, description).
		Preload("Translations").
		Find(&products).
		Error; err != nil {
		return nil, err
	}
	return products, nil
}
