package auth

import "context"

type AuthContract interface {
	Login(ctx context.Context, payload CredentialPayload) (TokenResponse, error)
	Refresh(ctx context.Context, oldrefreshToken string) (TokenResponse, error)
	Logout(ctx context.Context, refreshToken string) error
	Verify(ctx context.Context, token string) (TokenEntity, error)
}

type TokenContract interface {
	CreateToken(ctx context.Context, entity TokenEntity, minute TimeInt) (string, error)
	VerifyToken(ctx context.Context, token string) (TokenEntity, error)
}

type SessionContract interface {
	Create(ctx context.Context, payload SessionEntity) error
	Rotate(ctx context.Context, payload SessionRotatePayload) error
	Delete(ctx context.Context, payload SessionDeletePayload) error
	GetByID(ctx context.Context, sessionID string) (SessionEntity, error)
}

type CredentialContract interface {
	FindByEmail(ctx context.Context, email string) (CredentialEntity, error)
	Login(ctx context.Context, payload CredentialPayload) (CredentialEntity, error)
}
