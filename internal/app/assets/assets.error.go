package assets

import (
	"errors"
	"net/http"
)

// Sentinel errors. Usecase & repository layer harus me-wrap error asli
var (
	ErrAssetNotFound      = errors.New("asset not found")
	ErrAssetAlreadyExists = errors.New("asset already exists")
	ErrInvalidAssetInput  = errors.New("invalid asset input")
	ErrAssetTooLarge      = errors.New("asset size exceeds limit")
	ErrUnsupportedFormat  = errors.New("unsupported asset format")
	ErrStorageUnavailable = errors.New("storage backend unavailable")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrUnauthorized       = errors.New("unauthorized")
)

type CodeErr string

// Error Code Contract Frontend
const (
	CodeAssetNotFound      CodeErr = "ASSET_NOT_FOUND"
	CodeAssetAlreadyExists CodeErr = "ASSET_ALREADY_EXISTS"
	CodeInvalidInput       CodeErr = "INVALID_INPUT"
	CodeAssetTooLarge      CodeErr = "ASSET_TOO_LARGE"
	CodeUnsupportedFormat  CodeErr = "UNSUPPORTED_FORMAT"
	CodeStorageUnavailable CodeErr = "STORAGE_UNAVAILABLE"
	CodePermissionDenied   CodeErr = "PERMISSION_DENIED"
	CodeUnauthorized       CodeErr = "UNAUTHORIZED"
	CodeInternal           CodeErr = "INTERNAL_ERROR"
)

type entry struct {
	err        error
	code       CodeErr
	statusCode int
	message    string
}

var MapErr = []entry{
	{ErrAssetNotFound, CodeAssetNotFound, http.StatusNotFound, "Asset tidak ditemukan"},
	{ErrAssetAlreadyExists, CodeAssetAlreadyExists, http.StatusConflict, "Asset sudah ada"},
	{ErrInvalidAssetInput, CodeInvalidInput, http.StatusBadRequest, "Input tidak valid"},
	{ErrAssetTooLarge, CodeAssetTooLarge, http.StatusRequestEntityTooLarge, "Ukuran asset melebihi batas"},
	{ErrUnsupportedFormat, CodeUnsupportedFormat, http.StatusUnsupportedMediaType, "Format asset tidak didukung"},
	{ErrStorageUnavailable, CodeStorageUnavailable, http.StatusServiceUnavailable, "Layanan penyimpanan sedang tidak tersedia"},
	{ErrPermissionDenied, CodePermissionDenied, http.StatusForbidden, "Anda tidak memiliki akses"},
	{ErrUnauthorized, CodeUnauthorized, http.StatusUnauthorized, "Anda belum terautentikasi"},
}

type ErrorAssets struct {
	Code    CodeErr
	Message string
	Status  int
}

func TranslateErr(err error) (ErrorAssets, bool) {
	if err == nil {
		return ErrorAssets{},
			false
	}
	for _, e := range MapErr {
		if e.err == err {
			return ErrorAssets{Code: e.code, Message: e.message, Status: e.statusCode},
				true
		}
	}
	return ErrorAssets{Code: CodeInternal,
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError},
		true
}
