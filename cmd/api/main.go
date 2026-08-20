package main

import (
	"RewriteProject/internal/assets"
	"RewriteProject/internal/config"
	storage "RewriteProject/internal/localstorage"
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

	file, err := os.Open("./0428.gif")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	// 2. Ensure the file is closed to prevent resource leaks
	defer file.Close()

	res, err := assetUsecase.Upload(context.Background(), file, assets.UploadPolicy{Category: assets.CategoryProfile, UserID: "asdad", MaxSize: 10 << 20}) //10 mb
	log.Print(res, err)
}
