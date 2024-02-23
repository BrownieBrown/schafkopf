package utils

import (
	"errors"
	"github.com/BrownieBrown/schafkopf/internal/models"
	"os"
)

func LoadConfig() (models.Config, error) {
	sqliteConfig, err := loadSqliteConfig()
	if err != nil {
		return models.Config{}, err
	}

	return models.Config{SqliteConfig: sqliteConfig}, nil
}

func loadSqliteConfig() (models.SqliteConfig, error) {
	cfg := models.SqliteConfig{
		DataSourceName: os.Getenv("SQLITE_DATA_SOURCE_NAME"),
		DriverName:     os.Getenv("SQLITE_DRIVER_NAME"),
	}

	if cfg.DataSourceName == "" {
		return models.SqliteConfig{}, errors.New("SQLITE_DATA_SOURCE_NAME is not set")
	}

	if cfg.DriverName == "" {
		return models.SqliteConfig{}, errors.New("SQLITE_DRIVER_NAME is not set")
	}

	return cfg, nil
}
