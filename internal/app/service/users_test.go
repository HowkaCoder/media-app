package service

import (
	"media-app/internal/app/entity"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	s := NewUserService([]byte("test-secret"))

	hash, err := s.HashPassword("supersecret")
	if err != nil {
		t.Fatal(err)
	}
	if hash == "supersecret" {
		t.Fatal("password is stored as plain text")
	}
	if !s.CheckPassword(hash, "supersecret") {
		t.Fatal("correct password rejected")
	}
	if s.CheckPassword(hash, "wrong") {
		t.Fatal("wrong password accepted")
	}
}

func TestTokenRoundTrip(t *testing.T) {
	s := NewUserService([]byte("test-secret"))
	user := &entity.User{ID: 42, Username: "ali", Role: "admin"}

	token, err := s.GenerateAccessToken(user)
	if err != nil {
		t.Fatal(err)
	}

	claims, err := s.ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if claims.UserID != 42 || claims.Username != "ali" || claims.Role != "admin" {
		t.Fatalf("unexpected claims: %+v", claims)
	}
}

func TestTokenWrongSecret(t *testing.T) {
	token, err := NewUserService([]byte("one")).GenerateAccessToken(&entity.User{ID: 1})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := NewUserService([]byte("two")).ParseToken(token); err == nil {
		t.Fatal("token signed with another secret was accepted")
	}
}
