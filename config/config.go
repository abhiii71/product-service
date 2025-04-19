package config

import (
	"log"
	"os"
)

type Config struct {
	DBUrl     string
	RedisHost string
	RedisPort string
	RedisTTL  string
}

func LoadConfig() *Config {
	return &Config{
		DBUrl:     getEnv("DB_URL", "postgers://postgres:postgres@db:5432/productdb?sslmode=disable"),
		RedisHost: getEnv("REDIS_HOST", "redis"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
		RedisTTL:  getEnv("REDIS_TTL", "60"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("Using default  for %s: %s", key, fallback)
	return fallback
}
