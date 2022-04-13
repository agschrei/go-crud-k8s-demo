package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/agschrei/go-crud-k8s-demo/internal/config"
)

const version = "0.1"

func main() {
	start := time.Now()

	appConfig := setUpConfig()
	app := startApplication(appConfig)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", appConfig.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	appConfig.Logger.Printf("Application setup took %dms", time.Since(start).Milliseconds())
	appConfig.Logger.Printf("Starting server on port %d", appConfig.Port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func setUpConfig() *config.AppConfig {
	//if this flag is set we log the file location the message originated from
	tracingEnabled := getEnvFlag("ENABLE_TRACE")
	logger := setUpLogger(tracingEnabled)
	logger.Println("Logger bootstrapped.")

	var applicationPort uint16
	if val, set := getEnvVariable("APPLICATION_PORT"); set {
		port, err := strconv.Atoi(val)
		if err == nil {
			applicationPort = uint16(port)
		}
	} else {
		applicationPort = 8080
	}

	isDev := getEnvFlag("ENABLE_DEV")
	var env config.Environment
	if isDev {
		env = config.Development
	} else {
		env = config.Production
	}

	var conTimeout time.Duration
	if val, set := getEnvVariable("DB_TIMEOUT"); set {
		seconds, err := strconv.Atoi(val)
		if err != nil {
			logger.Fatalf("Configuration failed, DB_TIMEOUT does not contain valid integer, was: %s", val)
		}
		conTimeout = time.Duration(seconds) * time.Second
	} else {
		conTimeout = 30 * time.Second
	}

	dbConfig := &config.DbConfig{
		Hostname:          getEnvVariableOrDefault("DB_HOST", "db"),
		Port:              getEnvVariableOrDefault("DB_PORT", "5432"),
		SslDisabled:       getEnvFlag("DB_SSL_DISABLE"),
		User:              getEnvVariableOrDefault("DB_USER", ""),
		Pass:              getEnvVariableOrDefault("DB_PASS", ""),
		DbName:            getEnvVariableOrDefault("DB_NAME", "demo"),
		ConnectionTimeout: conTimeout,
	}

	appConfig := config.AppConfig{
		Environment: env,
		Logger:      logger,
		Port:        applicationPort,
		DbConfig:    dbConfig,
	}

	return &appConfig
}

func getEnvVariableOrDefault(name string, defaultVal string) string {
	if val, set := getEnvVariable(name); set {
		return val
	} else {
		return defaultVal
	}
}

func getEnvVariable(name string) (string, bool) {
	if val, set := os.LookupEnv(name); set {
		return val, true
	}
	return "", false
}

// getEnvFlag checks if a variable is set to a truthy value, if it is it returns true, otherwise (including if the variable is not set) it returns false
func getEnvFlag(name string) bool {
	val, set := getEnvVariable(name)
	if set {
		flag, err := strconv.ParseBool(val)
		if err != nil {
			return flag
		}
	}
	return false
}

//setUpLogger configures an application-level logger based on environment variables and returns it
func setUpLogger(tracingEnabled bool) *log.Logger {

	loggerFlags := log.Ldate | log.Ltime | log.LUTC | log.Lmsgprefix
	if tracingEnabled {
		loggerFlags = loggerFlags | log.Llongfile
	}

	logger := log.New(os.Stdout, "", loggerFlags)
	return logger
}
