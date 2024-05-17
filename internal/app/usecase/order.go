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

func (lu *orderUseCase) GetAllOrders() ([]entity.Order, error) {
	return lu.repo.GetAllSubOrders()
}

func (lu *orderUseCase) GetOrderByID(id uint) (*entity.Order, error) {
	return lu.repo.GetOrderById(id)
}

func (lu *orderUseCase) CreateOrder(order *entity.Order) error {
	return lu.repo.CreateOrder(order)
}

func (lu *orderUseCase) UpdateOrder(order *entity.Order, id uint) error {
	return lu.repo.UpdateOrder(order, id)
}

func (lu *orderUseCase) DeleteOrder(id uint) error {
	return lu.repo.DeleteOrder(id)
}
