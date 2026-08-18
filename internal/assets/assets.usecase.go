package assets

import (
	"RewriteProject/internal/app/pipeline"
	"RewriteProject/internal/config"
	"RewriteProject/internal/reader"
	"RewriteProject/internal/utils"
	"context"
	"errors"
	"io"
	"log"

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

func (u *Usecase) Upload(ctx context.Context, r io.Reader, upload UploadPolicy) (AssetMetaData, error) {
	log.Print("detect mime")
	mime, newR, err := reader.DetectMime(r)
	if err != nil {
		return AssetMetaData{}, err
	}
	log.Print("limit reader")
	limitedR := u.limitReader(newR, upload.MaxSize)
	log.Print("allowed type buffer")
	policyCategory := PolicyAsset[upload.Category]
	log.Print("policyCategory", policyCategory)
	typeBuffer, ok := isAllowedExt(Mime(mime))
	log.Print("typeBuffer", typeBuffer)
	if !ok || !isAllowedMimeByPolicy(policyCategory, typeBuffer.Mime) {
		return AssetMetaData{}, ErrAssetInvalidMime
	}
	log.Print("generate name")
	name := upload.UserID + "_" + uuid.NewString() + string(typeBuffer.Ext)
	log.Print("fullpath", name)
	log.Print("resolve path")
	fullpath, err := u.resolvePath(upload.Category, name)
	if err != nil {
		return AssetMetaData{}, err
	}

	log.Print("record metadata")
	metadata, err := u.AssetRepo.Record(ctx, RecordPayload{Status: AssetStatusPending, FileKey: fullpath, Filename: name, Mime: typeBuffer.Mime, Path: fullpath, AuthorID: upload.UserID, Category: upload.Category})
	if err != nil {
		return AssetMetaData{}, err
	}
	log.Print("storage write")
	res, err := u.StorageRepo.Write(ctx, limitedR, fullpath)
	if err != nil {
		return AssetMetaData{}, err
	}
	log.Print("check size")
	if res.Size > upload.MaxSize {
		return AssetMetaData{}, ErrAssetTooLarge
	}
	log.Print("update metadata")
	metadata, err = u.AssetRepo.Update(ctx, UpdatePayload{
		Id:        metadata.ID,
		Size:      res.Size,
		Status:    AssetStatusActive,
		OldStatus: metadata.Status,
		AuthorID:  upload.UserID,
	})
	if err != nil {
		return AssetMetaData{}, ErrAssetFailedCreate
	}

	return metadata, nil
}

func (u *Usecase) Delete(ctx context.Context, payload DeletePayload) error {
	meta, err := u.AssetRepo.GetByIDs(ctx, payload.IDs)
	if err != nil {
		return err
	}
	if len(meta) == 0 {
		return ErrAssetNotFound
	}
	workerCount := pipeline.CalculateWorkerCount(len(meta), pipeline.WorkerPoolConfig{
		MaxWorkers:    8,
		MinWorkers:    2,
		JobsPerWorker: 5,
	})
	pool := pipeline.NewWorkerPool[AssetMetaData, DeleteReaport](workerCount)
	jobs := pipeline.ProduceJob(ctx, meta, workerCount*2)
	reports := pool.Run(ctx, jobs, func(ctx context.Context, amd AssetMetaData) DeleteReaport {

		if amd.Path == "" {
			return DeleteReaport{id: amd.ID, err: ErrAssetInvalidPath}
		}
		delErr := u.StorageRepo.Delete(ctx, amd.Path)
		if delErr != nil {
			return DeleteReaport{id: amd.ID, err: delErr}
		}
		return DeleteReaport{id: amd.ID, err: nil}
	})
	res := pipeline.Collect(ctx, reports, len(meta))

	ids := utils.MapField(res, func(dr DeleteReaport) (string, bool) {
		if dr.err != nil && !errors.Is(dr.err, ErrAssetNotFound) {
			return "", false
		}
		return dr.id, true
	}) //get deleted ids successfull
	err = u.AssetRepo.MarkAsDeleted(ctx, DeletePayload{IDs: ids}) //mark as deleted
	if err != nil {
		return err
	}
	return nil
}
func (u *Usecase) Read(ctx context.Context, url string) (AssetMetaData, error) {
	return AssetMetaData{}, nil
}

func (u *Usecase) LinkingAsset(ctx context.Context, assetIDs []string, parentID string) error {
	err := u.AssetRepo.LinkingAsset(ctx, assetIDs, parentID)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) ReadByParentID(ctx context.Context, parentID []string) ([]AssetMetaData, error) {
	data, err := u.AssetRepo.ReadByParentID(ctx, parentID)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, ErrAssetNotFound
	}
	return data, nil
}

func (u *Usecase) ReleaseByParentID(ctx context.Context, parentID []string) ([]AssetMetaData, error) {
	data, err := u.AssetRepo.ReadByParentID(ctx, parentID)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, ErrAssetNotFound
	}
	return data, nil
}
