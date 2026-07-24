package storage

import (
	"RewriteProject/internal/app/assets"
	"RewriteProject/internal/app/config"
	"context"
	"io"
)

type LocalStorage struct {
	conf config.StorageConfig
}

func NewLocalStorage(conf config.StorageConfig) assets.ReposioturyContract {
	return &LocalStorage{conf: conf}
}

func (l *LocalStorage) Upload(ctx context.Context, reader io.Reader, path string) (assets.StoreResult, error) {
	return assets.StoreResult{}, nil
}

func (l *LocalStorage) Rewrite(ctx context.Context, reader io.Reader, path string) (assets.StoreResult, error) {
	return assets.StoreResult{}, nil
}

func (l *LocalStorage) Delete(ctx context.Context, path string) (string, error) {
	return "", nil
}

func (l *LocalStorage) GetFile(ctx context.Context, path string) (assets.StoreResult, error) {
	return assets.StoreResult{}, nil
}
