package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	OpenAIAPIKey  string
}

func Load() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8081"),
		OpenAIAPIKey:  getEnv("OPENAI_API_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
