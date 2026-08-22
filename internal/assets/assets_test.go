package assets_test

import (
	"RewriteProject/internal/assets"
	"RewriteProject/internal/config"
	"RewriteProject/internal/localstorage"
	"bytes"
	"io"
	"testing"
	"testing/iotest"
)

// Upload(ctx context.Context, reader io.Reader, upload UploadPolicy) (AssetMetaData, error)
// Delete(ctx context.Context, payload DeletePayload) error
// Read(ctx context.Context, url string) (AssetMetaData, error)
// LinkingAsset(ctx context.Context, assetID []string, parentID string) error
// ReadByParentID(ctx context.Context, parentID []string) ([]AssetMetaData, error)

type UsecaseContract string

const (
	Upload         UsecaseContract = "Upload"
	Delete         UsecaseContract = "Delete"
	Read           UsecaseContract = "Read"
	LinkingAsset   UsecaseContract = "LinkingAsset"
	ReadByParentID UsecaseContract = "ReadByParentID"
)

type UploadTestPayload struct {
	Name        string
	ExpectedErr error
	Usecase     UsecaseContract
	Policy      assets.UploadPolicy
	Reader      io.Reader
}

type ReadTestPayload struct {
	Name        string
	ExpectedErr error
	FileKey     string
}

type LinkingAssetTestPayload struct {
	Name        string
	ExpectedErr error
	AssetID     []string
	ParentID    string
}

type DeleteTestPayload struct {
	Name        string
	ExpectedErr error
	AssetID     []string
}

func NewUC(t *testing.T) assets.UsecaseContract {
	conf := config.StorageConfig{
		StorageRoot:         t.TempDir(),
		StoragePathTemp:     "temp",
		StoragePathUpload:   "upload",
		StoragePathDocument: "document",
		StoragePathAvatar:   "avatar",
		StoragePathVideo:    "video",
	}
	storage, err := localstorage.NewLocalStorage(conf)
	if err != nil {
		t.Fatal("Storage Failed Init")
	}
	assetRepo := assets.NewFakeRepo()
	assetManager := assets.NewUsecase(conf, storage, assetRepo)
	if assetManager == nil {
		t.Fatal("Asset Manager Failed Init")
	}
	return assetManager
}
func CreateUploadUsecase() []UploadTestPayload {
	return []UploadTestPayload{
		{
			Name:        "VALID_CATEGORY_UPLOAD_TEST",
			ExpectedErr: nil,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1024 * 1024 * 10,
				UserID:   "VALID_CATEGORY_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer([]byte("%PDF-1.4")),
		},
		{
			Name:        "VALID_CATEGORY_UPLOAD_TEST",
			ExpectedErr: nil,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1024 * 1024 * 10,
				UserID:   "VALID_CATEGORY_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer([]byte("%PDF-1.4")),
		},
		{
			Name:        "VALID_CATEGORY_UPLOAD_TEST",
			ExpectedErr: nil,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1024 * 1024 * 10,
				UserID:   "VALID_CATEGORY_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer([]byte("%PDF-1.4")),
		},
		{
			Name:        "INVALID_CATEGORY_UPLOAD_TEST",
			ExpectedErr: assets.ErrAssetInvalidMime,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.AssetCategory("invalid"),
				MaxSize:  1024 * 1024 * 10,
				UserID:   "INVALID_CATEGORY_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer([]byte("%PDF-1.4")),
		},
		{
			Name:        "INVALID_SIZE_UPLOAD_TEST",
			ExpectedErr: assets.ErrAssetTooLarge,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1,
				UserID:   "INVALID_SIZE_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer([]byte("%PDF-1.4")),
		},
		{
			Name:        "INVALID_READER_CATEGORY_UPLOAD_TEST",
			ExpectedErr: assets.ErrAssetInvalidMime,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.AssetCategory("invalid"),
				MaxSize:  1024 * 1024 * 10,
				UserID:   "INVALID_READER_MIME_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer([]byte("%PDF-1.4")),
		},
		{
			Name:        "INVALID_READER_ERROR_UPLOAD_TEST",
			ExpectedErr: assets.ErrAssetFailedCreate,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1024 * 1024 * 10,
				UserID:   "INVALID_READER_ERROR_UPLOAD_TEST",
			},
			Reader: iotest.ErrReader(assets.ErrAssetFailedCreate),
		},
		{
			Name:        "INVALID_READER_MIME_UPLOAD_TEST",
			ExpectedErr: assets.ErrAssetInvalidMime,
			Usecase:     Upload,
			Policy: assets.UploadPolicy{
				Category: assets.CategoryDocument,
				MaxSize:  1024 * 1024 * 10,
				UserID:   "INVALID_READER_MIME_UPLOAD_TEST",
			},
			Reader: bytes.NewBuffer([]byte("what is reader?")),
		},
	}
}

