package config

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"net/url"
	"os"
	"time"
)

type StorageConfig struct {
	Host     string        `yaml:"host" env-required:"true"`
	Port     string        `yaml:"port" env-required:"true"`
	User     string        `yaml:"user" env-required:"true"`
	Password string        `yaml:"password" env-required:"true"`
	Database string        `yaml:"database" env-required:"true"`
	Timeout  time.Duration `yaml:"timeout" env-default:"5"`
	MaxConns int32         `yaml:"max_connection" env-default:"10"`
}

func NewStorageConfig() (StorageConfig, error) {
	const op = "repository.postgres.config.NewStorageConfig"

	err := godotenv.Load(".env")
	if err != nil {
		return StorageConfig{}, fmt.Errorf("%s - error loading .env file: %w", op, err)
	}
	configPath := os.Getenv("POSTGRES_CONFIG_PATH")
	if configPath == "" {
		return StorageConfig{}, fmt.Errorf("%s - config %s not found, err: %w", op, configPath, err)
	}

	var cfg StorageConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return StorageConfig{}, fmt.Errorf("%s - error reading config: %w", op, err)
	}
	return cfg, nil
}

func NewPoolConfig(cfg *StorageConfig) (*pgxpool.Config, error) {
	const op = "repository.postgres.config.NewPoolConfig"

	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.User),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Timeout)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return poolConfig, nil
}

func NewConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	const op = "repository.postgres.config.NewConnection"

	conn, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return conn, nil
}
