package jwt

import (
	"RewriteProject/internal/config"
	"RewriteProject/test/auth"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	auth.TokenEntity
	jwt.RegisteredClaims
}

type Service struct {
	key    []byte
	Method *jwt.SigningMethodHMAC
}

func NewJWT(config *config.AuthConfig) auth.TokenContract {
	return &Service{key: []byte(config.SecretKeyJWT), Method: jwt.SigningMethodHS256}
}

func (r *Service) CreateToken(ctx context.Context, entity auth.TokenEntity, minutes auth.TimeInt) (string, error) {
	var secretKey = []byte(r.key)

	claims := AuthClaims{
		TokenEntity: entity,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(minutes) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "TodoApp",
		},
	}

	token := jwt.NewWithClaims(r.Method, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (r *Service) VerifyToken(ctx context.Context, tokenString string) (auth.TokenEntity, error) {
	claims := AuthClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrTokenInvalid
		}
		return []byte(r.key), nil
	})
	if err != nil {
		return auth.TokenEntity{}, auth.ErrTokenInvalid
	}

	if !token.Valid {
		return auth.TokenEntity{}, auth.ErrTokenInvalid
	}
	if claims.ExpiresAt.Before(time.Now()) {
		return auth.TokenEntity{}, auth.ErrTokenExpired
	}

	return claims.TokenEntity, nil
}