func TestUpload(t *testing.T) {
	assetsUC := NewUC(t)
	var validAssets []ReadTestPayload
	var linkingTestPayload []LinkingAssetTestPayload
	var validIDS []string
	var deleteTestPayload []DeleteTestPayload

	validAssets = append(validAssets, ReadTestPayload{
		Name:        "EMPTY_FILE_KEY_READ_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		FileKey:     "",
	}, ReadTestPayload{
		Name:        "NOT_EXISTED_CATEGORY_READ_TEST",
		ExpectedErr: assets.ErrAssetInvalidPath,
		FileKey:     "a1b2c3d4-e5f6-7890-1234-567890abcdef.pdf",
	}, ReadTestPayload{
		Name:        "NESTED_FILE_KEY_READ_TEST",
		ExpectedErr: assets.ErrAssetInvalidPath,
		FileKey:     "avatar/document/a1b2c3d4-e5f6-7890-1234-567890abcdef.pdf",
	})

	testUploads := CreateUploadUsecase()

	for _, tc := range testUploads {

		t.Run(tc.Name, func(t *testing.T) {
			res, err := assetsUC.Upload(t.Context(), tc.Reader, tc.Policy)
			if err == nil {
				validAssets = append(validAssets, ReadTestPayload{
					Name:        "VALID_READ_TEST",
					ExpectedErr: nil,
					FileKey:     res.FileKey,
				})
				validIDS = append(validIDS, res.ID)
			}
			if err != tc.ExpectedErr {
				t.Fatalf("Upload: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
		})
	}
	for _, tc := range validAssets {
		t.Run(tc.Name, func(t *testing.T) {
			f, err := assetsUC.Read(t.Context(), tc.FileKey)
			if err != tc.ExpectedErr {
				t.Fatalf("ReadByID: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
			if f != nil {
				defer f.Close()
			}
		})
	}

	linkingTestPayload = append(linkingTestPayload, LinkingAssetTestPayload{
		Name:        "VALID_LINKING_TEST",
		ExpectedErr: nil,
		AssetID:     validIDS,
		ParentID:    "VALID_LINKING_TEST",
	}, LinkingAssetTestPayload{
		Name:        "INVALID_LINKING_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     []string{"123e4567-e89b-12d3-a456-426614174000"},
		ParentID:    "INVALID_LINKING_TEST",
	}, LinkingAssetTestPayload{
		Name:        "DUP_PARENT_ID_LINKING_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     validIDS,
		ParentID:    "DUP_LINKING_TEST",
	}, LinkingAssetTestPayload{
		Name:        "EMPTY_PARENT_ID_LINKING_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     validIDS,
		ParentID:    "",
	})

	for _, tc := range linkingTestPayload {
		t.Run(tc.Name, func(t *testing.T) {
			err := assetsUC.LinkingAsset(t.Context(), tc.AssetID, tc.ParentID)
			if err != tc.ExpectedErr {
				t.Fatalf("LinkingAsset: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
		})
		t.Run(tc.Name+"_READ_BY_PARENT_ID", func(t *testing.T) {
			_, err := assetsUC.ReadByParentID(t.Context(), []string{tc.ParentID})
			if err != tc.ExpectedErr {
				t.Fatalf("ReadByParentID: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
		})

	}
	for _, tc := range linkingTestPayload {
		t.Run(tc.Name+"_RELEASE_BY_PARENT_ID", func(t *testing.T) {
			err := assetsUC.ReleaseByParentID(t.Context(), []string{tc.ParentID})
			if err != tc.ExpectedErr {
				t.Fatalf("ReleaseByParentID: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
		})
	}

	deleteTestPayload = append(deleteTestPayload, DeleteTestPayload{
		Name:        "VALID_DELETE_TEST",
		ExpectedErr: nil,
		AssetID:     validIDS,
	}, DeleteTestPayload{
		Name:        "INVALID_DELETE_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     []string{"123e4567-e89b-12d3-a456-426614174000"},
	}, DeleteTestPayload{
		Name:        "DUP_PARENT_ID_DELETE_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     validIDS,
	}, DeleteTestPayload{
		Name:        "EMPTY_PARENT_ID_DELETE_TEST",
		ExpectedErr: assets.ErrAssetNotFound,
		AssetID:     validIDS,
	})

	for _, tc := range deleteTestPayload {
		t.Run(tc.Name, func(t *testing.T) {
			err := assetsUC.Delete(t.Context(), assets.DeletePayload{IDs: tc.AssetID})
			if err != tc.ExpectedErr {
				t.Fatalf("Delete: expected =%v, got err=%v", tc.ExpectedErr, err)
			}
		})
	}

}
