package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/donskova1ex/effective_mobile/internal"
	"github.com/donskova1ex/effective_mobile/internal/middleware"
	"github.com/donskova1ex/effective_mobile/internal/processors"
	"github.com/donskova1ex/effective_mobile/internal/repositories"
	openapi "github.com/donskova1ex/effective_mobile/openapi"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logJSONHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logJSONHandler)

	slog.SetDefault(logger)

	logger.Info("Starting server")

	err := godotenv.Load("config/.env.dev")

	if err != nil {
		logger.Error("Error loading .env file", slog.String("err", err.Error()))
	}
	pgDSN := os.Getenv("POSTGRES_DSN")
	if pgDSN == "" {
		logger.Error("empty POSTGRES_DSN")
		os.Exit(1)
	}

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		logger.Error("empty API_PORT")
		os.Exit(1)
	}
	logger.Info("Config loaded", slog.String("port", apiPort))

	db, err := repositories.NewPostgresDB(ctx, pgDSN)
	if err != nil {
		logger.Error("Failed to connect to postgres", slog.String("err", err.Error()))
		os.Exit(1)
	}
	logger.Info("Connected to postgres")

	repository := repositories.NewRepository(db, logger)
	songsProcessor := processors.NewSongProcessor(repository, logger)

	DefaultAPIService := openapi.NewDefaultAPIService(songsProcessor, logger)
	DefaultAPIController := openapi.NewDefaultAPIController(DefaultAPIService)

	router := openapi.NewRouter(DefaultAPIController)
	requestLogger := middleware.RequestLogger(logger)
	router.Use(middleware.RequestIDMiddleware, requestLogger)

	httpServer := &http.Server{
		Addr:     ":" + apiPort,
		Handler:  router,
		ErrorLog: slog.NewLogLogger(logJSONHandler, slog.LevelError),
	}

	gracefulCloser := internal.NewGracefulCloser()

	gracefulCloser.Add(func() error {
		logger.Info("closing db connection")
		if err := db.Close(); err != nil {
			logger.Error("failed to close db connection", slog.String("err", err.Error()))
			return err
		}
		logger.Info("db connection closed")
		return nil
	})

	gracefulCloser.Add(func() error {
		logger.Info("closing http server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(ctx); err != nil {
			logger.Error("failed to shutdown http server", slog.String("err", err.Error()))
			return err
		}
		logger.Info("http server closed successfully")
		return nil
	})

	shutdownCtx, shutdownCancel := context.WithCancel(ctx)
	defer shutdownCancel()

	go func() {
		ctx, cancel := context.WithCancel(shutdownCtx)
		defer cancel()
		gracefulCloser.Run(ctx, logger)
		os.Exit(1)
	}()

	logger.Info("Server started", slog.String("port", apiPort))

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Failed to start server", slog.String("err", err.Error()))
	}
	logger.Info("Server successfully stopped")

}
