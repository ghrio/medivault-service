package server

import (
	"fmt"
	"log"
	"medivault-service/cmd/routes"
	"medivault-service/config"
	handler "medivault-service/internal/handler"
	"medivault-service/internal/middleware"
	repository "medivault-service/internal/repository/sqlc"
	"medivault-service/internal/services"
	"net/http"

	_ "medivault-service/docs" // Import the generated docs

	"github.com/jackc/pgx/v5"
)

// @title           MediVault Service API
// @version         1.0
// @description     API documentation for MediVault Service.
// @contact.name    API Support
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8080
// @BasePath        /api/v1

func InitServer(cfg *config.ServerConfig, conn *pgx.Conn) error {
	r := http.NewServeMux()

	// Registering Repositories
	userRepo := repository.NewUserRepository(conn)
	userService := services.NewUserService(userRepo)
	userHandler := handler.UserHandler{Service: userService}
	authHandler := handler.AuthenticationHandler{Service: userService}
	// Registering Middleware
	userRouter := routes.UserRouter(&userHandler)
	authenticationRouter := routes.AuthenticationRouter(&authHandler)
	// Define routes and their corresponding handlers
	routes := []struct {
		pattern string
		handler http.Handler
	}{
		{"/api/v1/users/", userRouter},
		{"/api/v1/auth/", authenticationRouter},
	}

	// Register routes dynamically
	for _, route := range routes {
		r.Handle(route.pattern, route.handler)
	}

	loggedHandler := middleware.Logging(r)

	srv := &http.Server{
		Handler:      loggedHandler,
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}
	log.Print(cfg.Host)
	log.Printf("Starting server on %s:%d...", cfg.Host, cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
