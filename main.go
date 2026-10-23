package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"media-app/internal"
	"media-app/internal/app/entity"
	"media-app/internal/app/handler"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
	"media-app/internal/config"
	"media-app/internal/otp"
	"media-app/internal/sms"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := internal.Init(cfg.DBPath)
	if err != nil {
		log.Fatalf("database: %v", err)
	}

	if err := os.MkdirAll(cfg.ImagesDir, 0o755); err != nil {
		log.Fatalf("images dir: %v", err)
	}

	var (
		otpStore *otp.Store
		sender   handler.Sender
	)
	if cfg.SMSEnabled() {
		rdb := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
		if err := rdb.Ping(ctx).Err(); err != nil {
			log.Fatalf("redis: %v", err)
		}
		otpStore = otp.NewStore(rdb)

		eskiz := sms.NewEskiz(cfg.SMSEmail, cfg.SMSPassword)
		if err := eskiz.Login(ctx); err != nil {
			log.Fatalf("sms: %v", err)
		}
		eskiz.StartRefresher(ctx, 20*24*time.Hour)
		sender = eskiz
	} else {
		log.Println("sms is disabled: ESKIZ_EMAIL or ESKIZ_PASSWORD is not set")
	}

	userService := service.NewUserService([]byte(cfg.JWTSecret))

	productHandler := handler.NewProductHandler(usecase.NewProductUseCase(repository.NewProductRepository(db)), cfg.ImagesDir, cfg.PublicURL)
	langHandler := handler.NewLanguageHandler(usecase.NewLanguageUseCase(repository.NewLanguageRepository(db)))
	translationHandler := handler.NewTranslationHandler(usecase.NewTranslationUseCase(repository.NewTranslationRepository(db)))
	userHandler := handler.NewUserHandler(usecase.NewUsersUseCase(repository.NewUserRepository(db)), userService, otpStore, sender)
	subcategoryUsecase := usecase.NewSubCategoryUseCase(repository.NewSubCategoryRepository(db))
	subcategoryHandler := handler.NewSubCategoryHandler(subcategoryUsecase)
	mainCategoryHandler := handler.NewMainCategoryHandler(usecase.NewMainCategoryUseCase(repository.NewMainCategoryRepository(db)), subcategoryUsecase)
	orderHandler := handler.NewOrderHandler(usecase.NewOrderUseCase(repository.NewOrderRepository(db)))

	app := fiber.New(fiber.Config{BodyLimit: 100 * 1024 * 1024})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	app.Static("/images", cfg.ImagesDir)

	auth := userHandler.AuthenticateToken
	admin := userHandler.AuthorizeRole("admin")

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)
	app.Post("/api/users/phone", userHandler.SendCode)

	app.Get("/metrics", auth, admin, metricsHandler(db))

	app.Get("/api/:lang/products", productHandler.GetAllProducts)
	app.Get("/api/:lang/products/:id", productHandler.GetProductByID)
	app.Get("/api/:lang/categories/:id/products", productHandler.GetProductsByCategory)
	app.Post("/api/:lang/products-filter", productHandler.GetProductsByFilter)
	app.Get("/api/:lang/products/name/:name", productHandler.GetProductsByName)
	app.Get("/api/:lang/productss", productHandler.GetProductsSortedByThreeParams)
	app.Post("/api/:lang/products", auth, admin, productHandler.CreateProduct)
	app.Patch("/api/:lang/products/:id", auth, admin, productHandler.UpdateProduct)
	app.Delete("/api/:lang/products/:id", auth, admin, productHandler.DeleteProduct)

	app.Get("/api/subcategories", subcategoryHandler.GetAllSubCategories)
	app.Get("/api/subcategories/:id", subcategoryHandler.GetSubCategoryByID)
	app.Post("/api/subcategories", auth, admin, subcategoryHandler.CreateSubCategory)
	app.Patch("/api/subcategories/:id", auth, admin, subcategoryHandler.UpdateSubCategory)
	app.Delete("/api/subcategories/:id", auth, admin, subcategoryHandler.DeleteSubCategory)

	app.Get("/:lang/api/maincategories", mainCategoryHandler.GetAllMainCategories)
	app.Get("/:lang/api/maincategories/:id", mainCategoryHandler.GetSingleMainCategory)
	app.Post("/:lang/api/maincategories", auth, admin, mainCategoryHandler.CreateMainCategory)
	app.Patch("/:lang/api/maincategories/:id", auth, admin, mainCategoryHandler.UpdateMainCategory)
	app.Delete("/:lang/api/maincategories/:id", auth, admin, mainCategoryHandler.DeleteMainCategory)

	app.Get("/api/languages", langHandler.GetAllLanguages)
	app.Get("/api/languages/:id", langHandler.GetLanguageByID)
	app.Post("/api/languages", auth, admin, langHandler.CreateLanguage)
	app.Patch("/api/languages/:id", auth, admin, langHandler.UpdateLanguage)
	app.Delete("/api/languages/:id", auth, admin, langHandler.DeleteLanguage)

	app.Get("/api/products/:product_id/translations", translationHandler.GetProductTranslationsByProductID)
	app.Get("/api/characteristics/:characteristic_id/translations", translationHandler.GetCharacteristicTranslationsByCharacteristicID)
	app.Post("/api/translations/product", auth, admin, translationHandler.CreateProductTranslation)
	app.Patch("/api/translations/product/:id", auth, admin, translationHandler.UpdateProductTranslation)
	app.Delete("/api/translations/product/:id", auth, admin, translationHandler.DeleteProductTranslation)
	app.Post("/api/translations/characteristic", auth, admin, translationHandler.CreateCharacteristicTranslation)
	app.Patch("/api/translations/characteristic/:id", auth, admin, translationHandler.UpdateCharacteristicTranslation)
	app.Delete("/api/translations/characteristic/:id", auth, admin, translationHandler.DeleteCharacteristicTranslation)

	app.Post("/api/orders", auth, orderHandler.CreateOrder)
	app.Get("/api/orders", auth, admin, orderHandler.GetAllOrders)
	app.Get("/api/orders/:id", auth, admin, orderHandler.GetOrderByID)
	app.Patch("/api/orders/:id", auth, admin, orderHandler.UpdateOrder)
	app.Delete("/api/orders/:id", auth, admin, orderHandler.DeleteOrder)

	app.Get("/api/users-profile", auth, userHandler.GetUserProfile)
	app.Get("/api/users", auth, admin, userHandler.GetAllUsers)
	app.Get("/api/users/:id", auth, userHandler.GetUserByID)
	app.Patch("/api/users/:id", auth, userHandler.UpdateUser)
	app.Delete("/api/users/:id", auth, admin, userHandler.DeleteUser)

	go runMetrics(ctx, db, 24*time.Hour)

	go func() {
		<-ctx.Done()
		if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
			log.Printf("shutdown: %v", err)
		}
	}()

	log.Printf("server is running on :%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}

func runMetrics(ctx context.Context, db *gorm.DB, every time.Duration) {
	ticker := time.NewTicker(every)
	defer ticker.Stop()
	for {
		if err := calculateMetrics(db); err != nil {
			log.Printf("metrics: %v", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func calculateMetrics(db *gorm.DB) error {
	if err := saveMetrics(db, "order_count", "Количество заказов", "COUNT(id)"); err != nil {
		return fmt.Errorf("order count: %w", err)
	}
	if err := saveMetrics(db, "total_revenue", "Общая выручка", "COALESCE(SUM(total_amount), 0)"); err != nil {
		return fmt.Errorf("total revenue: %w", err)
	}
	return nil
}

func saveMetrics(db *gorm.DB, metricType, description, aggregate string) error {
	periods := []struct {
		label  string
		offset time.Duration
	}{
		{"daily", 24 * time.Hour},
		{"weekly", 7 * 24 * time.Hour},
		{"monthly", 30 * 24 * time.Hour},
		{"yearly", 365 * 24 * time.Hour},
	}

	now := time.Now()
	for _, p := range periods {
		var value float64
		err := db.Model(&entity.Order{}).
			Where("created_at >= ?", now.Add(-p.offset)).
			Select(aggregate).
			Scan(&value).Error
		if err != nil {
			return err
		}

		metric := entity.Metric{
			MetricType:  p.label + "_" + metricType,
			Value:       value,
			Date:        now,
			Description: description + " за " + p.label,
		}
		if err := db.Create(&metric).Error; err != nil {
			return err
		}
	}
	return nil
}

func metricsHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var metrics []entity.Metric
		if err := db.Order("date desc").Limit(500).Find(&metrics).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(metrics)
	}
}
