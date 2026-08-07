package assets

import (
	"context"
	"io"
)

type RepositoryContract interface {
	Write(ctx context.Context, reader io.Reader, path string) (StoreResult, error)
	Delete(ctx context.Context, path string) error
	Read(ctx context.Context, path string) (io.ReadCloser, error)
}
type StoreResult struct {
	Path string
	Size int64
}

type UsecaseContract interface {
	Upload(ctx context.Context, reader io.Reader) error
	Update(ctx context.Context, id string, reader io.Reader) error
	Delete(ctx context.Context, id string) error
	Read(ctx context.Context, url string) error
}
