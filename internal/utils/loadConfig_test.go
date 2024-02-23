package utils

import (
	"github.com/BrownieBrown/schafkopf/internal/models"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	dataSourceName := "testDataSourceName"
	driverName := "testDriverName"
	setEnv(t, "SQLITE_DATA_SOURCE_NAME", dataSourceName)
	setEnv(t, "SQLITE_DRIVER_NAME", driverName)
	defer unsetEnv(t, "SQLITE_DATA_SOURCE_NAME")
	defer unsetEnv(t, "SQLITE_DRIVER_NAME")

	got, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() returned an error: %v", err)

	}
	sqlitecfg := models.SqliteConfig{DataSourceName: dataSourceName, DriverName: driverName}
	want := models.Config{SqliteConfig: sqlitecfg}

	assertConfig(t, got, want)
}

func TestLoadConfigMissingDataSourceName(t *testing.T) {
	setEnv(t, "SQLITE_DRIVER_NAME", "testDriverName")
	defer unsetEnv(t, "SQLITE_DRIVER_NAME")

	_, err := LoadConfig()
	if err == nil {
		t.Fatalf("LoadConfig() did not return an error")
	}
}

func TestLoadConfigMissingDriverName(t *testing.T) {
	setEnv(t, "SQLITE_DATA_SOURCE_NAME", "testDataSourceName")
	defer unsetEnv(t, "SQLITE_DATA_SOURCE_NAME")

	_, err := LoadConfig()
	if err == nil {
		t.Fatalf("LoadConfig() did not return an error")
	}
}

func assertConfig(t *testing.T, got, want models.Config) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func setEnv(t *testing.T, key, value string) {
	t.Helper()
	err := os.Setenv(key, value)
	if err != nil {
		t.Fatalf("could not set environment variable: %v", err)
	}
}

func unsetEnv(t *testing.T, key string) {
	t.Helper()
	err := os.Unsetenv(key)
	if err != nil {
		t.Fatalf("could not unset environment variable: %v", err)
	}
}
