package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"strconv"
)

type ProductHandler struct {
	productUsecase usecase.ProductUseCase
}

func NewProductHandler(useCase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUsecase: useCase}
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	//// Парсинг | анализирует данные, извлекает из них нужную информацию
	//form, err := c.MultipartForm()
	//if err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	//}
	//
	//// JSON | Эта строка извлекает данные JSON из формы, которые клиент отправил как json.
	//jsonData := form.Value["json"][0]
	//
	//var request struct {
	//	Product         entity.Product           `json:"product"`
	//	Images          []entity.Image           `json:"images"`
	//	Characteristics []*entity.Characteristic `json:"characteristics"`
	//}
	//// Этот блок кода разбирает данные JSON, полученные от клиента, и помещает их в структуру request.
	//if err := json.Unmarshal([]byte(jsonData), &request); err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	//}

	var request struct {
		Product         entity.Product           `json:"product"`
		Images          []entity.Image           `json:"images"`
		Characteristics []*entity.Characteristic `json:"characteristics"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// создает Product
	if err := ph.productUsecase.CreateProduct(&request.Product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 3 ": err.Error(), "request": request})
	}

	for _, image := range request.Images {

		image := entity.Image{
			ProductID: request.Product.ID,
			Path:      image.Path,
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

	// Обновляет продукт
	if err := ph.productUsecase.UpdateProduct(request.Product, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	// Обновляет  фотки
	for _, imageFile := range request.Images {

		if err := ph.productUsecase.UpdateImage(imageFile, imageFile.ID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

	}

	// Обновляет характеристики
	for _, Value := range request.Characteristics {
		if err := ph.productUsecase.UpdateCharacteristic(Value, Value.ID); err != nil {
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
		//return c.JSON(fiber.Map{"limit": limit})
		products, err = ph.productUsecase.GetProductsWithPagination(int(limit))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

	} else if minPrice > 0 || maxPrice > 0 {
		products, err = ph.productUsecase.GetProductsByPriceRange(uint(minPrice), uint(maxPrice))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	} else if value != "" || description != "" {
		products, err = ph.productUsecase.GetProductsByCharacteristics(value, description)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	} else {

		products, err = ph.productUsecase.GetAllProducts()
		if err != nil {
			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"Error": err.Error()})
		}

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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
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
