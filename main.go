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

	log.Println("starting server")

	log.Println("database initiating ")
	db := internal.Init()
	log.Println("database initiation complete")

	// PRODUCT
	productRepository := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUseCase(productRepository)
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

	// Subcategory
	subcategoryRepository := repository.NewSubCategoryRepository(db)
	subcategoryUsecase := usecase.NewSubCategoryUseCase(subcategoryRepository)
	subcategoryHandler := handler.NewSubCategoryHandler(subcategoryUsecase)

	//MainCategory
	mainCategoryRepository := repository.NewMainCategoryRepository(db)
	mainCategoryUsecase := usecase.NewMainCategoryUseCase(mainCategoryRepository)
	mainCategoryHandler := handler.NewMainCategoryHandler(mainCategoryUsecase, subcategoryUsecase)

	// Order
	orderRepository := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUseCase(orderRepository)
	orderHandler := handler.NewOrderHandler(orderUsecase, userUsecase)
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
	app.Get("/:lang/api/products", productHandler.GetAllProducts)
	app.Get("/:lang/api/products/:id", productHandler.GetProductByID)
	app.Get("/:lang/api/categories/:id/products", productHandler.GetProductsByCategory)

	app.Get("/api/subcategories", subcategoryHandler.GetAllSubCategories)
	app.Get("/api/subcategories/:id", subcategoryHandler.GetSubCategoryByID)
	app.Post("/api/subcategories", subcategoryHandler.CreateSubCategory)
	app.Patch("/api/subcategories/:id", subcategoryHandler.UpdateSubCategory)
	app.Delete("/api/subcategories/:id", subcategoryHandler.DeleteSubCategory)

	app.Get("/:lang/api/maincategories", mainCategoryHandler.GetAllMainCategories)
	app.Post("/:lang/api/maincategories", mainCategoryHandler.CreateMainCategory)
	app.Get("/:lang/api/maincategories/:id", mainCategoryHandler.GetSingleMainCategory)
	app.Patch("/:lang/api/maincategories/:id", mainCategoryHandler.UpdateMainCategory)
	app.Delete("/:lang/api/maincategories/:id", mainCategoryHandler.DeleteMainCategory)

	app.Get("/api/languages", langHandler.GetAllLanguages)
	app.Get("/api/languages/:id", langHandler.GetLanguageByID)
	app.Get("/api/products/:product_id/translations", translationHandler.GetProductTranslationsByProductID)
	app.Get("/api/characteristics/:characteristic_id/translations", translationHandler.GetCharacteristicTranslationsByCharacteristicID)

	app.Get("/api/orders", orderHandler.GetAllOrders)
	app.Get("/api/orders/:id", orderHandler.GetOrderById)
	app.Post("/api/orders", orderHandler.CreateOrder)
	app.Patch("/api/orders/:id", orderHandler.UpdateOrder)
	app.Delete("/api/orders/:id", orderHandler.DeleteOrder)

	api := app.Group("/api", userHandler.AuthenticateToken)
	//api := app.Group("/api")
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
