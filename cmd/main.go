package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"media-app/internal"
	"media-app/internal/app/handler"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
)

func main() {

	db := internal.Init()

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService()
	categoryUsecase := usecase.NewCategoryUseCase(categoryRepository, categoryService)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(categoryRepository)
	productUsecase := usecase.NewProductUseCase(productRepository, productService)
	productHandler := handler.NewProductHandler(productUsecase)

	langRepo := repository.NewLanguageRepository(db)
	langUsecase := usecase.NewLanguageUseCase(langRepo)
	langHandler := handler.NewLanguageHandler(langUsecase)

	translationRepository := repository.NewTranslationRepository(db)
	translationUsecase := usecase.NewTranslationUseCase(translationRepository)
	translationHandler := handler.NewTranslationHandler(translationUsecase)

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})
	api := app.Group("/api")
	api.Get("/categories", categoryHandler.GetAllCategories)
	api.Post("/categories", categoryHandler.CreateCategory)
	api.Patch("/categories/:id", categoryHandler.UpdateCategory)
	api.Get("/categories/:id", categoryHandler.GetCategoryByID)
	api.Delete("/categories/:id", categoryHandler.DeleteCategory)

	api.Get("/products", productHandler.GetAllProducts)
	api.Post("/products", productHandler.CreateProduct)
	api.Patch("/products/:id", productHandler.UpdateProduct)
	api.Get("/products/:id", productHandler.GetProductByID)
	api.Delete("/products/:id", productHandler.DeleteProduct)
	api.Get("/categories/:id/products", productHandler.GetProductsByCategory)

	api.Post("/images", productHandler.CreateImages)

	api.Get("/languages", langHandler.GetAllLanguages)
	api.Get("/languages/:id", langHandler.GetLanguageByID)
	api.Post("/languages", langHandler.CreateLanguage)
	api.Patch("/languages/:id", langHandler.UpdateLanguage)
	api.Delete("/languages/:id", langHandler.DeleteLanguage)

	api.Get("/products/:product_id/translations", translationHandler.GetProductTranslationsByProductID)
	api.Post("/translations/product", translationHandler.CreateProductTranslation)
	api.Patch("/translations/product/:id", translationHandler.UpdateProductTranslation)
	api.Delete("/translations/product/:id", translationHandler.DeleteProductTranslation)

	api.Get("/characteristics/:characteristic_id/translations", translationHandler.GetCharacteristicTranslationsByCharacteristicID)
	api.Post("/translations/characteristic", translationHandler.CreateCharacteristicTranslation)
	api.Patch("/translations/characteristic/:id", translationHandler.UpdateCharacteristicTranslation)
	api.Delete("/translations/characteristic/:id", translationHandler.DeleteCharacteristicTranslation)

	log.Println("Server is runnig on :8082")
	log.Fatal(app.Listen(":8082"))

}
