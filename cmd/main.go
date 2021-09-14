package main

import (
	"context"
	"log"
	"time"

	"faroukelkholy/survey/config"
	"faroukelkholy/survey/internal/server"
	"faroukelkholy/survey/internal/service/survey"
	"faroukelkholy/survey/internal/storage"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := storage.New(ctx, &storage.Options{
		DBHost: cfg.DBHost,
		DBPort: cfg.DBPort,
		DBName: cfg.DBName,
		DBUser: cfg.DBUser,
		DBPass: cfg.DBPass,
	})
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New()
	srv.AddSurveyFormRoutes(survey.New(repo))

	if err = srv.Start(cfg.HTTPPort); err != nil {
		log.Fatal(err)
	}
}
