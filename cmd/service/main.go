package main

import (
	"context"
	"fmt"
	"os"
	"subscription-service/internal/config"
	"subscription-service/internal/storage"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/rs/zerolog"
)

func main() {
	// 1. Инициализация логгера
	logMy := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()

	// 2. Создание конфигурации
	cfgMy := config.NewConfig()

	// 3. Установка уровня логирования из конфига
	if level, err := zerolog.ParseLevel(cfgMy.LogLevel); err == nil {
		logMy = logMy.Level(level)
	}

	// 4. Инициализация корневого контекста
	ctxRoot := context.Background()

	// 5. Подключение к PostgreSQL
	poolConfig, err := pgxpool.ParseConfig(cfgMy.DatabaseURL)
	if err != nil {
		logMy.Fatal().Err(err).Msg(fmt.Errorf("Failed to parse database config: %w", err).Error())
	}

	poolConfig.MaxConns = cfgMy.MaxConns
	poolConfig.MinConns = cfgMy.MinConns
	poolConfig.MaxConnLifetime = cfgMy.MaxConnLifetime
	poolConfig.MaxConnIdleTime = cfgMy.MaxConnIdleTime
	poolConfig.HealthCheckPeriod = cfgMy.HealthCheckPeriod

	pool, err := pgxpool.NewWithConfig(ctxRoot, poolConfig)
	if err != nil {
		logMy.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer pool.Close()

	pingCtx, cancel := context.WithTimeout(ctxRoot, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		logMy.Fatal().Err(err).Msg("Failed to ping database")
	}

	logMy.Info().
		Int32("max_conns", cfgMy.MaxConns).
		Int32("min_conns", cfgMy.MinConns).
		Str("max_conn_lifetime", cfgMy.MaxConnLifetime.String()).
		Str("max_conn_idle_time", cfgMy.MaxConnIdleTime.String()).
		Str("health_check_period", cfgMy.HealthCheckPeriod.String()).
		Msg("Successfully connected to PostgreSQL")

	// 6. Инициализация transaction manager
	trManager := manager.Must(trmpgx.NewDefaultFactory(pool))
	ctxGetter := trmpgx.DefaultCtxGetter

	// 7. Инициализация зависимостей
	storage := storage.NewStorage(pool, ctxGetter)
	subscriptionService := subscriptions.NewService(storage, trManager)

}
