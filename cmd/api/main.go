package main

import (
	"RewriteProject/internal/assets"
	"RewriteProject/internal/config"
	"RewriteProject/internal/storage"
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Println("Error loading config:", err)
		panic(err)
	}

	LocalStorageRepo, err := storage.NewLocalStorage(appConfig.StorageConfig)
	if err != nil {
		log.Println("Error loading storage:", err)
		panic(err)
	}
	// init repo asset
	assetRepo := assets.NewRepository()

	assetUsecase := assets.NewUsecase(appConfig.StorageConfig, LocalStorageRepo, assetRepo)
	_ = assetUsecase

	file, err := os.Open("./images1.jpg")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	// 2. Ensure the file is closed to prevent resource leaks
	defer file.Close()

	res, err := assetUsecase.Upload(context.Background(), file, assets.UploadPolicy{Category: assets.CategoryProfile, UserID: "asdad", MaxSize: 1 << 20})
	log.Print(res, err)
}
