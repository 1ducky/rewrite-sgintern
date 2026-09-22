package registry

import "RewriteProject/internal/app/services/auth"

type RegisterHandler struct {
	AuthHandler auth.Handler
}

func NewRegisterHandler(apps RegisterApp) RegisterHandler {
	return RegisterHandler{
		AuthHandler: auth.NewHandler(apps.authService),
	}
}
