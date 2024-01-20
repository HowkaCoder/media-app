package main

import (
	"log"
	"media-app/internal"
	"media-app/internal/app/handler"
	"media-app/internal/app/repository"
	"media-app/internal/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db := internal.Init()

	categoryRepository := repository.NewCategoryRepository(db)

	categoryUsecase := usecase.NewCategoryUseCase(categoryRepository)

	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	app := fiber.New()
	app.Get("/categories", categoryHandler.GetAllCategories)

	log.Println("Server is runnig on :8080")
	log.Fatal(app.Listen(":8080"))

}
