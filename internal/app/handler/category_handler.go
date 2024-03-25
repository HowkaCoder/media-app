package handler

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"strconv"

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
	limit, _ := strconv.Atoi(c.Query("limit"))

	offset, _ := strconv.Atoi(c.Query("offset"))

	var categories []entity.Category
	var err error
	if limit > 0 || offset > 0 {
		categories, err = h.categoryUseCase.GetCategoriesWithPagination(limit, offset)
	} else {
		categories, err = h.categoryUseCase.GetAllCategories()
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(categories)
}

func (h *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	var category entity.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.categoryUseCase.CreateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "category successfully created"})
}

func (h *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err := h.categoryUseCase.DeleteCategory(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Successfully deleted"})

}

func (h *categoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error 1": err.Error()})
	}

	var category entity.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 2": err.Error()})
	}

	if err := h.categoryUseCase.UpdateCategory(&category, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 3": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "updated successfully"})
}

func (h *categoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	category, err := h.categoryUseCase.GetCategoryByID(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(category)
}
