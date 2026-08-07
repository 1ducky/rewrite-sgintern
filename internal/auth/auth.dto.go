package auth

type AuthEntity struct {
	ID       string
	Username string
	Role     Role
	Version  int
}
type AuthLogin struct {
	ID       string
	Email    string
	Password string
	Username string
	Role     Role
}
type RegisterPayload struct {
	ID       string
	Email    string
	Password string
	Username string
}
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Session
type CreateSessionPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}
type UpdateSessionPayload struct {
	AccessToken     string `json:"access_token"`
	OldrefreshToken string `json:"old_refresh_token"`
	RefreshToken    string `json:"refresh_token"`
	UserID          string `json:"user_id"`
	Version         int    `json:"version"`
}

type DeleteSessionPayload struct {
	UserID       string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	Version      int    `json:"version"`
}
