package auth

type SessionRotatePayload struct {
	SessionID    string
	RefreshToken string
	Version      Version
	OldVersion   Version
}
type SessionDeletePayload struct {
	SessionID    string
	RefreshToken string
	Version      Version
}

type CredentialPayload struct {
	Email    string
	Password string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
