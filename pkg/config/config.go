package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type (
	// Config of service
	Config struct {
		InstanceConfig
		DBConfig
		ClientsConfig
		RedisConfig
	}

	InstanceConfig struct {
		ServiceName    string `env:"SERVICE_NAME,required"`
		HTTPPort       string `env:"HTTP_PORT, default=8081"`
		PublicURL      string `env:"PUBLIC_URL,required"`
		CSVFileMB      string `env:"CSV_FILE_MB, default=DBZURRESULTFinal_utf8.csv"`
		CSVFileRB      string `env:"CSV_FILE_RB, default=resultfizFinal_utf8.csv"`
		RunMBMigration bool   `env:"RUN_MB_MIGRATION, default=true"`
	}

	DBConfig struct {
		PGURL string `env:"PG_URL,required"`
	}

	ClientsConfig struct {
		AISuflerAPIURL string `env:"AI_SUFLER_API_URL,required"`
	}

	RedisConfig struct {
		RedisAddr string `env:"REDIS_ADDR,required"`
	}
)

// Get
func Get() (Config, error) {
	return parseConfig()
}

func parseConfig() (cfg Config, err error) {
	godotenv.Load()

	err = envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return cfg, fmt.Errorf("fill config: %w", err)
	}

	return cfg, nil
}
