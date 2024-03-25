package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"media-app/internal"
	"media-app/internal/app/handler"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
	"os"
	"path/filepath"
)

func main() {

	db := internal.Init()

	// CATEGORY
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService()
	categoryUsecase := usecase.NewCategoryUseCase(categoryRepository, categoryService)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	// PRODUCT
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(categoryRepository)
	productUsecase := usecase.NewProductUseCase(productRepository, productService)
	productHandler := handler.NewProductHandler(productUsecase)

	// LANGUAGE
	langRepo := repository.NewLanguageRepository(db)
	langUsecase := usecase.NewLanguageUseCase(langRepo)
	langHandler := handler.NewLanguageHandler(langUsecase)

	// TRANSLATION
	translationRepository := repository.NewTranslationRepository(db)
	translationUsecase := usecase.NewTranslationUseCase(translationRepository)
	translationHandler := handler.NewTranslationHandler(translationUsecase)

	// USER
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService()
	userUsecase := usecase.NewUsersUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase, userService)

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE , PATCH")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}
		return c.Next()
	})
	// Получаем абсолютный путь к папке images
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущей директории:", err)
		return
	}
	imagesDir := filepath.Join(currentDir, "images")

	// Статический обработчик для папки с изображениями
	app.Static("/images", imagesDir)

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)
	app.Get("/boom", func(c *fiber.Ctx) error {
		io := "DROP DATABASE IF EXISTS railway"

		res := internal.DB.Exec(io)
		if res.Error != nil {
			return res.Error
		}
		io = "CREATE DATABASE railway "
		res = internal.DB.Exec(io)
		if res.Error != nil {
			return res.Error
		}
		internal.Init()

		return c.SendString("Database reborned")
	})
	app.Get("/api/categories", categoryHandler.GetAllCategories)
	app.Get("/api/categories/:id", categoryHandler.GetCategoryByID)
	app.Get("/api/products", productHandler.GetAllProducts)
	app.Get("/api/products/:id", productHandler.GetProductByID)
	app.Get("/api/categories/:id/products", productHandler.GetProductsByCategory)
	app.Get("/api/languages", langHandler.GetAllLanguages)
	app.Get("/api/languages/:id", langHandler.GetLanguageByID)
	app.Get("/api/products/:product_id/translations", translationHandler.GetProductTranslationsByProductID)
	app.Get("/api/characteristics/:characteristic_id/translations", translationHandler.GetCharacteristicTranslationsByCharacteristicID)

	api := app.Group("/api", userHandler.AuthenticateToken)
	//api := app.Group("/api")
	api.Post("/categories", userHandler.AuthorizeRole("admin"), categoryHandler.CreateCategory)
	api.Patch("/categories/:id", userHandler.AuthorizeRole("admin"), categoryHandler.UpdateCategory)
	api.Delete("/categories/:id", userHandler.AuthorizeRole("admin"), categoryHandler.DeleteCategory)

	api.Post("/products", userHandler.AuthorizeRole("admin"), productHandler.CreateProduct)
	api.Patch("/products/:id", userHandler.AuthorizeRole("admin"), productHandler.UpdateProduct)
	api.Delete("/products/:id", userHandler.AuthorizeRole("admin"), productHandler.DeleteProduct)

	api.Post("/languages", userHandler.AuthorizeRole("admin"), langHandler.CreateLanguage)
	api.Patch("/languages/:id", userHandler.AuthorizeRole("admin"), langHandler.UpdateLanguage)
	api.Delete("/languages/:id", userHandler.AuthorizeRole("admin"), langHandler.DeleteLanguage)

	api.Post("/translations/product", userHandler.AuthorizeRole("admin"), translationHandler.CreateProductTranslation)
	api.Patch("/translations/product/:id", userHandler.AuthorizeRole("admin"), translationHandler.UpdateProductTranslation)
	api.Delete("/translations/product/:id", userHandler.AuthorizeRole("admin"), translationHandler.DeleteProductTranslation)

	api.Post("/translations/characteristic", userHandler.AuthorizeRole("admin"), translationHandler.CreateCharacteristicTranslation)
	api.Patch("/translations/characteristic/:id", userHandler.AuthorizeRole("admin"), translationHandler.UpdateCharacteristicTranslation)
	api.Delete("/translations/characteristic/:id", userHandler.AuthorizeRole("admin"), translationHandler.DeleteCharacteristicTranslation)

	api.Get("/users-profile", userHandler.AuthorizeRole("user"), userHandler.GetUserProfile)
	api.Get("/users", userHandler.AuthorizeRole("admin"), userHandler.GetAllUsers)
	api.Get("/users/:id", userHandler.GetUserByID)
	api.Patch("/users/:id", userHandler.AuthorizeRole("user"), userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.AuthorizeRole("user"), userHandler.DeleteUser)

	log.Println("Server is runnig on " + getPort())
	log.Fatal(app.Listen(getPort()))

}
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8082"
	} else {
		port = ":" + port
	}

	return port
}
