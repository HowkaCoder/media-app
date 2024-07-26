package handler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"image/jpeg"
	"image/png"
	"log"
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ProductHandler struct {
	productUsecase usecase.ProductUseCase
}

func NewProductHandler(useCase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUsecase: useCase}
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {

	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	c.Set("Access-Control-Allow-Credentials", "true")

	var request struct {
		Product         entity.Product           `json:"product"`
		Images          []entity.Image           `json:"images"`
		Characteristics []*entity.Characteristic `json:"characteristics"`
	}

	log.Println("...............Create Product...............")

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 1": err.Error()})
	}

	log.Println("...............Request...............")
	log.Println("INFO: ", request)

	// создает Product
	log.Println("...............Create Product...............")
	log.Println("INFO: ", request.Product)
	if err := ph.productUsecase.CreateProduct(&request.Product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 2 ": err.Error(), "request": request})
	}

	log.Println("...............Request Images...............")
	log.Println("INFO: ", request.Images)
	for _, image := range request.Images {

		path := strings.Split(image.Path, ",")
		log.Println("...............Base64URL of Image...............")
		log.Println(path[1])
		decodedImage, err := base64.StdEncoding.DecodeString(path[1])

		log.Println("...............Decoded Image...............")
		log.Println("INFO: ", decodedImage)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		var imageFormat string
		switch {
		case bytes.HasPrefix(decodedImage, []byte{0xFF, 0xD8}):
			imageFormat = ".jpg"
		case bytes.HasPrefix(decodedImage, []byte{0x89, 0x50, 0x4E, 0x47}):
			imageFormat = ".png"
		case bytes.HasPrefix(decodedImage, []byte{0x47, 0x49, 0x46, 0x38}):
			imageFormat = ".gif"
		case bytes.HasPrefix(decodedImage, []byte{0x42, 0x4D}):
			imageFormat = ".bmp"
		case bytes.HasPrefix(decodedImage, []byte{0x52, 0x49, 0x46, 0x46}): // TIFF
			imageFormat = ".tiff"
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported image format"})
		}
		fileName := uuid.New().String() + imageFormat
		FileName := uuid.New().String() + ".webp"
		log.Println("...............Image FileName...............")

		file, err := os.Create("images/" + fileName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer file.Close()

		log.Println("...............Empty File...............")
		log.Println("INFO: ", file)
		if _, err := file.Write(decodedImage); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		file1, err := os.Open("images/" + fileName)
		if err != nil {
			log.Fatalln(err)
		}

		if imageFormat == ".jpg" {
			img1, err := jpeg.Decode(file1)

			if err != nil {
				log.Fatalln(err)
			}
			output, err := os.Create("images/" + FileName)
			if err != nil {
				log.Fatal(err)
			}
			defer output.Close()

			options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
			if err != nil {
				log.Fatalln(err)
			}

			if err := webp.Encode(output, img1, options); err != nil {
				log.Fatalln(err)
			}

			if err := os.Remove("images/" + fileName); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #3": err.Error()})
			}

		}

		if imageFormat == ".png" {
			img1, err := png.Decode(file1)

			if err != nil {
				log.Fatalln(err)
			}
			log.Println(FileName)
			output, err := os.Create("images/" + FileName)
			if err != nil {
				log.Fatal(err)
			}
			defer output.Close()

			options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
			if err != nil {
				log.Fatalln(err)
			}

			if err := webp.Encode(output, img1, options); err != nil {
				log.Fatalln(err)
			}

			if err := os.Remove("images/" + fileName); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #3": err.Error()})
			}

		}

		image := entity.Image{
			ProductID: request.Product.ID,
			Path:      fmt.Sprintf("https://media-app-production.up.railway.app/images/%s", FileName),
		}

		log.Println("...............Entity Image...............")
		log.Println("INFO: ", image)
		if err := ph.productUsecase.CreateImage(&image); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 1": err.Error()})
		}
	}

	log.Println("...............Request Characteristics...............")
	log.Println("INFO: ", request.Characteristics)
	for _, chars := range request.Characteristics {
		chars.ProductID = request.Product.ID

		if err := ph.productUsecase.CreateCharacteristic(chars); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 2": err.Error()})
		}

		log.Println("INFO: ", chars)

	}

	return c.JSON(fiber.Map{"message": "product successfully created"})

}

