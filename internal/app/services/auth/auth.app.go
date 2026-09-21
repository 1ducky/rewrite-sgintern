package auth

import (
	"RewriteProject/internal/app/uow"
	"RewriteProject/internal/auth"
	"RewriteProject/internal/profile"
	"context"
	"strings"

	"github.com/google/uuid"
)

type TXUsecase struct {
	Auth    auth.UsecaseContract
	Profile profile.UsecaseContract
}

type AuthApp struct {
	uow     uow.UnitOfWork[TXUsecase]
	auth    auth.UsecaseContract
	profile profile.UsecaseContract
}

func NewAuthApp(authUsecase auth.UsecaseContract, profileUsecase profile.UsecaseContract, uow uow.UnitOfWork[TXUsecase]) *AuthApp {
	return &AuthApp{
		auth:    authUsecase,
		profile: profileUsecase,
		uow:     uow,
	}
}

func (a *AuthApp) Login(ctx context.Context, req LoginRequest) (auth.TokenResponse, error) {
	token, err := a.auth.Login(ctx, auth.LoginPayload{Email: req.Email, Password: req.Password})
	if err != nil {
		return auth.TokenResponse{}, err
	}
	return token, nil
}
func (a *AuthApp) Register(ctx context.Context, req RegisterRequest) error {
	err := a.uow.Do(ctx, func(ctx context.Context, r TXUsecase) error {
		userID := uuid.NewString()
		err := r.Auth.Register(ctx, auth.RegisterPayload{UserID: userID, Email: req.Email, Password: req.Password})
		if err != nil {
			return err
		}
		err = r.Profile.CreateUser(ctx, profile.CreateRequest{ID: userID, Username: req.Username, Tag: strings.Split(req.Username, "")[0] + userID[0:5]})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
