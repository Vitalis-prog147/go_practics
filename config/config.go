package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string
	DebugMode        bool
}

func Load() *Config {
	if err := godotenv.Load("config.yaml"); err != nil {
		log.Printf("Файл config.yaml не найден, используются переменные окружения системы: %v", err)
	}

	return &Config{
		TelegramBotToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
		DebugMode:        getEnvBool("BOT_DEBUG", false),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "true" || value == "1" || value == "yes"
	}
	return defaultValue
}
