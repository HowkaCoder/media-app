package handler

import (
	"fmt"
	"media-app/internal/app/entity"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type data struct {
	access_token  string `json:"accessToken"`
	refresh_token string `json:"refreshToken"`
	status        string `json:"status"`
	username      string `json:"username"`
}

type errorr struct {
	status  uint   `json:"status"`
	message string `json:"message"`
}

type UsersHandler struct {
	userUsecase usecase.UsersUseCase
	userService service.UserService
}
type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserHandler(useCase usecase.UsersUseCase, userService service.UserService) *UsersHandler {
	return &UsersHandler{userUsecase: useCase, userService: userService}
}

func (uh *UsersHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := uh.userUsecase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}

	return c.JSON(users)
}

func (uh *UsersHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}

	user, err := uh.userUsecase.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}

	return c.JSON(user)
}

func (uh *UsersHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Status": fiber.StatusBadRequest, "data": err.Error()})
	}

	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Status": fiber.StatusBadRequest, "data": err.Error()})
	}
	fmt.Println(user)
	if err := uh.userUsecase.UpdateUser(&user, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "data": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "successfully updated"})
}

func (uh *UsersHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}

	if err := uh.userUsecase.DeleteUser(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{"message": "successfully deleted user"})
}

func (uh *UsersHandler) Register(c *fiber.Ctx) error {
	//
	//file, err := c.FormFile("ava")
	//
	//if err != nil {
	//	log.Println("Error in uploading Image : ", err)
	//	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	//
	//}
	//
	//uniqueId := uuid.New()
	//filename := strings.Replace(uniqueId.String(), "-", "", -1)
	//fileExt := strings.Split(file.Filename, ".")[1]
	//image := fmt.Sprintf("%s.%s", filename, fileExt)
	//
	//// Get absolute path to the images folder
	//imagesDir, err := os.Getwd()
	//if err != nil {
	//	log.Println("Error getting working directory:", err)
	//	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	//}
	//imagesDir = fmt.Sprintf("%s/images", imagesDir)
	//
	//// Create the images folder if it doesn't exist
	//if err := os.MkdirAll(imagesDir, os.ModePerm); err != nil {
	//	log.Println("Error creating images folder:", err)
	//	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	//}
	//
	//// Save the image
	//err = c.SaveFile(file, fmt.Sprintf("%s/%s", imagesDir, image))
	//if err != nil {
	//	log.Println("Error in saving Image :", err, " image ", image)
	//	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	//}

	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return err
	//}
	//
	//user.Password = string(hashedPassword)

	//user.Ava = fmt.Sprintf("https://media-app-production.up.railway.app/images/%s", image)

	if err := uh.userUsecase.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func (uh *UsersHandler) Login(c *fiber.Ctx) error {
	var logins login

	err := c.BodyParser(&logins)
	var eri = errorr{
		status:  http.StatusBadRequest,
		message: err.Error(),
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"data": eri})
	}

	user, err := uh.userUsecase.FindUserByUsername(logins.Username)

	if err != nil {
		eri = errorr{
			status:  http.StatusNotFound,
			message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": eri})
	}

	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logins.Password)); err != nil {
	//	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": "Password doesnt exists"})
	//}

	if user.Password != logins.Password {
		eri = errorr{
			status:  http.StatusBadRequest,
			message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"data": eri})
	}
	accessToken, err := uh.userService.GenerateAccessToken(user)
	if err != nil {
		return err
	}

	refreshToken, err := uh.userService.GenerateRefreshToken(user)
	if err != nil {
		return err
	}

	var data = data{
		access_token:  accessToken,
		refresh_token: refreshToken,
		status:        user.Role,
		username:      user.Username,
	}
	return c.JSON(fiber.Map{
		"data": data,
	})
}

func (uh *UsersHandler) AuthenticateToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.ErrUnauthorized
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &entity.JWTCredentials{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return entity.SecretKey, nil
	})
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if claims, ok := token.Claims.(*entity.JWTCredentials); ok && token.Valid {
		c.Locals("role", claims.Role)

		return c.Next()
	}

	return nil
}

func (uh *UsersHandler) AuthorizeRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("role").(string)
		if !ok || userRole != role {
			return fiber.ErrForbidden
		}
		return c.Next()
	}
}

func (uh *UsersHandler) GetUserProfile(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.ErrUnauthorized
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &entity.JWTCredentials{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return entity.SecretKey, nil
	})
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if claims, ok := token.Claims.(*entity.JWTCredentials); ok && token.Valid {

		return c.JSON(fiber.Map{"claims": claims})
	}

	return fiber.ErrUnauthorized
}
