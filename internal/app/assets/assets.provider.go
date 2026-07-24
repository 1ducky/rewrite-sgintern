package assets

import (
	"context"
	"io"
)

type ReposioturyContract interface {
	Upload(ctx context.Context, reader io.Reader, path string) (StoreResult, error)
	Rewrite(ctx context.Context, reader io.Reader, path string) (StoreResult, error)
	Delete(ctx context.Context, path string) (string, error)
	GetFile(ctx context.Context, path string) (StoreResult, error)
}
type StoreResult struct {
	Filename string
	Path     string
}
