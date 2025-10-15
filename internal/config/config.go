package config

import (
	"time"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	// HTTP Server
	HTTPPort          string        `env:"HTTP_PORT" envDefault:"8080"`
	ReadTimeout       time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"10s"`
	WriteTimeout      time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"10s"`
	IdleTimeout       time.Duration `env:"HTTP_IDLE_TIMEOUT" envDefault:"60s"`
	ReadHeaderTimeout time.Duration `env:"HTTP_READ_HEADER_TIMEOUT" envDefault:"5s"`

	// Database
	DatabaseURL       string        `env:"DATABASE_URL" envDefault:"postgres://user:password@localhost:5432/subscriptions?sslmode=disable"`
	DBMaxOpenConns    int           `env:"DB_MAX_OPEN_CONNS" envDefault:"5"`
	DBMaxIdleConns    int           `env:"DB_MAX_IDLE_CONNS" envDefault:"2"`
	DBConnMaxIdleTime time.Duration `env:"DB_CONN_MAX_IDLE_TIME" envDefault:"2m"`
	DBConnMaxLifetime time.Duration `env:"DB_CONN_MAX_LIFETIME" envDefault:"30m"`

	// Logger
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
