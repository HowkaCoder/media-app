package handler

import (
	"context"
	"fmt"
	"log"
	"media-app/internal/app/entity"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

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
func verifyCode(phone, code string) bool {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	storedCode, err := rdb.Get(ctx, phone).Result()
	if err == redis.Nil {
		fmt.Println("Code not found for this phone number")
		return false
	} else if err != nil {
		log.Fatalf("Failed to get code from Redis: %v", err)
	}

	return storedCode == code
}
func (uh *UsersHandler) Register(c *fiber.Ctx) error {

	var user entity.User

	var request struct {
		Username  string ` json:"username"`
		Firstname string ` json:"firstname"`
		Lastname  string ` json:"lastname"`
		Age       uint   ` json:"age"`
		Phone     string ` json:"phone"`
		Address   string ` json:"address"`
		Password  string ` json:"password"`
		Role      string ` json:"role"`
		Ava       string ` json:"ava"`
		Code      string `  json:"code"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if verifyCode(request.Phone, request.Code) {
		user.Username = request.Username
		user.Firstname = request.Firstname
		user.Lastname = request.Lastname
		user.Age = request.Age
		user.Phone = request.Phone
		user.Address = request.Address
		user.Role = request.Role
		user.Ava = request.Ava

		if err := uh.userUsecase.CreateUser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "User registered successfully",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "User registered error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "..."})

}

func (uh *UsersHandler) Login(c *fiber.Ctx) error {
	var logins login

	err := c.BodyParser(&logins)

	if err != nil {
		var eri = errorr{
			status:  http.StatusBadRequest,
			message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"data": eri})
	}

	user, err := uh.userUsecase.FindUserByUsername(logins.Username)

	if err != nil {
		var eri = errorr{
			status:  http.StatusNotFound,
			message: err.Error(),
		}
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": eri})
	}

	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logins.Password)); err != nil {
	//	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": "Password doesnt exists"})
	//}

	if user.Password != logins.Password {
		var eri = errorr{
			status:  http.StatusBadRequest,
			message: err.Error(),
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"data 2": eri})
	}
	accessToken, err := uh.userService.GenerateAccessToken(user)
	if err != nil {
		return err
	}

	refreshToken, err := uh.userService.GenerateRefreshToken(user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{

		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"status":        user.Role,
		"username":      user.Username,
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
