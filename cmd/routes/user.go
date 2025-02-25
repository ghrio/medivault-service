package routes

import (
	"medivault-service/internal/handler"
	"net/http"
)

func UserRouter(userHandler *handler.UserHandler) *http.ServeMux {
	userMux := http.NewServeMux()
	userMux.HandleFunc("POST /create", userHandler.CreateUser)
	userMux.HandleFunc("GET /user/{id}", userHandler.GetUserByID)
	userMux.HandleFunc("GET /all", userHandler.GetAllUsers)
	userMux.HandleFunc("POST /email/{email}", userHandler.GetUserByEmail)
	userMux.HandleFunc("PUT /{id}", userHandler.UpdateUser)
	userMux.HandleFunc("DELETE /{id}", userHandler.DeleteUser)
	return userMux
}
