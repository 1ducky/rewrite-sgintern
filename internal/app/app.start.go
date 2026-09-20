package app

import (
	"RewriteProject/internal/assets"
	"RewriteProject/internal/config"
	"RewriteProject/internal/db"
	"RewriteProject/internal/localstorage"
)

type repos struct {
	assets  assets.AssetRepository
	storage assets.RepositoryContract
}

type ucs struct {
	assets assets.UsecaseContract
}

func createRepo(cfg config.AppConfig, db db.DBTX) *repos {
	storage, err := localstorage.NewLocalStorage(cfg.StorageConfig)
	if err != nil {
		panic(err)
	}
	return &repos{
		assets:  assets.NewMysqlRepo(db),
		storage: storage,
	}
}

func createUcs(cfg config.AppConfig, r *repos) *ucs {
	return &ucs{
		assets: assets.NewUsecase(cfg.StorageConfig, r.storage, r.assets),
	}
}

// func StartApp() {
// 	cfg,_ := config.LoadConfig()
// 	ctx := context.Background()
// 	db, err := db.NewMySQLDatabase(&cfg.DBConfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	repos := createRepo(cfg,db)
// 	ucs := createUcs(cfg,repos)
// 	handler :=

// 	// transport.Start()
// }
