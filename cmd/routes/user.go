package routes

import (
	handlers "medivault-service/internal/handler"
	"net/http"
)

func UserRouter(userHandler *handlers.UserHandler) *http.ServeMux {
	userMux := http.NewServeMux()
	userMux.HandleFunc("POST /users", userHandler.CreateUser)
	userMux.HandleFunc("GET /users/{id}", userHandler.GetUserByID)
	userMux.HandleFunc("GET /users", userHandler.GetAllUsers)
	userMux.HandleFunc("POST /users/email/{email}", userHandler.GetUserByEmail)
	userMux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	userMux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)
	return userMux
}
