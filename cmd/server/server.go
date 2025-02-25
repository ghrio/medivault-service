package server

import (
	"fmt"
	"log"
	"medivault-service/cmd/routes"
	"medivault-service/config"
	_ "medivault-service/docs" // Import the generated docs
	handler "medivault-service/internal/handler"
	"medivault-service/internal/middleware"
	repository "medivault-service/internal/repository/sqlc"
	"medivault-service/internal/services"
	"net/http"

	"github.com/jackc/pgx/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			MediVault Service API
//	@version		1.0
//	@description	API documentation for MediVault Service.
//	@contact.name	API Support
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@BasePath		/api/v1

func InitServer(cfg *config.ServerConfig, conn *pgx.Conn) error {

	swaggerURL := fmt.Sprintf("http://%s:%d/api/v1/swagger/doc.json", cfg.Host, cfg.Port)
	r := http.NewServeMux()

	// Registering Repositories
	userRepo := repository.NewUserRepository(conn)
	patientFile := repository.NewPatientFileRepository(conn)
	patientFileCollection := repository.NewPatientFileCollectionRepository(conn)

	// Registering Services
	userService := services.NewUserService(userRepo)
	patientFileService := services.NewPatientFileService(patientFile)
	patientFileCollectionService := services.NewPatientFileCollectionService(patientFileCollection)

	// Registering Handlers
	patientFileCollectionHandler := handler.PatientFileCollectionHandler{Service: patientFileCollectionService}
	userHandler := handler.UserHandler{Service: userService}
	authHandler := handler.AuthenticationHandler{Service: userService}
	patientFileHandler := handler.PatientFileHandler{Service: patientFileService}
	// Registering Routes
	patientFileCollectionRouter := routes.PatientFileCollectionRouter(&patientFileCollectionHandler)
	patientFileRouter := routes.PatientFileRouter(&patientFileHandler)
	userRouter := routes.UserRouter(&userHandler)
	authenticationRouter := routes.AuthenticationRouter(&authHandler)

	// Define routes and their corresponding handlers
	// Register subrouters with the main router
	r.Handle("/api/v1/patientFileCollections/", http.StripPrefix("/api/v1/patientFileCollections", patientFileCollectionRouter))
	r.Handle("/api/v1/patientFiles/", http.StripPrefix("/api/v1/patientFiles", patientFileRouter))
	r.Handle("/api/v1/users/", http.StripPrefix("/api/v1/users", userRouter))
	r.Handle("/api/v1/auth/", http.StripPrefix("/api/v1/auth", authenticationRouter))

	r.Handle("/api/v1/swagger/", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
	))
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
