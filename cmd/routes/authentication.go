package routes

import (
	handlers "medivault-service/internal/handler"
	"net/http"
)

func AuthenticationRouter(authHandler *handlers.AuthenticationHandler) *http.ServeMux {
	authMux := http.NewServeMux()
	authMux.HandleFunc("POST /login", authHandler.Login)
	authMux.HandleFunc("POST /logout", authHandler.Logout)
	authMux.HandleFunc("GET /refresh", authHandler.RefreshToken)
	return authMux
}
