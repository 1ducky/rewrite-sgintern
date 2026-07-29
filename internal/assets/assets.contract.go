package assets

import (
	"context"
	"io"
)

type RepositoryContract interface {
	Write(ctx context.Context, reader io.Reader, path string, fileName string) (StoreResult, error)
	Delete(ctx context.Context, path string) (string, error)
	Read(ctx context.Context, path string) (io.ReadCloser, error)
}
type StoreResult struct {
	Filename string
	Path     string
	Size     int64
}
