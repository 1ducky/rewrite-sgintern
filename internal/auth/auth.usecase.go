package auth

import (
	"RewriteProject/internal/utils"
	"context"
	"strings"
	"time"
)

type Service struct {
	CredentialRepository CredentialRepositoryContract
	Token                TokenContract
	SessionRepo          SessionRepositoryContract
}

func NewAuthService(CredentialRepo CredentialRepositoryContract, Token TokenContract, Session SessionRepositoryContract) UsecaseContract {
	return &Service{
		CredentialRepository: CredentialRepo,
		Token:                Token,
		SessionRepo:          Session,
	}
}

func (s *Service) Login(ctx context.Context, payload LoginPayload) (TokenResponse, error) {
	if payload.Email == "" || payload.Password == "" || !utils.IsEmailValid(payload.Email) || len(strings.Split(payload.Password, "")) < PasswordLength {
		return TokenResponse{}, ErrInvalidCredentials
	}
	user, err := s.CredentialRepository.Authentication(ctx, payload.Email)
	if err != nil {
		return TokenResponse{}, ErrInvalidCredentials
	}
	// Compare hash password
	if !utils.CheckHashString(user.Password, payload.Password) {
		return TokenResponse{}, ErrInvalidCredentials
	}

	token, err := s.Token.CreateToken(ctx, TokenEntity{ID: user.ID, Role: user.Role, Version: 0}, 5)
	if err != nil {
		return TokenResponse{}, err
	}
	refreshToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: user.ID, Role: user.Role, Version: 0}, 60*24*30)
	if err != nil {
		return TokenResponse{}, err
	}
	err = s.SessionRepo.Create(ctx, CreateSessionPayload{
		UserID:       user.ID,
		AccessToken:  token,
		RefreshToken: refreshToken,
	})
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil

}

func (s *Service) Register(ctx context.Context, payload RegisterPayload) error {
	// hash password
	switch {
	case len(strings.Split(payload.Password, "")) < PasswordLength:
		return ErrPasswordShort
	}

	password, err := utils.HashString(payload.Password)
	if err != nil {
		return err
	}

	// check email valid
	if !utils.IsEmailValid(payload.Email) {
		return ErrInvalidEmail
	}

	payload.Password = password
	return s.CredentialRepository.Registration(ctx, payload)

}

func (s *Service) Refresh(ctx context.Context, oldrefreshToken string) (TokenResponse, error) {
	user, err := s.Token.VerifyToken(ctx, oldrefreshToken)
	if err != nil {
		return TokenResponse{}, err
	}
	// Revoke Version
	session, err := s.SessionRepo.GetByRefreshToken(ctx, oldrefreshToken)
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
	err = s.SessionRepo.Rotate(ctx, UpdateSessionPayload{
		UserID:          user.ID,
		OldrefreshToken: oldrefreshToken,
		AccessToken:     newToken,
		RefreshToken:    newRefreshToken,
		Version:         nextVersion,
	})
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{AccessToken: newToken, RefreshToken: newRefreshToken}, nil
}

func (s *Service) Logout(ctx context.Context, oldrefreshToken string) error {
	// Revoke Version
	user, err := s.Token.VerifyToken(ctx, oldrefreshToken)
	if err != nil {
		return err
	}
	err = s.SessionRepo.Delete(ctx, DeleteSessionPayload{
		UserID:       user.ID,
		RefreshToken: oldrefreshToken,
	})

	if err != nil {
		return err
	}
	return nil
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
