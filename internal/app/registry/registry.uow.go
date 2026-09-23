package registry

import (
	authApp "RewriteProject/internal/app/services/auth"
	"RewriteProject/internal/app/uow"
	"RewriteProject/internal/auth"
	"RewriteProject/internal/auth/credential"
	"RewriteProject/internal/auth/jwt"
	"RewriteProject/internal/auth/session"
	"RewriteProject/internal/config"
	"RewriteProject/internal/db"
	"RewriteProject/internal/profile"
	"database/sql"
)

type RegisterUoW struct {
	RegisterCase uow.UnitOfWork[authApp.TXUsecase]
}

func NewUoWCase(persistent *sql.DB, conf config.AppConfig) RegisterUoW {
	return RegisterUoW{
		RegisterCase: buildRegisterCase(persistent, conf),
	}
}

func buildRegisterCase(persistent *sql.DB, conf config.AppConfig) uow.UnitOfWork[authApp.TXUsecase] {
	return uow.NewUoW(persistent, func(tx db.DBTX) authApp.TXUsecase {
		return authApp.TXUsecase{
			Auth:    auth.NewAuthService(credential.NewMySQLRepository(tx), jwt.NewJWT(conf.AuthConfig), session.NewMySQLRepository(tx)),
			Profile: profile.NewUsecase(profile.NewMySQLRepository(tx)),
		}
	})
}
