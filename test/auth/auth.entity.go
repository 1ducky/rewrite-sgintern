package auth

import "time"

type CredentialEntity struct {
	UserID   string
	Email    string
	Password string
	Role     Role
}

type SessionEntity struct {
	ID           string
	UserID       string
	RefreshToken string
	Version      Version
	DeviceName   string
	IP           string
	UserAgent    string
	ExpiresAt    time.Time
	RevokedAt    *time.Time
	CreatedAt    time.Time
}

type TokenEntity struct {
	UserID    string
	SessionID string
	Role      Role
	Version   Version
}
