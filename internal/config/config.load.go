package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	StorageConfig StorageConfig
	AuthConfig    AuthConfig
}

func LoadConfig() (AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return AppConfig{}, err
	}
	StorageConf := LoadStorageConfig()
	authConf := LoadAuthConfig()
	return AppConfig{StorageConfig: StorageConf, AuthConfig: authConf}, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
