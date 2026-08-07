package auth

import (
	"context"
	"time"
)

type Service struct {
	Repository RepositoryContract
	Token      TokenContract
	Session    SessionUsecase
}

func NewAuthService(Repo RepositoryContract, Token TokenContract, Session SessionUsecase) UsecaseContract {
	return &Service{
		Repository: Repo,
		Token:      Token,
	}
}

func (s *Service) Login(ctx context.Context, payload LoginPayload) (TokenResponse, error) {
	user, err := s.Repository.Login(ctx, payload.Email)
	if err != nil {
		return TokenResponse{}, err
	}
	// Compare hash password
	if user.Password != payload.Password {
		return TokenResponse{}, ErrInvalidPassword
	}
	
	token, err := s.Token.CreateToken(ctx, TokenEntity{ID: user.ID, Role: user.Role, Version: 0}, 5)
	if err != nil {
		return TokenResponse{}, err
	}
	refreshToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: user.ID, Role: user.Role, Version: 0}, 60*24*30)
	if err != nil {
		return TokenResponse{}, err
	}
	res, err := s.Session.Create(ctx, CreateSessionPayload{
		UserID:       user.ID,
		AccessToken:  token,
		RefreshToken: refreshToken,
	})
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil

}

func (s *Service) Refresh(ctx context.Context, oldrefreshToken string) (TokenResponse, error) {
	user, err := s.Token.VerifyToken(ctx, oldrefreshToken)
	if err != nil {
		return TokenResponse{}, err
	}
	// Revoke Version
	session, err := s.Session.GetByRefreshToken(ctx, user.ID, oldrefreshToken)
	if err != nil {
		return TokenResponse{}, err
	}
	if session.Version != user.Version {
		return TokenResponse{}, ErrTokenExpired
	}

	if !session.RevokeAt.IsZero() || session.RevokeAt.Before(time.Now()) {
		return TokenResponse{}, ErrTokenExpired
	}
	nextVersion := session.Version + 1

	newToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: user.ID, Role: user.Role, Version: nextVersion}, 5)
	if err != nil {
		return TokenResponse{}, err
	}
	newRefreshToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: user.ID, Role: user.Role, Version: nextVersion}, 60*24*30)
	if err != nil {
		return TokenResponse{}, err
	}
	res, err := s.Session.Rotate(ctx, UpdateSessionPayload{
		UserID:          user.ID,
		OldrefreshToken: oldrefreshToken,
		AccessToken:     newToken,
		RefreshToken:    newRefreshToken,
		Version:         nextVersion,
	})
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{AccessToken: res.AccessToken, RefreshToken: res.RefreshToken}, nil
}

func (s *Service) Logout(ctx context.Context, refreshToken string) (string, error) {
	// Revoke Version
	return "", nil
}

func (s *Service) Verify(ctx context.Context, token string) (TokenEntity, error) {
	user, err := s.Token.VerifyToken(ctx, token)
	if err != nil {
		return TokenEntity{}, err
	}
	// Call repository to check version
	version := 0

	if user.Version != version {
		return TokenEntity{}, ErrTokenExpired
	}
	return user, nil
}
