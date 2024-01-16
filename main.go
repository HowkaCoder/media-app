package main

import (
	"fmt"
	"log"
	"media-app/handler"
	"media-app/model"
	"media-app/repository"
	"media-app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate используется для создания таблицы в базе данных
	db.AutoMigrate(&model.User{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	// Обработчики для CRUD API
	app.Get("/users", userHandler.GetUsers)
	app.Get("/users/:id", userHandler.GetUser)
	app.Post("/users", userHandler.CreateUser)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)

	port := 3000
	err = app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
