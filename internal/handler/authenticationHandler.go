package handler

import (
	"medivault-service/internal/db/generated"
	"medivault-service/internal/services"
	"net/http"
)

type User struct {
	*generated.User
}

type AuthenticationHandler struct {
	Service services.UserService
}

func NewAuthenticationHandler(service services.UserService) *AuthenticationHandler {
	return &AuthenticationHandler{service}
}

// Login logs in a user
// @Summary      Login a user
// @Description  Logs in a user
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param        email     body  string  true  "User email"
// @Param        password  body  string  true  "User password"
// @Success      200   {object}  User
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /login [post]
func (ath *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	r.FormValue("email")
	r.FormValue("password")
	if r.FormValue("email") == "" || r.FormValue("password") == "" {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}
}

// Logout logs out a user
// @Summary      Logout a user
// @Description  Logs out a user
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Success      200   {object}  User
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /logout [post]
func (ath *AuthenticationHandler) Logout(w http.ResponseWriter, r *http.Request) {
}

// RefreshToken refreshes a user token
// @Summary      Refresh a user token
// @Description  Refreshes a user token
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Success      200   {object}  User
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /refresh [post]
func (ath *AuthenticationHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}
