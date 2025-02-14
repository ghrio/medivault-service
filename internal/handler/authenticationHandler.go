package handler

import "net/http"

type AuthenticationHandler struct {
}

// Login User
func (ath *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	r.FormValue("email")
	r.FormValue("password")
	if r.FormValue("email") == "" || r.FormValue("password") == "" {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}
}

// Logout User
func (ath *AuthenticationHandler) Logout(w http.ResponseWriter, r *http.Request) {
}

// Refresh Token
func (ath *AuthenticationHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}
