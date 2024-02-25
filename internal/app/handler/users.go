package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"media-app/internal/app/entity"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
	"os"
	"strconv"
	"strings"
)

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	age, _ := strconv.Atoi(form.Value["age"][0])
	phone, _ := strconv.Atoi(form.Value["phone"][0])
	user := &entity.User{
		Username:  form.Value["username"][0],
		Firstname: form.Value["firstname"][0],
		Lastname:  form.Value["lastname"][0],
		Age:       uint(age),
		Phone:     uint(phone),
		Address:   form.Value["address"][0],
		Password:  form.Value["password"][0],
		Role:      form.Value["role"][0],
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err := uh.userUsecase.UpdateUser(user, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}
	image, err := uh.userUsecase.GetAvaByUserID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := os.Remove(image.Path); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := uh.userUsecase.DeleteAva(image.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	photoFile, err := c.FormFile("avatar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
	}

	photoPath := "uploads/photo/" + photoFile.Filename
	if err := c.SaveFile(photoFile, photoPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save photo"})
	}
	if err := uh.userUsecase.CreateAva(&entity.Ava{UserID: uint(id), Path: photoPath}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}

	return c.JSON(fiber.Map{"message": "Successfully updated user"})
}

func (uh *UsersHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}

	image, err := uh.userUsecase.GetAvaByUserID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"Error": err})
	}
	if err := os.Remove(image.Path); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := uh.userUsecase.DeleteAva(image.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := uh.userUsecase.DeleteUser(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{"message": "successfully deleted user"})
}

func (uh *UsersHandler) CreateAva(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}

	photoFile, err := c.FormFile("avatar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
	}

	photoPath := "uploads/photo/" + photoFile.Filename
	if err := c.SaveFile(photoFile, photoPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save photo"})
	}
	if err := uh.userUsecase.CreateAva(&entity.Ava{UserID: uint(id), Path: photoPath}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}

	return c.JSON(fiber.Map{"message": "Successfully created ava"})

}

func (uh *UsersHandler) DeleteAvaByUserID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	image, err := uh.userUsecase.GetAvaByUserID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := os.Remove(image.Path); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	if err := uh.userUsecase.DeleteAva(image.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Ava successfully deleted"})
}

func (uh *UsersHandler) Register(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	age, _ := strconv.Atoi(form.Value["age"][0])
	phone, _ := strconv.Atoi(form.Value["phone"][0])
	user := entity.User{
		Username:  form.Value["username"][0],
		Firstname: form.Value["firstname"][0],
		Lastname:  form.Value["lastname"][0],
		Age:       uint(age),
		Phone:     uint(phone),
		Address:   form.Value["address"][0],
		Password:  form.Value["password"][0],
		Role:      form.Value["role"][0],
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	photoFile, err := c.FormFile("avatar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid form data"})
	}

	photoPath := "cmd/uploads/photo/" + photoFile.Filename

	if err := c.SaveFile(photoFile, photoPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to save photo", "Err": err.Error()})
	}

	//file, err := photoFile.Open()
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to open image file", "details": err.Error()})
	//}
	//defer file.Close()
	//
	//fileContent, err := ioutil.ReadAll(file)
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to read image file", "details": err.Error()})
	//}
	//
	//base64Image := base64.StdEncoding.EncodeToString(fileContent)
	//
	//decodedImage, err := base64.StdEncoding.DecodeString(base64Image)
	//if err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to decode base64 image", "details": err.Error()})
	//}
	//
	//photoPath := filepath.Join("uploads", "photo", fmt.Sprintf("%s.jpg", uuid.New().String())) // Generate a unique filename
	//if err := ioutil.WriteFile(photoPath, decodedImage, 0644); err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save photo", "details": err.Error()})
	//}
	//
	//if err := uh.userUsecase.CreateUser(user); err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	//}

	ava := &entity.Ava{
		UserID: user.ID,
		Path:   photoPath,
	}
	if err := uh.userUsecase.CreateAva(ava); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}

	if err := uh.userUsecase.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (uh *UsersHandler) Login(c *fiber.Ctx) error {
	var logins login
	if err := c.BodyParser(&logins); err != nil {
		return err
	}

	user, err := uh.userUsecase.FindUserByUsername(logins.Username)
	if err != nil {
		return c.JSON(fiber.Map{"Error": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logins.Password)); err != nil {
		return c.JSON(fiber.Map{"Error": "Password Error"})
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
		"user":          user,
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
