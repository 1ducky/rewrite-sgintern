package auth

import (
	"RewriteProject/internal/utils"
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
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
	credentials, err := s.CredentialRepository.Authentication(ctx, payload.Email)
	if err != nil {
		return TokenResponse{}, ErrInvalidCredentials
	}
	// Compare hash password
	if !utils.CheckHashString(credentials.Password, payload.Password) {
		return TokenResponse{}, ErrInvalidCredentials
	}

	accessTokenRevokedAt := time.Now().Add(time.Minute * time.Duration(AccessTokenDuration))
	refreshTokenRevokedAt := time.Now().Add(time.Minute * time.Duration(RefreshTokenDuration))
	SessionID := "SESSION_" + credentials.UserID + "_" + uuid.NewString()

	token, err := s.Token.CreateToken(ctx, TokenEntity{ID: SessionID, UserID: credentials.UserID, Role: credentials.Role, Version: 0, RevokeAt: accessTokenRevokedAt})
	if err != nil {
		return TokenResponse{}, err
	}
	refreshToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: SessionID, UserID: credentials.UserID, Role: credentials.Role, Version: 0, RevokeAt: refreshTokenRevokedAt})
	if err != nil {
		return TokenResponse{}, err
	}
	err = s.SessionRepo.Create(ctx, CreateSessionPayload{
		SessionID:    SessionID,
		AccessToken:  token,
		RefreshToken: refreshToken,
		RevokeAt:     accessTokenRevokedAt,
		UserID:       credentials.UserID,
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
	payload.CredentialID = uuid.NewString()
	return s.CredentialRepository.Registration(ctx, payload)

}

func (s *Service) Refresh(ctx context.Context, oldrefreshToken string) (TokenResponse, error) {
	token, err := s.Token.VerifyToken(ctx, oldrefreshToken)
	if err != nil {
		return TokenResponse{}, err
	}
	// Revoke Version
	session, err := s.SessionRepo.GetByRefreshToken(ctx, oldrefreshToken)
	if err != nil {
		return TokenResponse{}, err
	}
	if session.Version != token.Version {
		return TokenResponse{}, ErrTokenExpired
	}

	if !session.RevokeAt.IsZero() || session.RevokeAt.Before(time.Now()) {
		return TokenResponse{}, ErrTokenExpired
	}
	nextVersion := session.Version + 1
	accessTokenRevokedAt := time.Now().Add(time.Minute * time.Duration(AccessTokenDuration))
	refreshTokenRevokedAt := time.Now().Add(time.Minute * time.Duration(RefreshTokenDuration))
	role, err := s.CredentialRepository.GetRoleByUserId(ctx, session.UserID)
	if err != nil {
		return TokenResponse{}, err
	}

	newToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: session.ID, Role: role, Version: nextVersion, UserID: session.ID, RevokeAt: accessTokenRevokedAt})
	if err != nil {
		return TokenResponse{}, err
	}
	newRefreshToken, err := s.Token.CreateToken(ctx, TokenEntity{ID: session.ID, Role: role, Version: nextVersion, UserID: session.ID, RevokeAt: refreshTokenRevokedAt})
	if err != nil {
		return TokenResponse{}, err
	}
	err = s.SessionRepo.Rotate(ctx, UpdateSessionPayload{
		UserID:          session.UserID,
		OldrefreshToken: oldrefreshToken,
		AccessToken:     newToken,
		RefreshToken:    newRefreshToken,
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
