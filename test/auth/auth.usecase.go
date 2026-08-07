package auth

import (
	"RewriteProject/internal/utils"
	"context"
	"time"

	"github.com/google/uuid"
)

type Usecase struct {
	Token      TokenContract
	Session    SessionContract
	Credential CredentialContract
}

func NewUsecase(token TokenContract, session SessionContract, credential CredentialContract) AuthContract {
	return &Usecase{
		Token:      token,
		Session:    session,
		Credential: credential,
	}
}

func (u *Usecase) Login(ctx context.Context, payload CredentialPayload) (TokenResponse, error) {
	intialUUID := uuid.NewString()
	user, err := u.Credential.Login(ctx, payload)
	if err != nil {
		return TokenResponse{}, err
	}
	if !utils.CheckHashString(payload.Password, user.Password) {
		return TokenResponse{}, ErrInvalidCredential
	}
	accessToken, err := u.Token.CreateToken(ctx, TokenEntity{UserID: user.UserID, SessionID: intialUUID, Version: InitialSessionVersion}, AccessTokenExpired)
	if err != nil {
		return TokenResponse{}, err
	}
	refreshToken, err := u.Token.CreateToken(ctx, TokenEntity{UserID: user.UserID, SessionID: intialUUID, Version: InitialSessionVersion}, RefreshTokenExpired)
	if err != nil {
		return TokenResponse{}, err
	}
	err = u.Session.Create(ctx, SessionEntity{ID: intialUUID, RefreshToken: refreshToken, Version: InitialSessionVersion, UserID: user.UserID})
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil

}
func (u *Usecase) Refresh(ctx context.Context, oldrefreshToken string) (TokenResponse, error) {
	token, err := u.Token.VerifyToken(ctx, oldrefreshToken)
	if err != nil {
		return TokenResponse{}, err
	}
	currentSession, err := u.Session.GetByID(ctx, token.SessionID)
	if err != nil {
		return TokenResponse{}, err
	}
	if token.Version != currentSession.Version {
		return TokenResponse{}, ErrTokenExpired
	}
	if !currentSession.RevokedAt.IsZero() || currentSession.RevokedAt.Before(time.Now()) {
		return TokenResponse{}, ErrTokenExpired
	}
	// Rotate New Version
	newVersion := currentSession.Version + 1
	accessToken, err := u.Token.CreateToken(ctx, TokenEntity{UserID: token.UserID, SessionID: token.SessionID, Version: newVersion}, AccessTokenExpired)
	if err != nil {
		return TokenResponse{}, err
	}
	refreshToken, err := u.Token.CreateToken(ctx, TokenEntity{UserID: token.UserID, SessionID: token.SessionID, Version: newVersion}, RefreshTokenExpired)
	if err != nil {
		return TokenResponse{}, err
	}
	err = u.Session.Rotate(ctx, SessionRotatePayload{SessionID: token.SessionID, RefreshToken: refreshToken, Version: newVersion, OldVersion: currentSession.Version})
	if err != nil {
		return TokenResponse{}, err
	}
	return TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
func (u *Usecase) Logout(ctx context.Context, refreshToken string) error {
	token, err := u.Token.VerifyToken(ctx, refreshToken)
	if err != nil {
		return err
	}
	currentSession, err := u.Session.GetByID(ctx, token.SessionID)
	if err != nil {
		return err
	}
	if token.Version != currentSession.Version {
		return ErrTokenExpired
	}
	if !currentSession.RevokedAt.IsZero() || currentSession.RevokedAt.Before(time.Now()) {
		return ErrTokenExpired
	}
	err = u.Session.Delete(ctx, SessionDeletePayload{SessionID: token.SessionID, RefreshToken: refreshToken, Version: token.Version})
	if err != nil {
		return err
	}
	return nil
}
func (u *Usecase) Verify(ctx context.Context, token string) (TokenEntity, error) {
	return u.Token.VerifyToken(ctx, token)
}
