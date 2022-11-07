package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/Meystergod/online-store-api/pkg/logging"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PgConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewPgConfig(username string, password string, host string, port string, database string) *PgConfig {
	return &PgConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func NewClient(ctx context.Context, maxAttempts int, maxDelay time.Duration, cfg *PgConfig) (pool *pgxpool.Pool, err error) {
	logger := logging.GetLogger(ctx)

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	err = DoWithAttempts(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pgxCfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			logger.Fatalf("unable to parse config: %v\n", err)
		}

		pool, err = pgxpool.ConnectConfig(ctx, pgxCfg)
		if err != nil {
			logger.Info("failed to connect to postgres... going to do the next attempt")

			return err
		}

		return nil
	}, maxAttempts, maxDelay)

	if err != nil {
		logger.Fatal("all attempts are exceeded. unable to connect to postgres")
	}

	return pool, nil
}

func DoWithAttempts(fn func() error, maxAttempts int, delay time.Duration) error {
	var err error

	for maxAttempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			maxAttempts--
			continue
		}
		return nil
	}
	return err
}
