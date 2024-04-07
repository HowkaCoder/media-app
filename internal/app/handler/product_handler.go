package handler



import (
  "bytes"
  "encoding/base64"
  "fmt"
  "github.com/gofiber/fiber/v2"
  "github.com/google/uuid"
  "log"
  "media-app/internal/app/entity"
  "media-app/internal/app/usecase"
  "os"
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

    image := entity.Image{
      ProductID: request.Product.ID,
      Path:      fmt.Sprintf("https://media-app-production.up.railway.app/images/%s", fileName),
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

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
  c.Set("Access-Control-Allow-Origin", "*")
  c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
  c.Set("Access-Control-Allow-Headers", "Content-Type")
  c.Set("Access-Control-Allow-Credentials", "true")

  //id, err := strconv.Atoi(c.Params("id"))
  //if err != nil {
  //  return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
  //}

  var request struct {
    Product         entity.Product           `json:"product"`
    Images          []entity.Image          `json:"images"`
    Characteristics []*entity.Characteristic `json:"characteristics"`
  }

  if err := c.BodyParser(&request); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 1": err.Error()})
  }

  log.Println("...............Update Product...............")

  log.Println("\n\n\n ")
  log.Println("...............Request...............")
  log.Println(&request)
  //return c.JSON(fiber.Map{"id": id, "request": request, "product": request.Product})

  

log.Println("...............Request Product ID...............")
  
  log.Println(request.Product.ID)
  
  log.Println("...............Update Product...............")
  
  log.Println(request.Product)
  if err := ph.productUsecase.UpdateProduct(&request.Product, request.Product.ID); err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
  }

  oldImages, err := ph.productUsecase.GetImagesByProductID(request.Product.ID)
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 2": err.Error()})
  }



 oldPaths := make(map[string]entity.Image)
    for _ , img := range oldImages {
        oldPaths[img.Path] = img
    }




    for _, newImg := range request.Images {
        if oldImg, exists := oldPaths[newImg.Path]; !exists || oldImg.Path != newImg.Path {
            // Удаляем старое изображение, если путь изменился
            if exists {
                if err := os.Remove(oldImg.Path); err != nil {
                    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to remove old image"})
                }
                if err := ph.productUsecase.DeleteImage(oldImg.ID); err != nil {
                    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to delete old image record"})
                }
            }


      path := strings.Split(newImg.Path, ",")
    log.Println("path[1]   ", path[1])
    decodedImage, err := base64.StdEncoding.DecodeString(path[1])
    if err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 6": err.Error()})
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
    file, err := os.Create("images/" + fileName)
    if err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    defer file.Close()

    if _, err := file.Write(decodedImage); err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    image := entity.Image{
      ProductID: request.Product.ID,
      Path:      fmt.Sprintf("https://media-app-production.up.railway.app/images/%s", fileName),
    }

if err := ph.productUsecase.CreateImage(&image); err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 1": err.Error()})
    }

          
        }
    }










  
/*  for _, oldImage := range oldImages {
    path := strings.Split(oldImage.Path, "/")
    oldPath := filepath.Join(path[3], path[4])
    log.Println(oldPath)
    
    if err := os.Remove(oldPath); err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
    }

    if err := ph.productUsecase.DeleteImage(oldImage.ID); err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 3 ": err.Error()})
    }
  }
*/
  oldValues, err := ph.productUsecase.GetCharacteristicsByProductID(request.Product.ID)
  if err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 4 ": err.Error()})
  }
  for _, oldValue := range oldValues {
    if err := ph.productUsecase.DeleteCharacteristic(oldValue.ID); err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 5 ": err.Error()})
    }
  }

 /* for _, image := range request.Images {
    path := strings.Split(image.Path, ",")
    log.Println("path[1]   ", path[1])
    decodedImage, err := base64.StdEncoding.DecodeString(path[1])
    if err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error 6": err.Error()})
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
    file, err := os.Create("images/" + fileName)
    if err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    defer file.Close()

    if _, err := file.Write(decodedImage); err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    image := entity.Image{
      ProductID: request.Product.ID,
      Path:      fmt.Sprintf("https://media-app-production.up.railway.app/images/%s", fileName),
    }

if err := ph.productUsecase.CreateImage(&image); err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error 1": err.Error()})
    }
  }
*/
  for _, Value := range request.Characteristics {
    Value.ProductID = request.Product.ID
    
      log.Println("........................Value........................")
      log.Println(Value)
    if err := ph.productUsecase.CreateCharacteristic(Value); err != nil {
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

  } else if minPrice > 0 && maxPrice > 0 {
    products, err = ph.productUsecase.GetProductsByPriceRange(uint(minPrice), uint(maxPrice))
    if err != nil {
      return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
  } else if value != "" && description != "" {
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
