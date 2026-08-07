package auth

import "context"

type UsecaseContract interface {
	Login(ctx context.Context, payload LoginPayload) (TokenResponse, error)  //compare email and hash password
	Refresh(ctx context.Context, refreshToken string) (TokenResponse, error) //refresh old token expired,
	Logout(ctx context.Context, refreshToken string) (string, error)         //revoke token
	Verify(ctx context.Context, token string) (TokenEntity, error)           //verify valid token
}

type RepositoryContract interface {
	FindByEmail(ctx context.Context, email string) (AuthEntity, error)
	Login(ctx context.Context, email string) (AuthLogin, error)
	Create(ctx context.Context, entity RegisterPayload) (AuthEntity, error)
	GetVersion(ctx context.Context, userId string) (int, error)
}

type CredentialRepositoryContract interface {
	Login(ctx context.Context, payload LoginPayload) (CredentialEntity, error)  //find credential by email
	Update(ctx context.Context, payload LoginPayload) (CredentialEntity, error) //update credential
	RevokeVersion(ctx context.Context, Token string) error                      //update version to current version + 1 to revoke all token
}

type SessionUsecase interface {
	Create(ctx context.Context, payload CreateSessionPayload) (SessionEntity, error)                  //save token to session
	Rotate(ctx context.Context, payload UpdateSessionPayload) (SessionEntity, error)                  //update token
	Delete(ctx context.Context, payload DeleteSessionPayload) (SessionEntity, error)                  //delete token from session
	GetByAccessToken(ctx context.Context, UserID string, accessToken string) (SessionEntity, error)   //get token by access token
	GetByRefreshToken(ctx context.Context, userID string, refreshToken string) (SessionEntity, error) //get token by refresh token
}

type TokenContract interface {
	CreateToken(ctx context.Context, payload TokenEntity, minute int) (string, error) //create token with entity
	VerifyToken(ctx context.Context, token string) (TokenEntity, error)               //verify token and return entity
}

const ContextAuthEntityKey = "user-auth" //key to get user from context
