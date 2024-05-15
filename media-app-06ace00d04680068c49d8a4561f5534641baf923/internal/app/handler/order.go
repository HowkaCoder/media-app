package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"strconv"
)

type OrderHandler struct {
	orderUseCase usecase.OrderUseCase
	userUsecase  usecase.UsersUseCase
}

func NewOrderHandler(useCase usecase.OrderUseCase, userUseCase usecase.UsersUseCase) *OrderHandler {
	return &OrderHandler{userUsecase: userUseCase, orderUseCase: useCase}
}

func (o *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var request struct {
		ProductID uint        `json:"product_id"`
		User      entity.User `json:"user"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error #1": err.Error()})
	}

	log.Println("____________Create User____________")
	if err := o.userUsecase.CreateUser(&request.User); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #2": err.Error()})
	}
	log.Println("____________Created User____________")
	log.Println(&request.User)

	log.Println("____________Create Order____________")
	order := entity.Order{ProductID: request.ProductID, UserID: request.User.ID, Status: "active"}
	if err := o.orderUseCase.CreateOrder(&order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #3": err.Error()})
	}
	log.Println("____________Created Order____________")
	log.Println(&order)

	return c.JSON(fiber.Map{"message": "order successfully created!"})

}

func (o *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := o.orderUseCase.GetAllOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #1": err.Error()})
	}

	return c.JSON(orders)
}

func (o *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	var order entity.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error #1": err.Error()})
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error #2": err.Error()})
	}

	if err := o.orderUseCase.UpdateOrder(&order, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #3": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "order successfully updated!"})
}

func (o *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error #1": err.Error()})
	}

	if err := o.orderUseCase.DeleteOrder(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #2": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "order successfully deleted!"})
}

func (o *OrderHandler) GetOrderById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error #1": err.Error()})
	}

	order, err := o.orderUseCase.GetOrderByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #2": err.Error()})
	}
	return c.JSON(order)
}
