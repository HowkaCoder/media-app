package main

import (
	"fmt"
	"log"
	"media-app/internal"
  "media-app/internal/app/entity"
	"media-app/internal/app/handler"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
	"os"
	"path/filepath"
  "gorm.io/gorm"
  "time"
	"github.com/gofiber/fiber/v2"
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
	orderHandler := handler.NewOrderHandler(orderUsecase)
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




  app.Get("/metrics", GetMetricsHandler(db))


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
	app.Get("/api/:lang/products", productHandler.GetAllProducts)
	app.Get("/api/:lang/products/:id", productHandler.GetProductByID)
	app.Get("/api/:lang/categories/:id/products", productHandler.GetProductsByCategory)
	app.Post("/api/:lang/products-filter", productHandler.GetProductsByFilter)

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

	app.Get("/api/:lang/productss", productHandler.GetProductsSortedByThreeParams)
	app.Get("/api/languages", langHandler.GetAllLanguages)
	app.Get("/api/languages/:id", langHandler.GetLanguageByID)
	app.Get("/api/products/:product_id/translations", translationHandler.GetProductTranslationsByProductID)
	app.Get("/api/characteristics/:characteristic_id/translations", translationHandler.GetCharacteristicTranslationsByCharacteristicID)

	app.Get("/api/orders", orderHandler.GetAllOrders)
	app.Get("/api/orders/:id", orderHandler.GetOrderByID)
	app.Post("/api/orders", orderHandler.CreateOrder)
	app.Patch("/api/orders/:id", orderHandler.UpdateOrder)
	app.Delete("/api/orders/:id", orderHandler.DeleteOrder)

	api := app.Group("/api", userHandler.AuthenticateToken)
	//api := app.Group("/api")
	api.Post("/:lang/products", userHandler.AuthorizeRole("admin"), productHandler.CreateProduct)
	api.Patch("/:lang/products/:id", userHandler.AuthorizeRole("admin"), productHandler.UpdateProduct)
	api.Delete("/:lang/products/:id", userHandler.AuthorizeRole("admin"), productHandler.DeleteProduct)

	api.Post("/languages", userHandler.AuthorizeRole("admin"), langHandler.CreateLanguage)
	api.Patch("/languages/:id", userHandler.AuthorizeRole("admin"), langHandler.UpdateLanguage)
	api.Delete("/languages/:id", userHandler.AuthorizeRole("admin"), langHandler.DeleteLanguage)

	api.Post("/translations/product", userHandler.AuthorizeRole("admin"), translationHandler.CreateProductTranslation)
	api.Patch("/translations/product/:id", userHandler.AuthorizeRole("admin"), translationHandler.UpdateProductTranslation)
	api.Delete("/translations/product/:id", userHandler.AuthorizeRole("admin"), translationHandler.DeleteProductTranslation)

	api.Post("/translations/characteristic", userHandler.AuthorizeRole("admin"), translationHandler.CreateCharacteristicTranslation)
	api.Patch("/translations/characteristic/:id", userHandler.AuthorizeRole("admin"), translationHandler.UpdateCharacteristicTranslation)
	api.Delete("/translations/characteristic/:id", userHandler.AuthorizeRole("admin"), translationHandler.DeleteCharacteristicTranslation)

	api.Get("/users-profile", userHandler.GetUserProfile)
	api.Get("/users", userHandler.AuthorizeRole("admin"), userHandler.GetAllUsers)
	api.Get("/users/:id", userHandler.GetUserByID)
	api.Patch("/users/:id", userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.AuthorizeRole("admin"), userHandler.DeleteUser)






go func() {
		for {
			err := CalculateMetrics(db)
			if err != nil {
				log.Println("Error calculating metrics:", err)
			}
			time.Sleep(24 * time.Hour) // Выполнение каждые 24 часа
		}
	}()









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





func CalculateMetrics(db *gorm.DB) error {
	// Пример вычисления общего количества заказов за день
	var orderCount int64
	err := db.Model(&entity.Order{}).Where("created_at >= ?", time.Now().AddDate(0, 0, -1)).Count(&orderCount).Error
	if err != nil {
		return err
	}

	// Сохранение метрики
	metric := entity.Metric{
		MetricType:  "order_count",
		Value:       float64(orderCount),
		Date:        time.Now(),
		Description: "Количество заказов за день",
	}
	err = db.Create(&metric).Error
	if err != nil {
		return err
	}

	// Пример вычисления общей выручки за день
	var totalRevenue float64
	err = db.Model(&entity.Order{}).Select("SUM(total_amount)").Where("created_at >= ?", time.Now().AddDate(0, 0, -1)).Scan(&totalRevenue).Error
	if err != nil {
		return err
	}

	// Сохранение метрики
	metric = entity.Metric{
		MetricType:  "total_revenue",
		Value:       totalRevenue,
		Date:        time.Now(),
		Description: "Общая выручка за день",
	}
	err = db.Create(&metric).Error
	if err != nil {
		return err
	}















	// Подсчет количества заказов за неделю
var weeklyOrderCount int64
err = db.Model(&entity.Order{}).Where("created_at >= ?", time.Now().AddDate(0, 0, -7)).Count(&weeklyOrderCount).Error
if err != nil {
    return err
}

// Сохранение метрики количества заказов за неделю
weeklyMetric := entity.Metric{
    MetricType:  "weekly_order_count",
    Value:       float64(weeklyOrderCount),
    Date:        time.Now(),
    Description: "Количество заказов за неделю",
}
err = db.Create(&weeklyMetric).Error
if err != nil {
    return err
}

// Подсчет общей выручки за неделю
var weeklyTotalRevenue float64
err = db.Model(&entity.Order{}).Select("SUM(total_amount)").Where("created_at >= ?", time.Now().AddDate(0, 0, -7)).Scan(&weeklyTotalRevenue).Error
if err != nil {
    return err
}

// Сохранение метрики общей выручки за неделю
weeklyMetric = entity.Metric{
    MetricType:  "weekly_total_revenue",
    Value:       weeklyTotalRevenue,
    Date:        time.Now(),
    Description: "Общая выручка за неделю",
}
err = db.Create(&weeklyMetric).Error
if err != nil {
    return err
}














	// Подсчет количества заказов за месяц
var monthlyOrderCount int64
err = db.Model(&entity.Order{}).Where("created_at >= ?", time.Now().AddDate(0, -1, 0)).Count(&monthlyOrderCount).Error
if err != nil {
    return err
}

// Сохранение метрики количества заказов за месяц
monthlyMetric := entity.Metric{
    MetricType:  "monthly_order_count",
    Value:       float64(monthlyOrderCount),
    Date:        time.Now(),
    Description: "Количество заказов за месяц",
}
err = db.Create(&monthlyMetric).Error
if err != nil {
    return err
}

// Подсчет общей выручки за месяц
var monthlyTotalRevenue float64
err = db.Model(&entity.Order{}).Select("SUM(total_amount)").Where("created_at >= ?", time.Now().AddDate(0, -1, 0)).Scan(&monthlyTotalRevenue).Error
if err != nil {
    return err
}

// Сохранение метрики общей выручки за месяц
monthlyMetric = entity.Metric{
    MetricType:  "monthly_total_revenue",
    Value:       monthlyTotalRevenue,
    Date:        time.Now(),
    Description: "Общая выручка за месяц",
}
err = db.Create(&monthlyMetric).Error
if err != nil {
    return err
}

















	// Подсчет количества заказов за год
var yearlyOrderCount int64
err = db.Model(&entity.Order{}).Where("created_at >= ?", time.Now().AddDate(-1, 0, 0)).Count(&yearlyOrderCount).Error
if err != nil {
    return err
}

// Сохранение метрики количества заказов за год
yearlyMetric := entity.Metric{
    MetricType:  "yearly_order_count",
    Value:       float64(yearlyOrderCount),
    Date:        time.Now(),
    Description: "Количество заказов за год",
}
err = db.Create(&yearlyMetric).Error
if err != nil {
    return err
}

// Подсчет общей выручки за год
var yearlyTotalRevenue float64
err = db.Model(&entity.Order{}).Select("SUM(total_amount)").Where("created_at >= ?", time.Now().AddDate(-1, 0, 0)).Scan(&yearlyTotalRevenue).Error
if err != nil {
    return err
}

// Сохранение метрики общей выручки за год
yearlyMetric = entity.Metric{
    MetricType:  "yearly_total_revenue",
    Value:       yearlyTotalRevenue,
    Date:        time.Now(),
    Description: "Общая выручка за год",
}
err = db.Create(&yearlyMetric).Error
if err != nil {
    return err
}

	return nil
}






func GetMetricsHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var metrics []entity.Metric
		if err := db.Find(&metrics).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(metrics)
	}
}
