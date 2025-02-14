package main

import (
	"log"
	"medivault-service/cmd/server"
	"medivault-service/config"
)

func main() {
	cfg, err := config.LoadConfigENV()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Println("Config loaded")
	db, err := config.InitDBConnection(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db connection: %v", err)
	}
	log.Println("DB connection initialized")

	err = server.InitServer(&cfg.Server, db)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}
}
