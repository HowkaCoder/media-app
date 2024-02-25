package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
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

//func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
//
//	var request struct {
//		Product         entity.Product          `json:"product"`
//		Images          []entity.Image          `json:"images"`
//		Characteristics []entity.Characteristic `json:"characteristics"`
//	}
//
//	if err := c.BodyParser(&request); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 1": err.Error(), "request": request})
//	} else {
//		return c.JSON(fiber.Map{"request": request})
//	}
//
//if err := ph.productUsecase.CreateProduct(&request.Product); err != nil {
//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 3 ": err.Error(), "request": request})
//}
//
//form, err := c.MultipartForm()
//if err != nil {
//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 2 ": err.Error()})
//}
//
//images := form.File["images[]"]
//
//var base64Images []string
//
//for _, imageFile := range images {
//	file, err := imageFile.Open()
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to open image file", "details": err.Error()})
//	}
//	defer file.Close()
//
//	fileContent, err := ioutil.ReadAll(file)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to read image file", "details": err.Error()})
//	}
//
//	base64Image := base64.StdEncoding.EncodeToString(fileContent)
//	base64Images = append(base64Images, base64Image)
//}
//
//for _, imageData := range base64Images {
//	decodedImage, err := base64.StdEncoding.DecodeString(imageData)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to decode base64 image", "details": err.Error()})
//	}
//
//	photoPath := filepath.Join("uploads", "photo", fmt.Sprintf("%s.jpg", uuid.New().String())) // Generate a unique filename
//	if err := ioutil.WriteFile(photoPath, decodedImage, 0644); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save photo", "details": err.Error()})
//	}
//
//	image := entity.Image{
//		ProductID: request.Product.ID,
//		Path:      photoPath,
//	}
//
//	if err := ph.productUsecase.CreateImage(&image); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 1": err.Error()})
//	} else {
//		log.Println(request.Product.ID)
//	}
//}
//for _, chars := range request.Characteristics {
//	chars.ProductID = request.Product.ID
//	log.Println(chars)
//	if err := ph.productUsecase.CreateCharacteristic(&chars); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 2": err.Error()})
//	}
//
//	log.Println(chars)
//
//}
//return c.JSON(fiber.Map{"request": &request})

//}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	// Парсинг | анализирует данные, извлекает из них нужную информацию
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// JSON | Эта строка извлекает данные JSON из формы, которые клиент отправил как json.
	jsonData := form.Value["json"][0]

	var request struct {
		Product         entity.Product           `json:"product"`
		Images          []entity.Image           `json:"images"`
		Characteristics []*entity.Characteristic `json:"characteristics"`
	}
	// Этот блок кода разбирает данные JSON, полученные от клиента, и помещает их в структуру request.
	if err := json.Unmarshal([]byte(jsonData), &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// создает Product
	if err := ph.productUsecase.CreateProduct(&request.Product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 3 ": err.Error(), "request": request})
	}

	for _, image := range request.Images {

		// BASE64Decoding
		imageData, err := base64.URLEncoding.DecodeString(image.Path)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to decode base64 image", "details": err.Error()})
		}

		// Generate a unique filename
		photoPath := filepath.Join("uploads", "photo", fmt.Sprintf("%s.jpg", uuid.New().String()))
		if err := ioutil.WriteFile(photoPath, imageData, 0644); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save photo", "details": err.Error()})
		}

		image := entity.Image{
			ProductID: request.Product.ID,
			Path:      photoPath,
		}

		if err := ph.productUsecase.CreateImage(&image); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 1": err.Error()})
		}
	}

	for _, chars := range request.Characteristics {
		chars.ProductID = request.Product.ID
		log.Println(chars)
		if err := ph.productUsecase.CreateCharacteristic(chars); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 2": err.Error()})
		}

		log.Println(chars)

	}
	return c.JSON(fiber.Map{"request": &request})
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
		if err := ph.productUsecase.CreateImage(&image); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

	}

	for _, Value := range request.Characteristics {
		if err := ph.productUsecase.CreateCharacteristic(&Value); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		} else {
			log.Println(Value)
		}
	}

	return c.JSON(fiber.Map{"message": "successfully updated"})

}

func (ph *ProductHandler) GetAllProducts(c *fiber.Ctx) error {

	limit, _ := strconv.Atoi(c.Query("limit"))

	minPrice, _ := strconv.Atoi(c.Query("min"))

	maxPrice, _ := strconv.Atoi(c.Query("max"))

	value := c.Query("value")

	description := c.Query("description")

	var products []entity.Product
	var err error

	if limit > 0 {
		products, err = ph.productUsecase.GetProductsWithPagination(limit)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

	}
	if minPrice > 0 || maxPrice > 0 {
		products, err = ph.productUsecase.GetProductsByPriceRange(uint(minPrice), uint(maxPrice))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}
	if value != "" || description != "" {
		products, err = ph.productUsecase.GetProductsByCharacteristics(value, description)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}

	products, err = ph.productUsecase.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.JSON(products)
}

func (ph *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	product, err := ph.productUsecase.GetProductByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": id})
	}
	return c.JSON(product)
}

func (ph *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

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

	oldValues, err := ph.productUsecase.GetCharacteristicsByProductID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	for _, oldValue := range oldValues {
		if err := ph.productUsecase.DeleteCharacteristic(oldValue.ProductID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	if err := ph.productUsecase.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "successfully deleted"})

}

func (ph *ProductHandler) GetProductsByCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	sortOrder := c.Query("sort")

	if sortOrder != "cheap" && sortOrder != "expensive" {
		products, err := ph.productUsecase.GetProductsByCategoryID(uint(id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}
		return c.JSON(products)
	}

	products, err := ph.productUsecase.GetProductsSortedByPriceAndCategory(sortOrder, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)

}
