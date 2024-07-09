package handler

import (
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type languageHandler struct {
	languageUseCase usecase.LanguageUseCase
}

func NewLanguageHandler(languageUsecase usecase.LanguageUseCase) *languageHandler {
	return &languageHandler{
		languageUseCase: languageUsecase,
	}
}
func (h *languageHandler) GetAllLanguages(c *fiber.Ctx) error {
	languages, err := h.languageUseCase.GetAllLanguages()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(languages)
}

func (h *languageHandler) GetLanguageByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	language, err := h.languageUseCase.GetLanguageByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(language)
}

func (h *languageHandler) CreateLanguage(c *fiber.Ctx) error {
	var language entity.Language
	if err := c.BodyParser(&language); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.languageUseCase.CreateLanguage(&language); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(language)
}

func (h *languageHandler) UpdateLanguage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	var language entity.Language
	if err := c.BodyParser(&language); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.languageUseCase.UpdateLanguage(&language, uint(id)); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Language updated successfully"})
}

func (h *languageHandler) DeleteLanguage(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	if err := h.languageUseCase.DeleteLanguage(uint(id)); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Language deleted successfully"})
}
