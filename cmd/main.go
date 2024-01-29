package main

import (
	"log"
	"media-app/internal"
	"media-app/internal/app/handler"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db := internal.Init()

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService()
	categoryUsecase := usecase.NewCategoryUseCase(categoryRepository, categoryService)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	app := fiber.New()
	api := app.Group("/api")
	api.Get("/categories", categoryHandler.GetAllCategories)
	api.Post("/categories", categoryHandler.CreateCategory)
	api.Patch("/categories/:id", categoryHandler.UpdateCategory)
	api.Get("/categories/:id", categoryHandler.GetCategoryByID)
	api.Delete("/categories/:id", categoryHandler.DeleteCategory)

	log.Println("Server is runnig on :8081")
	log.Fatal(app.Listen(":8081"))

}
