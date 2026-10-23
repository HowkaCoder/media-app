package config

import (
	"errors"
	"os"
)

type Config struct {
	Port        string
	DBPath      string
	ImagesDir   string
	PublicURL   string
	JWTSecret   string
	RedisAddr   string
	SMSEmail    string
	SMSPassword string
}

func Load() (Config, error) {
	cfg := Config{
		Port:        getenv("PORT", "8082"),
		DBPath:      getenv("DB_PATH", "database/database.db"),
		ImagesDir:   getenv("IMAGES_DIR", "images"),
		PublicURL:   getenv("PUBLIC_URL", "http://localhost:8082"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		RedisAddr:   getenv("REDIS_ADDR", "localhost:6379"),
		SMSEmail:    os.Getenv("ESKIZ_EMAIL"),
		SMSPassword: os.Getenv("ESKIZ_PASSWORD"),
	}

	if cfg.JWTSecret == "" {
		return cfg, errors.New("JWT_SECRET is not set")
	}

	return cfg, nil
}

func (c Config) SMSEnabled() bool {
	return c.SMSEmail != "" && c.SMSPassword != ""
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
