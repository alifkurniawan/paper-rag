package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPAddr    string
	DatabaseURL string
	DBMaxConns  int32
	DBMinConss  int32
}

func Load() (Config, error) {
	maxConss, err := intEnv("DB_MAX_CONNS", 10)
	if err != nil {
		return Config{}, err
	}
	minConss, err := intEnv("DB_MIN_CONNS", 2)
	if err != nil {
		return Config{}, err
	}
	return Config{
		HTTPAddr:    ":8080",
		DatabaseURL: os.Getenv("DATABASE_URL"),
		DBMaxConns:  maxConss,
		DBMinConss:  minConss,
	}, nil
}

func env(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func intEnv(key string, defaultValue int32) (int32, error) {
	valueStr := env(key, "")
	if valueStr == "" {
		return defaultValue, nil
	}
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}
	return int32(valueInt), nil
}
