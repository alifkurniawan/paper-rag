package db

import (
	"context"
	"core/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, conf config.Config) (*pgxpool.Pool, error) {
	pgCfg, err := pgxpool.ParseConfig(conf.DatabaseURL)
	if err != nil {
		return nil, err
	}

	pgCfg.MaxConns = conf.DBMaxConns
	pgCfg.MinConns = conf.DBMinConss
	pgCfg.MaxConnLifetime = time.Hour
	pgCfg.MaxConnIdleTime = 30 * time.Minute
	pgCfg.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, pgCfg)
	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
