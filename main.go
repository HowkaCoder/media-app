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






	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"


	"strings"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"crypto/rand"
)
const (
	loginURL    = "https://notify.eskiz.uz/api/auth/login"
	refreshURL  = "https://notify.eskiz.uz/api/auth/refresh"
	tokenFile   = "token.txt"
	email       = "timajkeenks@gmail.com"
	password    = "IiUx7KGWsjy1L4dScoP7wlOurj9oNobByVdcXx5l"
	refreshTime = 20 * 24 * time.Hour // 20 дней
)
type TokenResponse struct {
	Message   string `json:"message"`
	Data      struct {
		Token string `json:"token"`
	} `json:"data"`
	TokenType string `json:"token_type"`
}
func main() {

	log.Println("starting servers")

	log.Println("database initiating ")
	db := internal.Init()
	log.Println("database initiation complete")




	log.Println("sms шлюз запускается ")
	if err := loginAndSaveToken(); err != nil {
		log.Fatal(err)
	}
	startTokenRefresher()
	log.Println("sms шлюз работает ")


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
  app.Get("/api/:lang/products/name/:name" , productHandler.GetProductsByName)

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


	app.Post("/api/users/phone" , func (c *fiber.Ctx) error {
		var request struct {
			Phone 		string     `phone`
		}

		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		sendSMS(request.Phone)


		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":"sms sended",
		})

	})

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
	// Calculate and save metrics for different periods
	if err := calculateAndSaveMetrics(db, "order_count", "Количество заказов", "Count", "id"); err != nil {
		return fmt.Errorf("error calculating order count metrics: %w", err)
	}

	if err := calculateAndSaveMetrics(db, "total_revenue", "Общая выручка", "SUM", "total_amount"); err != nil {
		return fmt.Errorf("error calculating total revenue metrics: %w", err)
	}

	return nil
}

func calculateAndSaveMetrics(db *gorm.DB, metricType, description, aggregateFunc, column string) error {
	// Define periods and their labels
	periods := []struct {
		label  string
		offset time.Duration
	}{
		{"daily", -24 * time.Hour},
		{"weekly", -7 * 24 * time.Hour},
		{"monthly", -30 * 24 * time.Hour},
		{"yearly", -365 * 24 * time.Hour},
	}

	for _, period := range periods {
		// Prepare the query based on the aggregate function
		var value float64
		query := db.Model(&entity.Order{}).Where("created_at >= ?", time.Now().Add(period.offset)).Where("deleted_at IS NULL")

		if aggregateFunc == "Count" {
			var count int64
			if err := query.Count(&count).Error; err != nil {
				return err
			}
			value = float64(count)
		} else if aggregateFunc == "SUM" {
			if err := query.Select(fmt.Sprintf("COALESCE(SUM(%s), 0)", column)).Scan(&value).Error; err != nil {
				return err
			}
		}

		// Save the metric
		metric := entity.Metric{
			MetricType:  fmt.Sprintf("%s_%s", period.label, metricType),
			Value:       value,
			Date:        time.Now(),
			Description: fmt.Sprintf("%s за %s", description, period.label),
		}
		if err := db.Create(&metric).Error; err != nil {
			return err
		}
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


func loginAndSaveToken() error {
	resp, err := http.Post(loginURL, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(fmt.Sprintf("email=%s&password=%s", email, password))))
	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response failed: %v", err)
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return fmt.Errorf("unmarshal failed: %v", err)
	}

	token := tokenResp.Data.Token
	if err := ioutil.WriteFile(tokenFile, []byte(token), 0644); err != nil {
		return fmt.Errorf("write token failed: %v", err)
	}

	fmt.Println("Token saved successfully")
	return nil
}

func refreshAndSaveToken() error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, refreshURL, nil)
	if err != nil {
		return fmt.Errorf("refresh token failed: %v", err)
	}

	token, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return fmt.Errorf("read token failed: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+string(token))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("refresh request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading refresh response failed: %v", err)
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return fmt.Errorf("unmarshal failed: %v", err)
	}

	newToken := tokenResp.Data.Token
	if err := ioutil.WriteFile(tokenFile, []byte(newToken), 0644); err != nil {
		return fmt.Errorf("write new token failed: %v", err)
	}

	fmt.Println("Token refreshed successfully")
	return nil
}

func startTokenRefresher() {
	ticker := time.NewTicker(refreshTime)
	go func() {
		for {
			<-ticker.C
			if err := refreshAndSaveToken(); err != nil {
				log.Printf("Error refreshing token: %v", err)
			}
		}
	}()
}





const smsURL = "https://notify.eskiz.uz/api/message/sms/send"

func generateCode() string {
	code := make([]byte, 3)
	rand.Read(code)
	return fmt.Sprintf("%06d", int(code[0])%1000000)
}

func sendSMS(phone string) error {
	code := generateCode()

	token, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return fmt.Errorf("read token failed: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", smsURL, strings.NewReader(fmt.Sprintf("mobile_phone=%s&message=%s&from=верефикация кода", phone, code)))
	if err != nil {
		return fmt.Errorf("request creation failed: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+string(token))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending SMS failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to send SMS: %s", string(body))
	}

	// Сохранение кода в Redis
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if err := rdb.Set(ctx, phone, code, 0).Err(); err != nil {
		return fmt.Errorf("failed to save code in Redis: %v", err)
	}

	fmt.Printf("SMS sent successfully to %s\n", phone)
	return nil
}
