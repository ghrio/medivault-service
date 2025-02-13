package server

import (
	"fmt"
	"log"
	"medivault-service/config"
	"net/http"
)

func InitServer(cfg *config.ServerConfig) error {
	r := http.NewServeMux()

	r.HandleFunc("GET /v1", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /v1")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GET /v1"))
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	log.Printf("Starting server on %s:%d...", cfg.Host, cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
