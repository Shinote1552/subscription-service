package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"subscription-service/internal/config"
	"subscription-service/internal/http/handlers/subscriptionhand"
	"subscription-service/internal/http/middlewares"
	"subscription-service/internal/services/subscription"
	"subscription-service/internal/storage"
	"syscall"
	"time"

	"github.com/gorilla/mux"
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

	// 6. Инициализация transaction manager для pgx/v5
	trManager := manager.Must(trmpgx.NewDefaultFactory(pool))
	ctxGetter := trmpgx.DefaultCtxGetter

	// 7. Инициализация зависимостей
	subscriptionRepo := storage.NewStorage(pool, trManager, ctxGetter)
	subscriptionService := subscription.NewService(subscriptionRepo)

	// 8. Создание роутера и настройка маршрутов
	router := mux.NewRouter()

	// Global middlewares
	router.Use(middlewares.Logger(logMy))
	router.Use(middlewares.Compressor(logMy))
	router.Use(middlewares.Recovery(logMy))

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := pool.Ping(r.Context()); err != nil {
			http.Error(w, "Database unavailable", http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "healthy"}`))
	}).Methods("GET")

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Subscription routes (CRUDL + агрегация)
	api.Handle("/subscription", subscriptionhand.Create(subscriptionService, logMy)).Methods("POST")         // Create
	api.Handle("/subscription/{id}", subscriptionhand.Get(subscriptionService, logMy)).Methods("GET")        // Read
	api.Handle("/subscription/{id}", subscriptionhand.Update(subscriptionService, logMy)).Methods("PUT")     // Update
	api.Handle("/subscription/{id}", subscriptionhand.Delete(subscriptionService, logMy)).Methods("DELETE")  // Delete
	api.Handle("/subscription", subscriptionhand.List(subscriptionService, logMy)).Methods("GET")            // List
	api.Handle("/subscription/summary", subscriptionhand.Summary(subscriptionService, logMy)).Methods("GET") // Summary

	// 9. Создание и настройка HTTP сервера
	httpServer := &http.Server{
		Addr:              ":" + cfgMy.HTTPPort,
		Handler:           router,
		ReadTimeout:       cfgMy.ReadTimeout,
		WriteTimeout:      cfgMy.WriteTimeout,
		IdleTimeout:       cfgMy.IdleTimeout,
		ReadHeaderTimeout: cfgMy.ReadHeaderTimeout,
	}

	// 10. shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Запуск сервера
	go func() {
		logMy.Info().
			Str("port", cfgMy.HTTPPort).
			Dur("read_timeout", cfgMy.ReadTimeout).
			Dur("write_timeout", cfgMy.WriteTimeout).
			Msg("Starting HTTP server")

		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logMy.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	// Ожидание сигнала завершения
	sig := <-stop
	logMy.Info().Str("signal", sig.String()).Msg("Received signal, shutting down server...")

	// shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(ctxRoot, 30*time.Second)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logMy.Error().Err(err).Msg("Server shutdown error")
	} else {
		logMy.Info().Msg("Server shutdown completed successfully")
	}
}
