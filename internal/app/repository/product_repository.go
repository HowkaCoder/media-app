package repository

import (
	"errors"
	"media-app/internal/app/entity"
	"strconv"
	"sync"

	"gorm.io/gorm"
)

type ProductRepository interface {

	// PRODUCT - CRUD FUNCTIONS
	GetProductsSortedByThreeParams(name , price , discount string) ([]entity.Product , error)
	GetProductsSortedByPriceAndCategory(sortOrder string, categoryID uint, language string) ([]entity.Product, error)
	GetProductsByPriceRange(minPrice, maxPrice uint, language string) ([]entity.Product, error)
	GetProductsByCharacteristics(value, description, language string) ([]entity.Product, error)
	GetProductsByCategoryID(id uint, language string) ([]entity.Product, error)
	GetAllProducts(language string) ([]entity.Product, error)
	GetProductsWithPagination(limit int,  offset int) ([]entity.Product, error)
	GetProductsByFilter(discount []string, minPrice string , maxPrice string , subcategoryID []string , brand string) ([]entity.Product , error)
	GetProductByID(id uint, language string) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product, id uint) error
	DeleteProduct(id uint) error
	

	// IMAGE - CRUD FUNCTIONS

	CreateImage(image *entity.Image) error
	GetImageByID(id uint) (*entity.Image, error)
	GetImagesByProductID(product_id uint) ([]entity.Image, error)
	DeleteImage(id uint) error
	UpdateImage(image entity.Image, id uint) error
	// Characteristics - CRUD FUNCTIONS

	CreateCharacteristic(characteristic *entity.Characteristic) error
	GetCharacteristicsByProductID(product_id uint) ([]entity.Characteristic, error)
	DeleteCharacteristic(id uint) error
	UpdateCharacteristic(characteristic entity.Characteristic, id uint) error
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

func (pr *productRepository) UpdateImage(image entity.Image, id uint) error {
	var Eimage entity.Image
	if err := pr.db.First(&Eimage, id).Error; err != nil {
		return err
	}

	if image.Path != "" {
		Eimage.Path = image.Path
	}

	return pr.db.Save(&Eimage).Error
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

func (pr *productRepository) UpdateCharacteristic(characteristic entity.Characteristic, id uint) error {
	var Echar entity.Characteristic
	if err := pr.db.First(&Echar, id).Error; err != nil {
		return err
	}

	if characteristic.Value != "" {
		Echar.Value = characteristic.Value
	}
	if characteristic.Description != "" {
		Echar.Description = characteristic.Description
	}

	return pr.db.Save(&Echar).Error
}

// PRODUCT FUNCTIONS


func (pr *productRepository) GetProductsSortedByThreeParams(name , price , discount string) ([]entity.Product , error) {
	var products []entity.Product

	query := pr.db.Model(&entity.Product{})

	if name != ""  {
		if name == "asc" {
			query = query.Order("name asc")
		} else if name == "desc" {
			query = query.Order("name desc")
		}
	}
		
	if price != "" {
		if price == "asc" {
			query = query.Order("price asc")
		} else if price == "desc" {
			query = query.Order("price desc")
		}
	}

	if discount != "" {
		if discount == "asc" {
			query = query.Order("discount asc")
		} else if discount == "desc" {
			query = query.Order("discount desc")
		}
	}

	if err := query.Find(&products).Error; err != nil {
		return nil , err
	}

	return products,nil
}














func (pr *productRepository) CreateProduct(product *entity.Product) error {
	return pr.db.Create(product).Error
}

func (pr *productRepository) UpdateProduct(product *entity.Product, id uint) error {
	var eProduct *entity.Product
	if err := pr.db.First(&eProduct, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return	errors.New("Product not found")
		}
	}

	if product.SubCategoryID != 0 {
		if product.SubCategoryID != eProduct.SubCategoryID {
			eProduct.SubCategoryID = product.SubCategoryID
		}
	}
	if product.Discount != 0 {
		eProduct.Discount = product.Discount
	}
	if product.Name != "" {
		eProduct.Name = product.Name
	}
	if product.Description != "" {
		eProduct.Description = product.Description
	}
	if product.Price != 0 {
		eProduct.Price = product.Price
	}
	if product.Quantity != 0 {
		eProduct.Quantity = product.Quantity
	}
	if product.Language != "" {
		eProduct.Language = product.Language
	}
	return pr.db.Save(&eProduct).Error
}

func (pr *productRepository) GetAllProducts(language string) ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.Where("language = ?", language).Preload("Images").Preload("Characteristics").Preload("Category").Preload("Translations").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productRepository) GetProductByID(id uint, language string) (*entity.Product, error) {
	var product *entity.Product
	if err := pr.db.Where("language = ?", language).Preload("Category").Preload("Images").Preload("Characteristics").Preload("Translations").First(&product, id).Error; err != nil {
		return nil, err
	}

	return product, nil

}

