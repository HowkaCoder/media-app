package handler

import (
	"errors"
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductHandler struct {
	productUsecase usecase.ProductUseCase
	imagesDir      string
	publicURL      string
}

type productRequest struct {
	Product         entity.Product          `json:"product"`
	Images          []entity.Image          `json:"images"`
	Characteristics []entity.Characteristic `json:"characteristics"`
}

func NewProductHandler(useCase usecase.ProductUseCase, imagesDir, publicURL string) *ProductHandler {
	return &ProductHandler{
		productUsecase: useCase,
		imagesDir:      imagesDir,
		publicURL:      strings.TrimSuffix(publicURL, "/"),
	}
}

func (ph *ProductHandler) GetProductsByName(c *fiber.Ctx) error {
	products, err := ph.productUsecase.GetProductsByName(c.Params("name"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var request productRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ph.productUsecase.CreateProduct(&request.Product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	for _, img := range request.Images {
		if err := ph.addImage(request.Product.ID, img.Path); err != nil {
			return ph.imageError(c, err)
		}
	}

	if err := ph.createCharacteristics(request.Product.ID, request.Characteristics); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "product successfully created", "id": request.Product.ID})
}

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	var request productRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	productID := uint(id)
	request.Product.ID = productID

	if err := ph.productUsecase.UpdateProduct(&request.Product, productID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	oldImages, err := ph.productUsecase.GetImagesByProductID(productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	keep := make(map[uint]entity.Image)
	for _, img := range request.Images {
		if img.ID != 0 {
			keep[img.ID] = img
		}
	}

	for _, old := range oldImages {
		img, ok := keep[old.ID]
		if ok {
			if err := ph.productUsecase.UpdateImage(img, img.ID); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
			}
			continue
		}
		if err := ph.deleteImage(old); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	for _, img := range request.Images {
		if img.ID != 0 {
			continue
		}
		if err := ph.addImage(productID, img.Path); err != nil {
			return ph.imageError(c, err)
		}
	}

	if err := ph.deleteCharacteristics(productID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := ph.createCharacteristics(productID, request.Characteristics); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "successfully updated"})
}

func (ph *ProductHandler) GetProductsByFilter(c *fiber.Ctx) error {
	var request struct {
		Discount      []string `json:"discount"`
		MinPrice      string   `json:"minPrice"`
		MaxPrice      string   `json:"maxPrice"`
		SubcategoryID []string `json:"categoryID"`
		Brand         string   `json:"brand"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	products, err := ph.productUsecase.GetProductByFilter(request.Discount, request.MinPrice, request.MaxPrice, request.SubcategoryID, request.Brand)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (ph *ProductHandler) GetProductsSortedByThreeParams(c *fiber.Ctx) error {
	products, err := ph.productUsecase.GetProductsSorterByThreeParams(c.Query("name"), c.Query("price"), c.Query("discount"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (ph *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	language := c.Params("lang")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	minPrice, _ := strconv.Atoi(c.Query("min"))
	maxPrice, _ := strconv.Atoi(c.Query("max"))
	value := c.Query("value")
	description := c.Query("description")

	var (
		products []entity.Product
		err      error
	)

	switch {
	case limit > 0:
		products, err = ph.productUsecase.GetProductsWithPagination(limit, offset)
	case minPrice > 0 && maxPrice > 0:
		products, err = ph.productUsecase.GetProductsByPriceRange(uint(minPrice), uint(maxPrice), language)
	case value != "" && description != "":
		products, err = ph.productUsecase.GetProductsByCharacteristics(value, description, language)
	default:
		products, err = ph.productUsecase.GetAllProducts(language)
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (ph *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	product, err := ph.productUsecase.GetProductByID(uint(id), c.Params("lang"))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(product)
}

func (ph *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	productID := uint(id)

	images, err := ph.productUsecase.GetImagesByProductID(productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	for _, img := range images {
		if err := ph.deleteImage(img); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	if err := ph.deleteCharacteristics(productID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ph.productUsecase.DeleteProduct(productID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "successfully deleted"})
}

func (ph *ProductHandler) GetProductsByCategory(c *fiber.Ctx) error {
	language := c.Params("lang")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	sortOrder := c.Query("sort")
	var products []entity.Product
	if sortOrder == "cheap" || sortOrder == "expensive" {
		products, err = ph.productUsecase.GetProductsSortedByPriceAndCategory(sortOrder, uint(id), language)
	} else {
		products, err = ph.productUsecase.GetProductsByCategoryID(uint(id), language)
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (ph *ProductHandler) addImage(productID uint, dataURL string) error {
	name, err := saveImageAsWebp(ph.imagesDir, dataURL)
	if err != nil {
		return err
	}

	image := entity.Image{
		ProductID: productID,
		Path:      ph.publicURL + "/images/" + name,
	}
	if err := ph.productUsecase.CreateImage(&image); err != nil {
		removeImage(ph.imagesDir, name)
		return err
	}
	return nil
}

func (ph *ProductHandler) deleteImage(img entity.Image) error {
	if err := removeImage(ph.imagesDir, img.Path); err != nil {
		return err
	}
	return ph.productUsecase.DeleteImage(img.ID)
}

func (ph *ProductHandler) createCharacteristics(productID uint, chars []entity.Characteristic) error {
	for i := range chars {
		chars[i].ID = 0
		chars[i].ProductID = productID
		if err := ph.productUsecase.CreateCharacteristic(&chars[i]); err != nil {
			return err
		}
	}
	return nil
}

func (ph *ProductHandler) deleteCharacteristics(productID uint) error {
	old, err := ph.productUsecase.GetCharacteristicsByProductID(productID)
	if err != nil {
		return err
	}
	for _, ch := range old {
		if err := ph.productUsecase.DeleteCharacteristic(ch.ID); err != nil {
			return err
		}
	}
	return nil
}

func (ph *ProductHandler) imageError(c *fiber.Ctx, err error) error {
	if errors.Is(err, errBadImage) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
}
