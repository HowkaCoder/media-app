package repository

import (
	"gorm.io/gorm"
	"media-app/internal/app/entity"
)

type OrderRepository interface {
	GetAllSubOrders() ([]entity.Order, error)
	GetOrderById(id uint) (*entity.Order, error)
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

func (r *orderRepository) GetAllSubOrders() ([]entity.Order, error) {
	var order []entity.Order
	if err := r.db.Preload("Product").Preload("User").Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) GetOrderById(id uint) (*entity.Order, error) {
	var order *entity.Order
	if err := r.db.Preload("Product").Preload("User").First(&order, id).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) CreateOrder(order *entity.Order) error {
	if err := r.db.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) UpdateOrder(order *entity.Order, id uint) error {
	var eOrder *entity.Order
	if err := r.db.First(&eOrder, id).Error; err != nil {
		return err
	}

	if order.Status != "" {
		eOrder.Status = order.Status
	}
	if order.ProductID != 0 {
		eOrder.ProductID = order.ProductID
	}
	if order.UserID != 0 {
		eOrder.UserID = order.UserID
	}

	if err := r.db.Save(eOrder).Error; err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) DeleteOrder(id uint) error {
	var eOrder *entity.Order
	if err := r.db.First(&eOrder, id).Error; err != nil {
		return err
	}
	if err := r.db.Delete(eOrder).Error; err != nil {
		return err
	}
	return nil
}
