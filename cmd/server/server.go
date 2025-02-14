package server

import (
	"fmt"
	"log"
	"medivault-service/config"
	"medivault-service/internal/middleware"
	"net/http"

	_ "medivault-service/docs" // Import the generated docs

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           MediVault Service API
// @version         1.0
// @description     API documentation for MediVault Service.
// @contact.name    API Support
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8080
// @BasePath        /api/v1

func InitServer(cfg *config.ServerConfig) error {
	r := http.NewServeMux()

	// Serve Swagger UI
	r.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s:%d/swagger/doc.json", cfg.Host, cfg.Port)), // The URL pointing to API definition
	))

	// Registering Middleware
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
