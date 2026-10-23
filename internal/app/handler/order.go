package handler

import (
	"errors"
	"media-app/internal/app/entity"
	"media-app/internal/app/repository"
	"media-app/internal/app/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase usecase.OrderUseCase) *orderHandler {
	return &orderHandler{
		orderUseCase: orderUseCase,
	}
}

func (h *orderHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.orderUseCase.GetAllOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(orders)
}

func (h *orderHandler) GetOrderByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	order, err := h.orderUseCase.GetOrderByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(order)
}

func (h *orderHandler) CreateOrder(c *fiber.Ctx) error {
	var order entity.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if userID, ok := c.Locals("user_id").(uint); ok {
		order.UserID = userID
	}
	err := h.orderUseCase.CreateOrder(&order)
	if errors.Is(err, repository.ErrOutOfStock) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *orderHandler) UpdateOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var order entity.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.orderUseCase.UpdateOrder(&order, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Order updated successfully"})
}

func (h *orderHandler) DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.orderUseCase.DeleteOrder(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Order deleted successfully"})
}
