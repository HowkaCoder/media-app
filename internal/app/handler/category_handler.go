package handler

import (
	"media-app/internal/app/usecase"

	"github.com/gofiber/fiber/v2"
)

type categoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUseCase) *categoryHandler {
	return &categoryHandler{
		categoryUseCase: categoryUsecase,
	}
}

func (h *categoryHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.categoryUseCase.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(categories)
}
