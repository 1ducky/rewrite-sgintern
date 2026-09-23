package registry

import (
	authApp "RewriteProject/internal/app/services/auth"
)

type RegisterApp struct {
	authService authApp.AuthApp
}

func NewRegisterApp(ucs RegisterUC, uows RegisterUoW) RegisterApp {
	return RegisterApp{
		authService: *authApp.NewAuthApp(ucs.auth, ucs.profile, uows.RegisterCase),
	}
}
