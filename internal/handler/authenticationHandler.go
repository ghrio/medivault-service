package handler

import "net/http"

type AuthenticationHandler struct {
}

// Login User
func (ath AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {

}

// Logout User
func (ath AuthenticationHandler) Logout(w http.ResponseWriter, r *http.Request) {
}

// Refresh Token
func (ath AuthenticationHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}
