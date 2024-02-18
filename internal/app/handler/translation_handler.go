package handler

import (
	"github.com/gofiber/fiber/v2"
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"strconv"
)

type TranslationHandler struct {
	translationUsecase usecase.TranslationUseCase
}

func NewTranslationHandler(useCase usecase.TranslationUseCase) *TranslationHandler {
	return &TranslationHandler{translationUsecase: useCase}
}

func (th *TranslationHandler) GetProductTranslationsByProductID(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("productID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	translations, err := th.translationUsecase.GetProductTranslationsByProductID(uint(productID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(translations)
}
func (th *TranslationHandler) CreateProductTranslation(c *fiber.Ctx) error {
	var request struct {
		ProductID  uint   `json:"product_id"`
		LanguageID uint   `json:"language_id"`
		Name       string `json:"name"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	translation := &entity.ProductTranslations{
		ProductID:  request.ProductID,
		LanguageID: request.LanguageID,
		Name:       request.Name,
	}

	if err := th.translationUsecase.CreateProductTranslation(translation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully Created"})
}

func (th *TranslationHandler) UpdateProductTranslation(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	var request struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	translation := &entity.ProductTranslations{
		Name: request.Name,
	}

	if err := th.translationUsecase.UpdateProductTranslation(translation, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully Updated"})
}

func (th *TranslationHandler) DeleteProductTranslation(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	if err := th.translationUsecase.DeleteProductTranslation(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully Deleted"})
}

func (th *TranslationHandler) GetCharacteristicTranslationsByCharacteristicID(c *fiber.Ctx) error {
	characteristicID, err := strconv.Atoi(c.Params("characteristicID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	translations, err := th.translationUsecase.GetCharacteristicTranslationsByCharacteristicID(uint(characteristicID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(translations)
}
func (th *TranslationHandler) CreateCharacteristicTranslation(c *fiber.Ctx) error {
	var request struct {
		CharacteristicID uint   `json:"characteristic_id"`
		LanguageID       uint   `json:"language_id"`
		Value            string `json:"value"`
		Description      string `json:"description"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	translation := &entity.CharacteristicTranslation{
		CharacteristicID: request.CharacteristicID,
		LanguageID:       request.LanguageID,
		Value:            request.Value,
		Description:      request.Description,
	}

	if err := th.translationUsecase.CreateCharacteristicTranslation(translation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully Created"})
}

func (th *TranslationHandler) UpdateCharacteristicTranslation(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	var request struct {
		Value       string `json:"value"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	translation := &entity.CharacteristicTranslation{
		Value:       request.Value,
		Description: request.Description,
	}

	if err := th.translationUsecase.UpdateCharacteristicTranslation(translation, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully Updated"})
}

func (th *TranslationHandler) DeleteCharacteristicTranslation(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	if err := th.translationUsecase.DeleteCharacteristicTranslation(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Successfully Deleted"})
}
