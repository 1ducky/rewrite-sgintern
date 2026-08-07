package auth

import (
	"errors"
	"net/http"
)

// Sentinel errors. Usecase & repository layer harus me-wrap error asli
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredential  = errors.New("invalid password")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrTokenInvalid       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
	ErrRoleInvalid        = errors.New("invalid role")
)

type CodeErr string

// Error Code Contract Frontend
const (
	CodeUserNotFound       CodeErr = "USER_NOT_FOUND"
	CodeInvalidCredential  CodeErr = "INVALID_CREDENTIAL"
	CodeEmailAlreadyExists CodeErr = "EMAIL_ALREADY_EXISTS"
	CodeUnauthorized       CodeErr = "UNAUTHORIZED"
	CodeTokenInvalid       CodeErr = "TOKEN_INVALID"
	CodeTokenExpired       CodeErr = "TOKEN_EXPIRED"
	CodeRoleInvalid        CodeErr = "ROLE_INVALID"
	CodeInternal           CodeErr = "INTERNAL_ERROR"
)

type entry struct {
	err        error
	code       CodeErr
	statusCode int
	message    string
}

var MapErr = []entry{
	{ErrUserNotFound, CodeUserNotFound, http.StatusNotFound, "Pengguna tidak ditemukan"},
	{ErrInvalidCredential, CodeInvalidCredential, http.StatusUnauthorized, "Credential salah"},
	{ErrEmailAlreadyExists, CodeEmailAlreadyExists, http.StatusConflict, "Email sudah terdaftar"},
	{ErrUnauthorized, CodeUnauthorized, http.StatusUnauthorized, "Anda belum terautentikasi"},
	{ErrTokenInvalid, CodeTokenInvalid, http.StatusUnauthorized, "Token tidak valid"},
	{ErrTokenExpired, CodeTokenExpired, http.StatusUnauthorized, "Token sudah kedaluwarsa"},
}

type ErrorAuth struct {
	Code    CodeErr
	Status  int
	Message string
}

func TranslateErr(err error) (ErrorAuth, bool) {
	if err == nil {
		return ErrorAuth{},
			false
	}
	for _, e := range MapErr {
		if errors.Is(err, e.err) {
			return ErrorAuth{e.code, e.statusCode, e.message},
				true
		}
	}
	return ErrorAuth{Code: CodeInternal,
			Message: "Internal Server Error",
			Status:  http.StatusInternalServerError},
		true
}
