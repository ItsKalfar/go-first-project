package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string

	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {

	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using default or environment values")
	}

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", ":8080"),
		DBUser:                 getEnv("DB_USER", ""),
		DBPassword:             getEnv("DB_PASSWORD", ""),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", ""), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", ""),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600 * 24 * 7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}