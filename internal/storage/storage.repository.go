package storage

import (
	"RewriteProject/internal/assets"
	"RewriteProject/internal/config"
	"context"
	"io"
	"os"
	"path"
)

type LocalStorage struct {
	conf config.StorageConfig
}

func NewLocalStorage(conf config.StorageConfig) (assets.RepositoryContract, error) {
	// Initial Folder
	dirs := []string{
		conf.StorageRoot,
		path.Join(conf.StorageRoot, conf.StoragePathTemp),
		path.Join(conf.StorageRoot, conf.StoragePathUpload),
		path.Join(conf.StorageRoot, conf.StoragePathDocument),
		path.Join(conf.StorageRoot, conf.StoragePathAvatar),
	}
	_AttemtDir := 10

	var attempDir []string

	for _AttemtDir > 0 {
		for _, dir := range dirs {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				attempDir = append(attempDir, dir)
			}
		}
		dirs = attempDir
		attempDir = []string{}
		_AttemtDir--
	}

	if len(attempDir) > 0 {
		return nil, assets.ErrStorageUnavailable
	}

	return &LocalStorage{conf: conf}, nil
}

func (l *LocalStorage) Write(ctx context.Context, reader io.Reader, path string) (assets.StoreResult, error) {
	dst, err := os.Create(path)
	if err != nil {
		return assets.StoreResult{}, assets.ErrAssetFailedCreate
	}
	defer dst.Close()

	size, err := io.Copy(dst, reader)
	if err != nil {
		return assets.StoreResult{}, assets.ErrAssetFailedCreate
	}

	return assets.StoreResult{Path: path, Size: size}, nil
}

func (l *LocalStorage) Delete(ctx context.Context, path string) error {
	err := os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return assets.ErrAssetNotFound
		}
		return err
	}
	return nil
}

func (l *LocalStorage) Read(ctx context.Context, path string) (io.ReadCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, assets.ErrAssetNotFound
		}
		return nil, err
	}
	return f, nil
}
