package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	StorageConfig StorageConfig
}

func LoadConfig() (AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return AppConfig{}, err
	}
	StorageConf := LoadStorageConfig()
	return AppConfig{StorageConfig: StorageConf}, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
