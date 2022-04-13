package driver

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/agschrei/go-crud-k8s-demo/internal/config"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const maxOpenDbCon = 5
const maxIdleDbCon = 3
const maxDbConKeepAlive = 5 * time.Minute

type PsqlDriverManager struct {
	dbConfig *config.DbConfig
	logger   *log.Logger
}

func NewPsqlDriverManager(dbConfig *config.DbConfig, logger *log.Logger) *PsqlDriverManager {
	return &PsqlDriverManager{
		dbConfig: dbConfig,
		logger:   logger,
	}
}

// NewDatabase creates a database using the configuration associated with the PsqlDriverManager
func (manager *PsqlDriverManager) NewDatabase() (*sql.DB, error) {
	manager.logger.Print("Creating database connection...")

	dsn, err := DsnFromDbConfig(manager.dbConfig)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbCon)
	db.SetMaxIdleConns(maxIdleDbCon)
	db.SetConnMaxLifetime(maxDbConKeepAlive)

	ctx, cancel := context.WithTimeout(context.Background(), manager.dbConfig.ConnectionTimeout)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	manager.logger.Print("Successfully created database connection.")

	return db, nil
}

func DsnFromDbConfig(dbConf *config.DbConfig) (string, error) {
	//TODO: consider using text/template to make this code a bit cleaner
	var sb strings.Builder

	prefix := "postgresql://"
	host := dbConf.Hostname
	port := dbConf.Port
	dbName := dbConf.DbName
	user := dbConf.User
	pass := dbConf.Pass
	conTimeout := int(dbConf.ConnectionTimeout.Seconds())
	sslDisabled := dbConf.SslDisabled

	hasParams := conTimeout != 0 || sslDisabled

	if host == "" {
		return "", errors.New("cannot construct DSN because no host is set in DbConfig")
	}

	sb.WriteString(prefix)
	if user != "" {
		sb.WriteString(user)
		if pass != "" {
			sb.WriteString(":")
			sb.WriteString(pass)
		}
		sb.WriteString("@")
	}
	sb.WriteString(host)

	if port != "" {
		sb.WriteString(":")
		sb.WriteString(port)
	}

	if dbName != "" {
		sb.WriteString("/")
		sb.WriteString(dbName)
	}

	if hasParams {
		sb.WriteString("?")
		atLeastOne := false
		if conTimeout != 0 {
			sb.WriteString("connect_timeout=")
			sb.WriteString(strconv.Itoa(int(conTimeout)))
			atLeastOne = true
		}
		if sslDisabled {
			if atLeastOne {
				sb.WriteString("&")
			}
			sb.WriteString("sslmode=disable")
		}
	}
	return sb.String(), nil
}
