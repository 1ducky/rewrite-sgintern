package assets

import (
	"context"
	"io"
)

type Usecase struct {
	StorageRepo RepositoryContract
}

func NewService(StorageRepo RepositoryContract) UsecaseContract {
	return &Usecase{StorageRepo: StorageRepo}
}

func (u *Usecase) Upload(ctx context.Context, reader io.Reader) error {
	return nil
}
func (u *Usecase) Update(ctx context.Context, id string, reader io.Reader) error {
	return nil
}
func (u *Usecase) Delete(ctx context.Context, id string) error {
	return nil
}
func (u *Usecase) Read(ctx context.Context, url string) error {
	return nil
}