func (pr *productRepository) GetProductsWithPagination(limit int , offset int) ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.Offset(offset).Preload("Images").Preload("Characteristics").Preload("Category").Preload("Translations").Limit(limit).Find(&products).Error; err != nil {
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

func (pr *productRepository) GetProductsByCategoryID(id uint, language string) ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.Where("sub_category_id = ? AND language = ?", id, language).Preload("Translations").Preload("Category").Preload("Images").Preload("Characteristics").Find(&products); err != nil {

	}
	return products, nil
}
/*

func (pr *productRepository) GetProductsByFilter(discount []string , minPrice string , maxPrice string , subcategoryID []string) ([]entity.Product , error) {

	var products []entity.Product

	query := pr.db.Model(&entity.Product{})

	// Проверка на наличие значений для фильтрации
	if len(subcategoryID) != 0 {
    for _ , value := range subcategoryID {
      id , _ := strconv.Atoi(value)
		  query = query.Where("sub_category_id = ?", uint(id))
    }  
  }
	if minPrice != "" {
    id , _ := strconv.Atoi(minPrice)
		query = query.Where("price >= ?", uint(id))
	}
	if maxPrice != "" {
    id , _ := strconv.Atoi(maxPrice)
		query = query.Where("price <= ?", uint(id))
	}

	if len(discount) != 0 {
    for _ , value := range discount {
      id , _ := strconv.Atoi(value)
		  query = query.Where("discount = ?" , uint(id))
	  }
  }
	


	// Проверка на наличие значений для сортировки
	

	if err := query.Preload("Category").Preload("Images").Preload("Characteristics").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

*/







func (pr *productRepository) GetProductsByFilter(discount []string, minPrice string, maxPrice string, subcategoryID []string , brand string) ([]entity.Product, error) {
	var products []entity.Product
	query := pr.db.Model(&entity.Product{})

	var wg sync.WaitGroup
	var mu sync.Mutex
	errors := make(chan error, 1)

	// Обработка фильтрации по subcategoryID
	wg.Add(1)
	go func() {
		defer wg.Done()
		if len(subcategoryID) != 0 {
			mu.Lock()         // Блокируем мьютекс перед изменением запроса
			defer mu.Unlock() // Разблокируем мьютекс после завершения функции
			for _, value := range subcategoryID {
				id, err := strconv.Atoi(value)
				if err != nil {
					errors <- err
					return
				}
				query = query.Where("sub_category_id = ?", uint(id))
			}
		}
	}()

	// Обработка фильтрации по minPrice
	wg.Add(1)
	go func() {
		defer wg.Done()
		if minPrice != "" {
			id, err := strconv.Atoi(minPrice)
			if err != nil {
				errors <- err
				return
			}
			mu.Lock()
			query = query.Where("price >= ?", uint(id))
			mu.Unlock()
		}
	}()




	// Обработка по brand

	wg.Add(1)
	go func() {
		defer wg.Done()
		if brand != "" {
			mu.Lock()
			query = query.Where("brand = ?" , brand)
			mu.Unlock()
		}
		
	}()





	// Обработка фильтрации по maxPrice
	wg.Add(1)
	go func() {
		defer wg.Done()
		if maxPrice != "" {
			id, err := strconv.Atoi(maxPrice)
			if err != nil {
				errors <- err
				return
			}
			mu.Lock()
			query = query.Where("price <= ?", uint(id))
			mu.Unlock()
		}
	}()

	// Обработка фильтрации по discount
	wg.Add(1)
	go func() {
		defer wg.Done()
		if len(discount) != 0 {
			mu.Lock()
			defer mu.Unlock()
			for _, value := range discount {
				id, err := strconv.Atoi(value)
				if err != nil {
					errors <- err
					return
				}
				query = query.Where("discount = ?", uint(id))
			}
		}
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
	close(errors)

	if err := <-errors; err != nil {
		return nil, err
	}

	// Выполняем запрос с предзагрузками
	if err := query.Preload("Category").Preload("Images").Preload("Characteristics").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}























func (pr *productRepository) GetProductsByPriceRange(minPrice, maxPrice uint, language string) ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.Where("price >= ? AND price <= ? AND language = ?", minPrice, maxPrice, language).Preload("Translations").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productRepository) GetProductsSortedByPriceAndCategory(sortOrder string, categoryID uint, language string) ([]entity.Product, error) {
	var products []entity.Product
	order := "price ASC"
	if sortOrder == "expensive" {
		order = "price DESC"
	}
	if err := pr.db.Where("category_id = ? AND language = ?", categoryID, language).Order(order).Preload("Translations").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *productRepository) GetProductsByCharacteristics(value, description, language string) ([]entity.Product, error) {
	var products []entity.Product
	if err := pr.db.
		Joins("JOIN characteristics ON products.id = characteristics.product_id").
		Where("value = ? AND description = ? AND language=?", value, description, language).
		Preload("Translations").
		Preload("Category").
		Preload("Images").
		Preload("Characteristics").
		Find(&products).
		Error; err != nil {
		return nil, err
	}
	return products, nil
}
