package api

import (
	"RewriteProject/internal/app/assets"
	"RewriteProject/internal/app/config"
	"RewriteProject/internal/app/storage"
)

func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	LocalStorageRepo := storage.NewLocalStorage(appConfig.StorageConfig)
	assetSetvice := assets.NewService(LocalStorageRepo)
	_ = assetSetvice
}
