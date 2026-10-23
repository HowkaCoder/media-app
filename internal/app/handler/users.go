package handler

import (
	"context"
	"errors"
	"log"
	"media-app/internal/app/entity"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
	"media-app/internal/otp"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var phoneRe = regexp.MustCompile(`^\+?998\d{9}$`)

type UsersHandler struct {
	userUsecase usecase.UsersUseCase
	userService service.UserService
	otp         *otp.Store
	sender      Sender
}

type Sender interface {
	Send(ctx context.Context, phone, text string) error
}

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerRequest struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       uint   `json:"age"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Password  string `json:"password"`
	Ava       string `json:"ava"`
	Code      string `json:"code"`
}

type updateUserRequest struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       uint   `json:"age"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Ava       string `json:"ava"`
}

func NewUserHandler(useCase usecase.UsersUseCase, userService service.UserService, otpStore *otp.Store, sender Sender) *UsersHandler {
	return &UsersHandler{userUsecase: useCase, userService: userService, otp: otpStore, sender: sender}
}

func (uh *UsersHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := uh.userUsecase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (uh *UsersHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if !canAccessUser(c, uint(id)) {
		return fiber.ErrForbidden
	}

	user, err := uh.userUsecase.GetUserByID(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (uh *UsersHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if !canAccessUser(c, uint(id)) {
		return fiber.ErrForbidden
	}

	var req updateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if req.Role != "" && c.Locals("role") != "admin" {
		return fiber.ErrForbidden
	}

	user := entity.User{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Age:       req.Age,
		Phone:     req.Phone,
		Address:   req.Address,
		Role:      req.Role,
		Ava:       req.Ava,
	}
	if req.Password != "" {
		hash, err := uh.userService.HashPassword(req.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		user.Password = hash
	}

	if err := uh.userUsecase.UpdateUser(&user, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "successfully updated"})
}

func (uh *UsersHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := uh.userUsecase.DeleteUser(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "successfully deleted user"})
}

func (uh *UsersHandler) SendCode(c *fiber.Ctx) error {
	if uh.otp == nil || uh.sender == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "sms is disabled"})
	}

	var req struct {
		Phone string `json:"phone"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if !phoneRe.MatchString(req.Phone) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid phone"})
	}

	code, err := uh.otp.Issue(c.Context(), req.Phone)
	if errors.Is(err, otp.ErrTooManyRequests) {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": err.Error()})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := uh.sender.Send(c.Context(), req.Phone, "Kod podtverzhdeniya: "+code); err != nil {
		log.Printf("send sms to %s: %v", req.Phone, err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "failed to send sms"})
	}

	return c.JSON(fiber.Map{"message": "sms sent"})
}

func (uh *UsersHandler) Register(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if req.Username == "" || len(req.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username is required and password must be at least 8 characters"})
	}

	if uh.otp != nil {
		ok, err := uh.otp.Verify(c.Context(), req.Phone, req.Code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid or expired code"})
		}
	}

	hash, err := uh.userService.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	user := entity.User{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Age:       req.Age,
		Phone:     req.Phone,
		Address:   req.Address,
		Password:  hash,
		Role:      "user",
		Ava:       req.Ava,
	}
	if err := uh.userUsecase.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "successfully created user", "id": user.ID})
}

func (uh *UsersHandler) Login(c *fiber.Ctx) error {
	var req login
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := uh.userUsecase.FindUserByUsername(req.Username)
	if err != nil || !uh.userService.CheckPassword(user.Password, req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid username or password"})
	}

	accessToken, err := uh.userService.GenerateAccessToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	refreshToken, err := uh.userService.GenerateRefreshToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
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
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return fiber.ErrUnauthorized
	}

	claims, err := uh.userService.ParseToken(strings.TrimPrefix(authHeader, "Bearer "))
	if err != nil {
		return fiber.ErrUnauthorized
	}

	c.Locals("user_id", claims.UserID)
	c.Locals("role", claims.Role)
	return c.Next()
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
	id, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.ErrUnauthorized
	}

	user, err := uh.userUsecase.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}

func canAccessUser(c *fiber.Ctx, id uint) bool {
	if c.Locals("role") == "admin" {
		return true
	}
	userID, ok := c.Locals("user_id").(uint)
	return ok && userID == id
}