//
//func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
//
//	id, err := strconv.Atoi(c.Params("id"))
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
//	}
//
//	var request struct {
//		Product         entity.Product          `json:"product"`
//		Images          []entity.Image          `json:"images"`
//		Characteristics []entity.Characteristic `json:"characteristics"`
//	}
//
//	if err := c.BodyParser(&request); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//	}
//	//return c.JSON(fiber.Map{"id": id, "request": request, "product": request.Product})
//
//	if err := ph.productUsecase.UpdateProduct(&request.Product, uint(id)); err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
//	}
//
//	log.Println("...............Old Images Removing...............")
//	//oldImages, err := ph.productUsecase.GetImagesByProductID(uint(id))
//	//if err != nil {
//	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//	//}
//	//for _, oldImage := range oldImages {
//	//	path := strings.Split(oldImage.Path, "/")
//	//	oldPath := filepath.Join(path[3], path[4])
//	//	log.Println(oldPath)
//	//	if err := os.Remove(oldPath); err != nil {
//	//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
//	//	}
//	//
//	//	if err := ph.productUsecase.DeleteImage(oldImage.ID); err != nil {
//	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//	//	}
//	//}
//
//	oldImages, err := ph.productUsecase.GetImagesByProductID(request.Product.ID)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 2": err.Error()})
//	}
//
//	oldPaths := make(map[string]entity.Image)
//	for _, img := range oldImages {
//		oldPaths[img.Path] = img
//	}
//
//	for _, newImg := range request.Images {
//		if oldImg, exists := oldPaths[newImg.Path]; !exists || oldImg.Path != newImg.Path {
//			// Удаляем старое изображение, если путь изменился
//			if exists {
//				if err := os.Remove(oldImg.Path); err != nil {
//					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to remove old image"})
//				}
//				if err := ph.productUsecase.DeleteImage(oldImg.ID); err != nil {
//					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to delete old image record"})
//				}
//			}
//
//			path := strings.Split(newImg.Path, ",")
//			log.Println("path[1]   ", path[1])
//			decodedImage, err := base64.StdEncoding.DecodeString(path[1])
//			if err != nil {
//				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 6": err.Error()})
//			}
//			var imageFormat string
//			switch {
//			case bytes.HasPrefix(decodedImage, []byte{0xFF, 0xD8}):
//				imageFormat = ".jpg"
//			case bytes.HasPrefix(decodedImage, []byte{0x89, 0x50, 0x4E, 0x47}):
//				imageFormat = ".png"
//			case bytes.HasPrefix(decodedImage, []byte{0x47, 0x49, 0x46, 0x38}):
//				imageFormat = ".gif"
//			case bytes.HasPrefix(decodedImage, []byte{0x42, 0x4D}):
//				imageFormat = ".bmp"
//			case bytes.HasPrefix(decodedImage, []byte{0x52, 0x49, 0x46, 0x46}): // TIFF
//				imageFormat = ".tiff"
//			default:
//				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported image format"})
//			}
//			fileName := uuid.New().String() + imageFormat
//			file, err := os.Create("images/" + fileName)
//			if err != nil {
//				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
//			}
//			defer file.Close()
//
//			if _, err := file.Write(decodedImage); err != nil {
//				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
//			}
//
//			image := entity.Image{
//				ProductID: request.Product.ID,
//				Path:      fmt.Sprintf("https://media-app-production.up.railway.app/images/%s", fileName),
//			}
//
//			if err := ph.productUsecase.CreateImage(&image); err != nil {
//				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 1": err.Error()})
//			}
//
//		}
//	}
//
//	log.Println("...............Old Characteristics Removing...............")
//	oldValues, err := ph.productUsecase.GetCharacteristicsByProductID(request.Product.ID)
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//	}
//	for _, oldValue := range oldValues {
//		if err := ph.productUsecase.DeleteCharacteristic(oldValue.ProductID); err != nil {
//			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//		}
//	}
//
//	log.Println("...............Update Request Images...............")
//	log.Println("INFO: ", request.Images)
//	for _, image := range request.Images {
//
//		path := strings.Split(image.Path, ",")
//		log.Println("...............Base64URL of Image...............")
//		log.Println(path[1])
//		decodedImage, err := base64.StdEncoding.DecodeString(path[1])
//
//		log.Println("...............Decoded Image...............")
//		log.Println("INFO: ", decodedImage)
//		if err != nil {
//			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//		}
//		var imageFormat string
//		switch {
//		case bytes.HasPrefix(decodedImage, []byte{0xFF, 0xD8}):
//			imageFormat = ".jpg"
//		case bytes.HasPrefix(decodedImage, []byte{0x89, 0x50, 0x4E, 0x47}):
//			imageFormat = ".png"
//		case bytes.HasPrefix(decodedImage, []byte{0x47, 0x49, 0x46, 0x38}):
//			imageFormat = ".gif"
//		case bytes.HasPrefix(decodedImage, []byte{0x42, 0x4D}):
//			imageFormat = ".bmp"
//		case bytes.HasPrefix(decodedImage, []byte{0x52, 0x49, 0x46, 0x46}): // TIFF
//			imageFormat = ".tiff"
//		default:
//			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported image format"})
//		}
//		fileName := uuid.New().String() + imageFormat
//		log.Println("...............Image FileName...............")
//
//		file, err := os.Create("images/" + fileName)
//		if err != nil {
//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
//		}
//		defer file.Close()
//
//		log.Println("...............Empty File...............")
//		log.Println("INFO: ", file)
//		if _, err := file.Write(decodedImage); err != nil {
//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
//		}
//
//		image := entity.Image{
//			ProductID: request.Product.ID,
//			Path:      fmt.Sprintf("https://media-app-production.up.railway.app/images/%s", fileName),
//		}
//
//		log.Println("...............Entity Image...............")
//		log.Println("INFO: ", image)
//		if err := ph.productUsecase.CreateImage(&image); err != nil {
//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 1": err.Error()})
//		}
//	}
//
//	log.Println("...............Update Request Characteristics...............")
//	log.Println("INFO: ", request.Characteristics)
//	for _, chars := range request.Characteristics {
//		chars.ProductID = request.Product.ID
//
//		if err := ph.productUsecase.CreateCharacteristic(&chars); err != nil {
//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 2": err.Error()})
//		}
//
//		log.Println("INFO: ", chars)
//
//	}
//
//	return c.JSON(fiber.Map{"message": "successfully updated"})
//
//}

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var request struct {
		Product         entity.Product          `json:"product"`
		Images          []entity.Image          `json:"images"`
		Characteristics []entity.Characteristic `json:"characteristics"`
	}

	log.Println("....................................UPDATE PRODUCT STARTS....................................")

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error #1": err.Error(),
		})
	}

	log.Println("....................................PRODUCT UPDATING ....................................")

	if err := ph.productUsecase.UpdateProduct(&request.Product, request.Product.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #2": err.Error()})
	}

	log.Println("....................................PRODUCT UPDATING 1....................................")

	oldImages, err := ph.productUsecase.GetImagesByProductID(request.Product.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error #3": err.Error()})
	}

	log.Println("....................................PRODUCT UPDATING 2....................................")
	toDelete := []uint{}
	toAdd := []entity.Image{}

	for _, oldImage := range oldImages {
		found := false
		for _, eImage := range request.Images {
			if eImage.ID == oldImage.ID {
				found = true
				if err := ph.productUsecase.UpdateImage(eImage, eImage.ID); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #4": err.Error()})
				}
				break
			}
		}
		if !found {
			toDelete = append(toDelete, oldImage.ID)
		}
	}

	log.Println("....................................PRODUCT UPDATING 3....................................")
	for _, requestImage := range request.Images {
		log.Println("INFO: ", requestImage)
		if requestImage.ID == 0 {
			toAdd = append(toAdd, requestImage)
		}
		log.Println("INFO: ", toAdd)
	}

	log.Println("....................................PRODUCT UPDATING 4....................................")

	if len(toDelete) > 0 {

		log.Println("...............Old Images Removing...............")
		for _, i := range toDelete {

			dImage, err := ph.productUsecase.GetImageByID(i)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #5": err.Error()})
			}
			path := strings.Split(dImage.Path, "/")
			oldPath := filepath.Join(path[3], path[4])
			log.Println(oldPath)
			if err := os.Remove(oldPath); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #6": err.Error()})
			}

			if err := ph.productUsecase.DeleteImage(dImage.ID); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error #7": err.Error()})
			}

		}
	}

	log.Println("....................................PRODUCT UPDATING 5....................................")
	for _, image := range toAdd {
		log.Println("...............New Images Adding...............")

		path := strings.Split(image.Path, ",")
		log.Println("...............Base64URL of Image...............")
		log.Println(path[1])
		decodedImage, err := base64.StdEncoding.DecodeString(path[1])

		log.Println("...............Decoded Image...............")
		log.Println("INFO: ", decodedImage)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error #8": err.Error()})
		}
		var imageFormat string
		switch {
		case bytes.HasPrefix(decodedImage, []byte{0xFF, 0xD8}):
			imageFormat = ".jpg"
		case bytes.HasPrefix(decodedImage, []byte{0x89, 0x50, 0x4E, 0x47}):
			imageFormat = ".png"
		case bytes.HasPrefix(decodedImage, []byte{0x47, 0x49, 0x46, 0x38}):
			imageFormat = ".gif"
		case bytes.HasPrefix(decodedImage, []byte{0x42, 0x4D}):
			imageFormat = ".bmp"
		case bytes.HasPrefix(decodedImage, []byte{0x52, 0x49, 0x46, 0x46}): // TIFF
			imageFormat = ".tiff"
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported image format"})
		}
		fileName := uuid.New().String() + imageFormat
		FileName := uuid.New().String() + ".webp"
		log.Println("...............Image FileName...............")

		file, err := os.Create("images/" + fileName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		defer file.Close()

		log.Println("...............Empty File...............")
		log.Println("INFO: ", file)
		if _, err := file.Write(decodedImage); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		log.Println("here/////")

		file1, err := os.Open("images/" + fileName)
		if err != nil {
			log.Fatalln(err)
		}

		if imageFormat == ".jpg" {
			img1, err := jpeg.Decode(file1)

			if err != nil {
				log.Fatalln(err)
			}
			output, err := os.Create("images/" + FileName)
			if err != nil {
				log.Fatal(err)
			}
			defer output.Close()

			options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
			if err != nil {
				log.Fatalln(err)
			}

			if err := webp.Encode(output, img1, options); err != nil {
				log.Fatalln(err)
			}

			if err := os.Remove("images/" + fileName); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #3": err.Error()})
			}

		}

		if imageFormat == ".png" {
			img1, err := png.Decode(file1)

			if err != nil {
				log.Fatalln(err)
			}
			log.Println(FileName)
			output, err := os.Create("images/" + FileName)
			if err != nil {
				log.Fatal(err)
			}
			defer output.Close()

			options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
			if err != nil {
				log.Fatalln(err)
			}

			if err := webp.Encode(output, img1, options); err != nil {
				log.Fatalln(err)
			}

			if err := os.Remove("images/" + fileName); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #3": err.Error()})
			}

		}

		image := entity.Image{
			ProductID: request.Product.ID,
			Path:      fmt.Sprintf("https://media-app-production.up.railway.app/images/%s", FileName),
		}

		log.Println("...............Entity Image...............")
		log.Println("INFO: ", image)
		if err := ph.productUsecase.CreateImage(&image); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 1": err.Error()})
		}

	}

	log.Println("....................................PRODUCT UPDATING 6....................................")

	log.Println("...............Old Characteristics Removing...............")
	oldValues, err := ph.productUsecase.GetCharacteristicsByProductID(request.Product.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	for _, oldValue := range oldValues {
		if err := ph.productUsecase.DeleteCharacteristic(oldValue.ID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	}

	log.Println("...............Update Request Characteristics...............")
	log.Println("INFO: ", request.Characteristics)
	for _, chars := range request.Characteristics {
		chars.ProductID = request.Product.ID
		chars.ID = 0

		if err := ph.productUsecase.CreateCharacteristic(&chars); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 2": err.Error()})
		}

		log.Println("INFO: ", chars)

	}

	return c.JSON(fiber.Map{
		"message": "successfully updated",
	})
}




