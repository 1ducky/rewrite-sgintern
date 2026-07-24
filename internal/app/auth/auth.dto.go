package auth

import "context"

type AuthContract interface {
	Login(ctx context.Context) (string, error)
	Register(ctx context.Context) (string, error)
	Refresh(ctx context.Context) (string, error)
	Logout(ctx context.Context) (string, error)
	Verify(ctx context.Context) (string, error)
}
