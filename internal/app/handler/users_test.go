package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"media-app/internal"
	"media-app/internal/app/repository"
	"media-app/internal/app/service"
	"media-app/internal/app/usecase"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func setupUsersApp(t *testing.T) *fiber.App {
	t.Helper()
	db, err := internal.Init(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}

	h := NewUserHandler(
		usecase.NewUsersUseCase(repository.NewUserRepository(db)),
		service.NewUserService([]byte("test-secret")),
		nil,
		nil,
	)

	app := fiber.New()
	app.Post("/register", h.Register)
	app.Post("/login", h.Login)
	app.Get("/api/users-profile", h.AuthenticateToken, h.GetUserProfile)
	app.Patch("/api/users/:id", h.AuthenticateToken, h.UpdateUser)
	app.Get("/api/users", h.AuthenticateToken, h.AuthorizeRole("admin"), h.GetAllUsers)
	return app
}

func doJSON(t *testing.T, app *fiber.App, method, url, token string, body any) (int, map[string]any) {
	t.Helper()
	var r io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		r = bytes.NewReader(b)
	}
	req := httptest.NewRequest(method, url, r)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	out := map[string]any{}
	json.NewDecoder(resp.Body).Decode(&out)
	return resp.StatusCode, out
}

func register(t *testing.T, app *fiber.App, username string) uint {
	t.Helper()
	status, body := doJSON(t, app, http.MethodPost, "/register", "", map[string]any{
		"username":  username,
		"password":  "password123",
		"firstname": "Test",
		"lastname":  "User",
		"phone":     "99890" + fmt.Sprint(len(username)) + "000000",
		"address":   "Tashkent",
		"age":       20,
	})
	if status != http.StatusCreated {
		t.Fatalf("register %s: status %d, body %v", username, status, body)
	}
	return uint(body["id"].(float64))
}

func loginAs(t *testing.T, app *fiber.App, username, password string) (int, string) {
	t.Helper()
	status, body := doJSON(t, app, http.MethodPost, "/login", "", map[string]any{
		"username": username,
		"password": password,
	})
	token, _ := body["access_token"].(string)
	return status, token
}

func TestLogin(t *testing.T) {
	app := setupUsersApp(t)
	register(t, app, "ali")

	status, token := loginAs(t, app, "ali", "password123")
	if status != http.StatusOK || token == "" {
		t.Fatalf("login failed: %d", status)
	}

	if status, _ := loginAs(t, app, "ali", "wrong-password"); status != http.StatusUnauthorized {
		t.Fatalf("wrong password: expected 401, got %d", status)
	}
	if status, _ := loginAs(t, app, "nobody", "password123"); status != http.StatusUnauthorized {
		t.Fatalf("unknown user: expected 401, got %d", status)
	}
}

func TestRegisterAlwaysCreatesUserRole(t *testing.T) {
	app := setupUsersApp(t)
	doJSON(t, app, http.MethodPost, "/register", "", map[string]any{
		"username": "hacker",
		"password": "password123",
		"role":     "admin",
	})
	_, token := loginAs(t, app, "hacker", "password123")

	if status, _ := doJSON(t, app, http.MethodGet, "/api/users", token, nil); status != http.StatusForbidden {
		t.Fatalf("expected 403 for admin route, got %d", status)
	}
}

func TestPasswordIsNotReturned(t *testing.T) {
	app := setupUsersApp(t)
	register(t, app, "ali")
	_, token := loginAs(t, app, "ali", "password123")

	status, body := doJSON(t, app, http.MethodGet, "/api/users-profile", token, nil)
	if status != http.StatusOK {
		t.Fatalf("profile: %d", status)
	}
	if _, ok := body["password"]; ok {
		t.Fatal("password is exposed in profile response")
	}
}

func TestUserCannotUpdateOtherUser(t *testing.T) {
	app := setupUsersApp(t)
	aliID := register(t, app, "ali")
	otherID := register(t, app, "bobur")
	_, token := loginAs(t, app, "ali", "password123")

	path := fmt.Sprintf("/api/users/%d", otherID)
	if status, _ := doJSON(t, app, http.MethodPatch, path, token, map[string]any{"firstname": "X"}); status != http.StatusForbidden {
		t.Fatalf("expected 403 when updating another user, got %d", status)
	}

	path = fmt.Sprintf("/api/users/%d", aliID)
	if status, _ := doJSON(t, app, http.MethodPatch, path, token, map[string]any{"role": "admin"}); status != http.StatusForbidden {
		t.Fatalf("expected 403 when user changes own role, got %d", status)
	}
	if status, _ := doJSON(t, app, http.MethodPatch, path, token, map[string]any{"firstname": "Ali"}); status != http.StatusOK {
		t.Fatalf("expected 200 when updating own profile, got %d", status)
	}
}

func TestProtectedRouteWithoutToken(t *testing.T) {
	app := setupUsersApp(t)
	if status, _ := doJSON(t, app, http.MethodGet, "/api/users-profile", "", nil); status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", status)
	}
}
