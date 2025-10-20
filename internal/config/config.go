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

	// Database (pgx pool)
	DatabaseURL       string        `env:"DATABASE_URL" envDefault:"postgres://user:password@localhost:5432/subscriptions?sslmode=disable"`
	MaxConns          int32         `env:"DB_MAX_CONNS" envDefault:"25"`
	MinConns          int32         `env:"DB_MIN_CONNS" envDefault:"5"`
	MaxConnLifetime   time.Duration `env:"DB_MAX_CONN_LIFETIME" envDefault:"1h"`
	MaxConnIdleTime   time.Duration `env:"DB_MAX_CONN_IDLE_TIME" envDefault:"30m"`
	HealthCheckPeriod time.Duration `env:"DB_HEALTH_CHECK_PERIOD" envDefault:"1m"`

	// Logger
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
}

func NewConfig() *Config {
	cfg := &Config{}
	env.Parse(cfg)
	return cfg
}
