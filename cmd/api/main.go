package main

import (
	"log"
	"medivault-service/config"
	"medivault-service/server"
)

func main() {
	cfg, err := config.LoadConfigENV()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = config.InitDBConnection(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database connection: %v", err)
	}

	err = server.InitServer(&cfg.Server)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}
}
