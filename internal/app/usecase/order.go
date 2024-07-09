package usecase

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
)

type OrderUseCase interface {
	GetAllOrders() ([]entity.Order, error)
	GetOrderByID(id uint) (*entity.Order, error)
	CreateOrder(order *entity.Order) error
	UpdateOrder(order *entity.Order, id uint) error
	DeleteOrder(id uint) error
}

type orderUseCase struct {
	repo repository.OrderRepository
}

func NewOrderUseCase(orderRepository repository.OrderRepository) *orderUseCase {
	return &orderUseCase{repo: orderRepository}
}

func (ou *orderUseCase) GetAllOrders() ([]entity.Order, error) {
	return ou.repo.GetAllOrders()
}

func (ou *orderUseCase) GetOrderByID(id uint) (*entity.Order, error) {
	return ou.repo.GetOrderByID(id)
}

func (ou *orderUseCase) CreateOrder(order *entity.Order) error {
	return ou.repo.CreateOrder(order)
}

func (ou *orderUseCase) UpdateOrder(order *entity.Order, id uint) error {
	return ou.repo.UpdateOrder(order, id)
}

func (ou *orderUseCase) DeleteOrder(id uint) error {
	return ou.repo.DeleteOrder(id)
}
