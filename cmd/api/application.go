package main

import (
	"database/sql"

	"github.com/agschrei/integration-test-sample/internal/config"
	"github.com/agschrei/integration-test-sample/internal/driver"
)

type Application struct {
	config *config.AppConfig
	dbCon  *sql.DB
}

func startApplication(config *config.AppConfig) *Application {
	pgdriver := driver.NewPsqlDriverManager(config.DbConfig, config.Logger)
	db, err := pgdriver.NewDatabase()
	if err != nil {
		config.Logger.Fatalf("Could not establish database connection: %s", err)
	}
	return &Application{config, db}
}
