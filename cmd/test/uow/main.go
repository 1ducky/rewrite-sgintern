package main

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
	"context"
)

func main() {
	conf, _ := config.LoadConfig()

	dbmysql, err := db.NewMySQLDatabase(&conf.DBConfig)
	if err != nil {
		return
	}
	authUC := auth.NewAuthService(credential.NewMySQLRepository(dbmysql), jwt.NewJWT(conf.AuthConfig), session.NewMySQLRepository(dbmysql))
	profileUC := profile.NewUsecase(profile.NewMySQLRepository(dbmysql))

	register := uow.NewUoW(dbmysql, func(tx db.DBTX) authApp.TXUsecase {
		return authApp.TXUsecase{
			Auth:    auth.NewAuthService(credential.NewMySQLRepository(tx), jwt.NewJWT(conf.AuthConfig), session.NewMySQLRepository(tx)),
			Profile: profile.NewUsecase(profile.NewMySQLRepository(tx)),
		}
	})
	authService := authApp.NewAuthApp(authUC, profileUC, register)

	authService.Register(context.Background(), authApp.RegisterRequest{
		Email:    "test@gmail.com",
		Password: "18282828",
		Username: "User1",
	})

}
