package registry

import (
	"RewriteProject/internal/auth"
	"RewriteProject/internal/auth/credential"
	"RewriteProject/internal/auth/jwt"
	"RewriteProject/internal/auth/session"
	"RewriteProject/internal/config"
	"RewriteProject/internal/db"
	"RewriteProject/internal/profile"
)

type RegisterRepo struct {
	Cred    auth.CredentialRepositoryContract
	Token   auth.TokenContract
	Session auth.SessionRepositoryContract
	Profile profile.RepoContract
}

func NewRegisterRepo(persistent db.DBTX, conf config.AppConfig) RegisterRepo {
	return RegisterRepo{
		Cred:    credential.NewMySQLRepository(persistent),
		Token:   jwt.NewJWT(conf.AuthConfig),
		Session: session.NewMySQLRepository(persistent),
		Profile: profile.NewMySQLRepository(persistent),
	}
}
