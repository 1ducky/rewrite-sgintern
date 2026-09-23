package auth

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	AuthApp AuthApp
}

func NewHandler(authService AuthApp) Handler {
	return Handler{
		AuthApp: authService,
	}
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) error {
	var payload LoginRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	res, err := h.AuthApp.Login(r.Context(), payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	cookie := http.Cookie{
		Name:     CookieSessionName,
		Value:    res.token.AccessToken, // The unencrypted data
		Path:     CookieSessionPath,     // Accessible across the entire site
		MaxAge:   CookieSessionMaxAge,   // Cookie expires in 1 hour (in seconds)
		HttpOnly: CookieSessionHTTPOnly, // Prevents JavaScript access
		Secure:   CookieSessionSecure,   // Recommended: Only sent over HTTPS
		SameSite: http.SameSiteLaxMode,  // Recommended: Protects against CSRF
	}

	// Send the cookie to the client in the response header
	http.SetCookie(w, &cookie)

	return json.NewEncoder(w).Encode(res.Profile)
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) error {
	var payload RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	err = h.AuthApp.Register(r.Context(), payload)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(nil)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}