func (ph *ProductHandler) GetProductsByFilter(c *fiber.Ctx) error {
//	param , _ := strconv.Atoi( c.Query("discount"))
	//minPrice , _ := strconv.Atoi(c.Query("minPrice"))
//	maxPrice , _ := strconv.Atoi(c.Query("maxPrice"))
	//subcategoryID , _ := strconv.Atoi(c.Query("categoryID"))


 var request struct {
    Discount        []string    `json:"discount"`
    MinPrice        string       `json:"minPrice"`
    MaxPrice        string       `json:"maxPrice"`
    SubcategoryID   []string     `json:"categoryID"`
    Brand           string 	 `json:"brand"`
  }




  if err := c.BodyParser(&request); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error":err.Error()})
  }
  fmt.Println(request)
  fmt.Println(&request)


	fmt.Println(request.SubcategoryID)
	products , err := ph.productUsecase.GetProductByFilter(request.Discount , request.MinPrice , request.MaxPrice , request.SubcategoryID , request.Brand)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error":err.Error()})
	}

	return c.JSON(products)
}


func (ph *ProductHandler) GetProductsSortedByThreeParams(c *fiber.Ctx) error {
	name := c.Query("name")
	price := c.Query("price")
	discount := c.Query("discount")

	products , err := ph.productUsecase.GetProductsSorterByThreeParams(name , price , discount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 1":err.Error})
	}

	return c.JSON(products)
}





