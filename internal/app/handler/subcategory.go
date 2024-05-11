package handler

import (
	"github.com/gofiber/fiber/v2"
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"strconv"
)

type SubCategoryHandler struct {
	subCategoryUsecase usecase.SubCategoryUseCase
}

func NewSubCategoryHandler(subCategoryUsecase usecase.SubCategoryUseCase) *SubCategoryHandler {
	return &SubCategoryHandler{
		subCategoryUsecase: subCategoryUsecase,
	}
}

func (h *SubCategoryHandler) GetAllSubCategories(c *fiber.Ctx) error {
	subcategories, err := h.subCategoryUsecase.GetAllSubCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(subcategories)
}

func (h *SubCategoryHandler) GetSubCategoryByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	subcategory, err := h.subCategoryUsecase.GetSubCategoryById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(subcategory)
}

func (h *SubCategoryHandler) CreateSubCategory(c *fiber.Ctx) error {
	var subcategory entity.SubCategory
	if err := c.BodyParser(&subcategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	if err := h.subCategoryUsecase.CreateSubCategory(&subcategory); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "SubCategory created successfully"})
}

func (h *SubCategoryHandler) UpdateSubCategory(c *fiber.Ctx) error {
	var subcategory entity.SubCategory
	if err := c.BodyParser(&subcategory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	if err := h.subCategoryUsecase.UpdateSubCategory(&subcategory, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "SubCategory updated successfully"})
}

func (h *SubCategoryHandler) DeleteSubCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := h.subCategoryUsecase.DeleteSubCategory(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.JSON(fiber.Map{"Message": "subcategory deleted successfully"})
}
