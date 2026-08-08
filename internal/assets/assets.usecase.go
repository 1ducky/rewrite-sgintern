package assets

import (
	"RewriteProject/internal/auth"
	"RewriteProject/internal/config"
	"RewriteProject/internal/reader"
	"context"
	"io"

	"github.com/google/uuid"
)

type Usecase struct {
	StorageRepo RepositoryContract
	AssetRepo   AssetRepository
	conf        config.StorageConfig
}

func NewUsecase(config config.StorageConfig, StorageRepo RepositoryContract, AssetRepo AssetRepository) UsecaseContract {
	return &Usecase{StorageRepo: StorageRepo, AssetRepo: AssetRepo, conf: config}
}

func (u *Usecase) Upload(ctx context.Context, r io.Reader, upload UploadPolicy) error {
	user, ok := auth.GetContext(ctx)
	if !ok {
		return auth.ErrUnauthorized
	}
	ext, newR, err := reader.DetectMime(r)
	if err != nil {
		return err
	}
	limitedR := u.limitReader(newR, upload.MaxSize)
	policyCategory := PolicyAsset[upload.Category]
	mm, ok := isAllowedExt(ext)
	if !ok || !isAllowedMimeByPolicy(policyCategory, mm.Mime) {
		return ErrAssetInvalidMime
	}
	name := user.ID + "_" + uuid.NewString() + string(mm.Ext)
	fullpath, err := u.resolvePath(upload.Category, name)
	md, err := u.AssetRepo.Record(ctx, RecordPayload{Status: AssetStatusPending, Filename: name, Mime: mm.Mime, Url: fullpath, AuthorID: user.ID})
	if err != nil {
		return ErrAssetFailedCreate
	}
	res, err := u.StorageRepo.Write(ctx, limitedR, fullpath)
	if err != nil {
		return ErrAssetFailedCreate
	}
	if res.Size > upload.MaxSize {
		return ErrAssetTooLarge
	}
	_, err = u.AssetRepo.Update(ctx, UpdatePayload{
		Id:        md.ID,
		Size:      res.Size,
		Status:    AssetStatusActive,
		OldStatus: md.Status,
		AuthorID:  user.ID,
	})
	if err != nil {
		return ErrAssetFailedCreate
	}

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
