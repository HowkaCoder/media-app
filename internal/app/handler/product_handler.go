package handler

import (
	"github.com/gofiber/fiber/v2"
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"os"
	"path/filepath"
	"strconv"
)

type ProductHandler struct {
	productUsecase usecase.ProductUseCase
}

func NewProductHandler(useCase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUsecase: useCase}
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {

	var request struct {
		Product         entity.Product          `json:"product"`
		Images          []entity.Image          `json:"images"`
		Characteristics []entity.Characteristic `json:"characteristics"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ph.productUsecase.CreateProduct(request.Product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	c.Set("Access-Control-Allow-Credentials", "true")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	images := form.File["images[]"]

	for _, imageFile := range images {
		photoPath := filepath.Join("uploads", "photo", imageFile.Filename)

		if err := c.SaveFile(imageFile, photoPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save photo", "error": err.Error()})
		}

		image := entity.Image{
			ProductID: request.Product.ID,
			Path:      photoPath,
		}
		if err := ph.productUsecase.CreateImage(image); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

	}
	for _, chars := range request.Characteristics {
		chars.ProductID = request.Product.ID
		if err := ph.productUsecase.CreateCharacteristic(chars); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}
	}
	return c.JSON(fiber.Map{"message": "Successfully Created"})

}

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	var request struct {
		Product         entity.Product          `json:"product"`
		Images          []entity.Image          `json:"images"`
		Characteristics []entity.Characteristic `json:"characteristics"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ph.productUsecase.UpdateProduct(request.Product, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	images := form.File["images[]"]

	oldImages, err := ph.productUsecase.GetImagesByProductID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	for _, oldImage := range oldImages {
		if err := os.Remove(oldImage.Path); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

		if err := ph.productUsecase.DeleteImage(oldImage.ID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	oldValues, err := ph.productUsecase.GetCharacteristicsByProductID(request.Product.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	for _, oldValue := range oldValues {
		if err := ph.productUsecase.DeleteCharacteristic(oldValue.ProductID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	for _, imageFile := range images {
		photoPath := filepath.Join("uploads", "photo", imageFile.Filename)

		if err := c.SaveFile(imageFile, photoPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save photo", "error": err.Error()})
		}

		image := entity.Image{
			ProductID: request.Product.ID,
			Path:      photoPath,
		}
		if err := ph.productUsecase.CreateImage(image); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

	}

	for _, Value := range request.Characteristics {
		if err := ph.productUsecase.CreateCharacteristic(Value); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}
	}

	return c.JSON(fiber.Map{"message": "successfully updated"})

}
