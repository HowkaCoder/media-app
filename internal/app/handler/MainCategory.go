package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"media-app/internal/app/entity"
	"media-app/internal/app/usecase"
	"strconv"
)

type MainCategoryHandler struct {
	MainCategoryUsecase usecase.MainCategoryUseCase
	SubCategoryUsecase  usecase.SubCategoryUseCase
}

func NewMainCategoryHandler(useCase usecase.MainCategoryUseCase, subcategoryUseCase usecase.SubCategoryUseCase) *MainCategoryHandler {
	return &MainCategoryHandler{MainCategoryUsecase: useCase, SubCategoryUsecase: subcategoryUseCase}
}

func (mh *MainCategoryHandler) GetAllMainCategories(c *fiber.Ctx) error {
	language := c.Params("lang")

	categories, err := mh.MainCategoryUsecase.GetAllMainCategories(language)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error  #1": err.Error()})
	}
	return c.JSON(categories)
}

func (mh *MainCategoryHandler) CreateMainCategory(c *fiber.Ctx) error {
	var request struct {
		MainCategory entity.MainCategory  `json:"main_category"`
		SubCategory  []entity.SubCategory `json:"sub_category"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error  #1": err.Error()})
	}

	log.Println("...............Request...............")
	log.Println("INFO: ", request)

	// создаем MainCategory
	log.Println("...............Create MainCategory...............")
	log.Println("INFO: ", request.MainCategory)
	if err := mh.MainCategoryUsecase.CreateMainCategory(&request.MainCategory); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error  #2": err.Error()})
	}
	log.Println("Crated MainCategory: ", request.MainCategory)

	// создаем SubCategory
	log.Println("...............Create SubCategories...............")
	log.Println("INFO: ", request.SubCategory)
	for _, subCategory := range request.SubCategory {

		subCategory.MainCategoryID = request.MainCategory.ID

		if err := mh.SubCategoryUsecase.CreateSubCategory(&subCategory); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error  #3": err.Error()})
		}

		log.Println("Crated SubCategory: ", request.SubCategory)

	}

	return c.JSON(fiber.Map{"message": "product successfully created"})

}

func (mh *MainCategoryHandler) UpdateMainCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error  #1": err.Error()})
	}

	oldMainCategory, err := mh.MainCategoryUsecase.GetSingleMainCategory(uint(id))

	for _, subCategory := range oldMainCategory.SubCategories {
		if err := mh.SubCategoryUsecase.DeleteSubCategory(subCategory.ID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error  #2": err.Error()})
		}
	}

	if err := mh.MainCategoryUsecase.DeleteMainCategory(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error  #3": err.Error()})
	}

	var request struct {
		MainCategory entity.MainCategory  `json:"main_category"`
		SubCategory  []entity.SubCategory `json:"sub_category"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error  #1": err.Error()})
	}

	log.Println("...............Request...............")
	log.Println("INFO: ", request)

	// создаем MainCategory
	log.Println("...............Create MainCategory...............")
	log.Println("INFO: ", request.MainCategory)
	if err := mh.MainCategoryUsecase.CreateMainCategory(&request.MainCategory); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error  #2": err.Error()})
	}
	log.Println("Crated MainCategory: ", request.MainCategory)

	// создаем SubCategory
	log.Println("...............Create SubCategories...............")
	log.Println("INFO: ", request.SubCategory)
	for _, subCategory := range request.SubCategory {

		subCategory.MainCategoryID = request.MainCategory.ID

		if err := mh.SubCategoryUsecase.CreateSubCategory(&subCategory); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error  #3": err.Error()})
		}

		log.Println("Crated SubCategory: ", request.SubCategory)

	}

	return c.JSON(fiber.Map{"message": "product successfully updated"})

}

func (mh *MainCategoryHandler) DeleteMainCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error  #1": err.Error()})
	}

	oldMainCategory, err := mh.MainCategoryUsecase.GetSingleMainCategory(uint(id))

	for _, subCategory := range oldMainCategory.SubCategories {
		if err := mh.SubCategoryUsecase.DeleteSubCategory(subCategory.ID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error  #2": err.Error()})
		}
	}

	if err := mh.MainCategoryUsecase.DeleteMainCategory(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error  #3": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "product successfully deleted"})
}

func (mh *MainCategoryHandler) GetSingleMainCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error  #1": err.Error()})
	}

	maincategory, err := mh.MainCategoryUsecase.GetSingleMainCategory(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error #2": err.Error()})
	}

	return c.JSON(maincategory)
}
