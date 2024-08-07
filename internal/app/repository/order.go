package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"media-app/internal/app/entity"
)

type OrderRepository interface {
	GetAllOrders() ([]entity.Order, error)
	GetOrderByID(id uint) (*entity.Order, error)
	CreateOrder(order *entity.Order) error
	UpdateOrder(order *entity.Order, id uint) error
	DeleteOrder(id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (or *orderRepository) GetAllOrders() ([]entity.Order, error) {
	var orders []entity.Order
	if err := or.db.Preload("Products.Image").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (or *orderRepository) GetOrderByID(id uint) (*entity.Order, error) {
	var order entity.Order
	if err := or.db.Preload("Products.Image").First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Record not found")
		}
		return nil, err
	}
	return &order, nil
}

func (or *orderRepository) CreateOrder(order *entity.Order) error {
    for i, value := range order.Products {
        productID := value.ProductID
        var product entity.Product
        if err := or.db.First(&product, productID).Error; err != nil {
            return err
        }



        if product.Quantity != 0 {
          product.Quantity = product.Quantity - 1
        }

        // Обновляем поля order.Products[i]
        order.Products[i].Description  = product.Description   
        order.Products[i].Title = product.Name
        order.Products[i].Price = product.Price
        order.Products[i].Discount = product.Discount
        
        if err := or.db.Save(&product).Error; err != nil {
          return err
        }


	fmt.Println(order.Products[i])
    }

	fmt.Println(order)
    // Сохраняем заказ вместе с обновленными продуктами
    return or.db.Create(&order).Error
}

func (or *orderRepository) UpdateOrder(order *entity.Order, id uint) error {
	var existingOrder entity.Order
	if err := or.db.First(&existingOrder, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Record not found")
		}
		return err
	}
	return or.db.Model(&existingOrder).Updates(order).Error
}

func (or *orderRepository) DeleteOrder(id uint) error {
	var order entity.Order
	if err := or.db.First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Record not found")
		}
		return err
	}
	return or.db.Delete(&order).Error
}
