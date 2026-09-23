package auth

import (
	"RewriteProject/internal/auth"
	"RewriteProject/internal/profile"
)

type LoginResource struct {
	token   auth.TokenResponse
	Profile profile.Profile
}
