package main

import (
	"RewriteProject/internal/assets"
	"RewriteProject/internal/config"
	"RewriteProject/internal/storage"
)

func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	LocalStorageRepo, err := storage.NewLocalStorage(appConfig.StorageConfig)
	if err != nil {
		panic(err)
	}
	// init repo asset
	assetRepo := assets.NewRepository()

	assetUsecase := assets.NewUsecase(appConfig.StorageConfig, LocalStorageRepo, assetRepo)
	_ = assetUsecase
}
