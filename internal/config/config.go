package config

import (
	"log"
	"time"
)

type Environment uint8

const (
	Development Environment = iota
	Production
)

type DbConfig struct {
	Hostname          string
	Port              string
	SslDisabled       bool
	User              string
	Pass              string
	DbName            string
	ConnectionTimeout time.Duration
}

type AppConfig struct {
	Environment Environment
	Logger      *log.Logger
	Port        uint16
	DbConfig    *DbConfig
}
