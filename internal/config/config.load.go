package config

import (
	"log"
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
		log.Print("use callback value")
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
