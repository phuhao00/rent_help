package config

import (
	"os"
)

type Config struct {
	Port      string
	MongoURI  string
	JWTSecret string
	DBName    string
}

func Load() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		MongoURI:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
		JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
		DBName:    getEnv("DB_NAME", "rent_help"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