func (ph *ProductHandler) GetAllProducts(c *fiber.Ctx) error {

	language := c.Params("lang")

	limit, _ := strconv.Atoi(c.Query("limit"))

	offseti, _ := strconv.Atoi(c.Query("offset"))

	minPrice, _ := strconv.Atoi(c.Query("min"))

	maxPrice, _ := strconv.Atoi(c.Query("max"))

	value := c.Query("value")

	description := c.Query("description")

	var products []entity.Product
	var err error

	if limit > 0 {
		//return c.JSON(fiber.Map{"limit": limit})
		products, err = ph.productUsecase.GetProductsWithPagination(int(limit), int(offseti))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

	} else if minPrice > 0 && maxPrice > 0 {
		products, err = ph.productUsecase.GetProductsByPriceRange(uint(minPrice), uint(maxPrice), language)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	} else if value != "" && description != "" {
		products, err = ph.productUsecase.GetProductsByCharacteristics(value, description, language)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	} else {

		products, err = ph.productUsecase.GetAllProducts(language)
		if err != nil {
			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"Error": err.Error()})
		}

	}
	return c.JSON(products)
}

func (ph *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	language := c.Params("lang")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	product, err := ph.productUsecase.GetProductByID(uint(id), language)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.JSON(product)
}

func (ph *ProductHandler) DeleteProduct(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error #1": err.Error()})
	}

	oldImages, err := ph.productUsecase.GetImagesByProductID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #2": err.Error()})
	}

	log.Println("...............Old Images Removing...............")
	for _, i := range oldImages {

		path := strings.Split(i.Path, "/")
		oldPath := filepath.Join(path[3], path[4])
		log.Println(oldPath)
		if err := os.Remove(oldPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #3": err.Error()})
		}

		if err := ph.productUsecase.DeleteImage(i.ID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error #4": err.Error()})
		}

	}

	log.Println("...............Old Characteristics Removing...............")
	oldValues, err := ph.productUsecase.GetCharacteristicsByProductID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error #5": err.Error()})
	}
	for _, oldValue := range oldValues {
		if err := ph.productUsecase.DeleteCharacteristic(oldValue.ID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error #6": err.Error()})
		}
	}

	if err := ph.productUsecase.DeleteProduct(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #7": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "successfully deleted"})

}

func (ph *ProductHandler) GetProductsByCategory(c *fiber.Ctx) error {
	language := c.Params("lang")
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	sortOrder := c.Query("sort")

	if sortOrder != "cheap" && sortOrder != "expensive" {
		products, err := ph.productUsecase.GetProductsByCategoryID(uint(id), language)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}
		return c.JSON(products)
	}

	products, err := ph.productUsecase.GetProductsSortedByPriceAndCategory(sortOrder, uint(id), language)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)

}
