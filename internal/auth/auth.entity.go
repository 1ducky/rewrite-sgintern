package auth

import "time"

type TokenEntity struct {
	ID      string
	Role    Role
	Version int
}

type CredentialEntity struct {
	ID       string
	Email    string
	Password string
	Version  int
}

type SessionEntity struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	RevokeAt     time.Time `json:"revoke_at"`
	Version      int       `json:"version"`
}
