package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"media-app/internal/app/handler"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"media-app/internal"
)

func TestCategoryHandler_GetAllCategories(t *testing.T) {
	app := setupTestApp()

	req, _ := http.NewRequest("GET", "https://media-app-97ql.onrender.com/api/categories", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	fmt.Println(resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCategoryHandler_GetCategoryByID(t *testing.T) {
	app := setupTestApp()

	req, _ := http.NewRequest("GET", "/api/categories/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func setupTestApp() *fiber.App {
	db := internal.Init()
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService()
	categoryUsecase := usecase.NewCategoryUseCase(categoryRepository, categoryService)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	app := fiber.New()

	app.Get("/api/categories", categoryHandler.GetAllCategories)
	app.Get("/api/categories/:id", categoryHandler.GetCategoryByID)
	app.Post("/api/categories", categoryHandler.CreateCategory)
	app.Patch("/api/categories/:id", categoryHandler.UpdateCategory)
	app.Delete("/api/categories/:id", categoryHandler.DeleteCategory)

	return app
}
