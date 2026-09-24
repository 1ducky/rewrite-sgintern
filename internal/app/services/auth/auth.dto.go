package auth

import (
	"RewriteProject/internal/auth"
)

type LoginResource struct {
	token   auth.TokenResponse
	Profile SessionProfile
}

type SessionProfile struct {
	UserID   string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Tag      string `json:"tag"`
	Avatar   string `json:"avatar"`
}
