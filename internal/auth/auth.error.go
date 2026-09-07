package auth

import (
	"errors"
	"net/http"
)

// Sentinel errors. Usecase & repository layer harus me-wrap error asli
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid password")
	ErrPasswordShort      = errors.New("password is too short")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidEmail       = errors.New("email is invalid")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrTokenInvalid       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
	ErrRoleInvalid        = errors.New("invalid role")
)

type CodeErr string

// Error Code Contract Frontend
const (
	CodeUserNotFound       CodeErr = "USER_NOT_FOUND"
	CodeInvalidCredentials CodeErr = "INVALID_CREDENTIALS"
	CodePasswordShort      CodeErr = "PASSWORD_SHORT"
	CodeEmailAlreadyExists CodeErr = "EMAIL_ALREADY_EXISTS"
	CodeInvalidEmail       CodeErr = "INVALID_EMAIL"
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
	{ErrInvalidCredentials, CodeInvalidCredentials, http.StatusUnauthorized, "Password salah"},
	{ErrEmailAlreadyExists, CodeEmailAlreadyExists, http.StatusConflict, "Email sudah terdaftar"},
	{ErrUnauthorized, CodeUnauthorized, http.StatusUnauthorized, "Anda belum terautentikasi"},
	{ErrTokenInvalid, CodeTokenInvalid, http.StatusUnauthorized, "Token tidak valid"},
	{ErrTokenExpired, CodeTokenExpired, http.StatusUnauthorized, "Token sudah kedaluwarsa"},
	{ErrInvalidEmail, CodeInvalidEmail, http.StatusBadRequest, "Email tidak valid"},
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
