package main

import (
	"github.com/agschrei/go-crud-k8s-demo/internal/config"
	"github.com/agschrei/go-crud-k8s-demo/internal/driver"
	"github.com/agschrei/go-crud-k8s-demo/internal/repository"
)

type application struct {
	config     *config.AppConfig
	repository repository.Repository
}

func startApplication(config *config.AppConfig) *application {
	pgdriver := driver.NewPsqlDriverManager(config.DbConfig, config.Logger)
	db, err := pgdriver.NewDatabase()
	if err != nil {
		config.Logger.Fatalf("Could not establish database connection: %s", err)
	}

	repository := repository.NewPostgresRepository(db)

	return &application{
		config:     config,
		repository: repository,
	}
}
