package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string `yaml:"telegram_bot_token"`
	DebugMode        bool   `yaml:"bot_debug"`
}

type yamlConfig struct {
	TelegramBotToken string `yaml:"telegram_bot_token"`
	BotDebug         bool   `yaml:"bot_debug"`
}

func Load() *Config {
	cfg := &Config{
		TelegramBotToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
		DebugMode:        getEnvBool("BOT_DEBUG", false),
	}

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("Файл config.yaml не найден, используются переменные окружения: %v", err)
		return cfg
	}

	var yamlCfg yamlConfig
	if err := yaml.Unmarshal(data, &yamlCfg); err != nil {
		log.Printf("Ошибка при чтении config.yaml, используются переменные окружения: %v", err)
		return cfg
	}

	if yamlCfg.TelegramBotToken != "" {
		cfg.TelegramBotToken = yamlCfg.TelegramBotToken
	}
	cfg.DebugMode = yamlCfg.BotDebug

	return cfg
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
