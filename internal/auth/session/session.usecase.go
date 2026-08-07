package session

import (
	"RewriteProject/internal/auth"
	"context"
)

type Service struct {
}

func NewUsecase() auth.SessionUsecase {
	return &Service{}
}

func (s *Service) Create(ctx context.Context, payload auth.CreateSessionPayload) (auth.SessionEntity, error) {
	return auth.SessionEntity{}, nil
}
func (s *Service) Rotate(ctx context.Context, payload auth.UpdateSessionPayload) (auth.SessionEntity, error) {
	return auth.SessionEntity{}, nil
}

func (s *Service) Delete(ctx context.Context, payload auth.DeleteSessionPayload) error {
	return nil
}

func (s *Service) GetByAccessToken(ctx context.Context, userID, accessToken string) (auth.SessionEntity, error) {
	return auth.SessionEntity{}, nil
}

func (s *Service) GetByRefreshToken(ctx context.Context, userID, refreshToken string) (auth.SessionEntity, error) {
	return auth.SessionEntity{}, nil
}
