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

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type")
		c.Set("Access-Control-Allow-Credentials", "true")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
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

	log.Println("Server is runnig on :8082")
	log.Fatal(app.Listen(":8082"))

}
