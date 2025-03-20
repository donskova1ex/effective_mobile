package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	PgDSN string
	Port  string
}

func LoadConfig() (*config, error) {
	err := godotenv.Load("./env.dev")
	if err != nil {
		return nil, fmt.Errorf("failed to load env file: %w", err)
	}

	cfg := &config{}

	cfg.PgDSN = os.Getenv("PG_DSN")
	if cfg.PgDSN == "" {
		return nil, fmt.Errorf("PG_DSN is not set")
	}

	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		return nil, fmt.Errorf("PORT is not set")
	}
	return cfg, nil
}
