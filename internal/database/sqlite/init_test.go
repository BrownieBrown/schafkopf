package sqlite

import (
	"github.com/BrownieBrown/schafkopf/internal/models"
	"testing"
)

func TestOpen(t *testing.T) {
	db := NewDatabase()
	cfg := models.SqliteConfig{DataSourceName: ":memory:", DriverName: "sqlite3"}

	err := db.Open(cfg)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer func(db *Database) {
		err := db.Close()
		if err != nil {
			t.Fatalf("Failed to close database: %v", err)
		}
	}(db)
}

func TestOpenError(t *testing.T) {
	db := NewDatabase()
	cfg := models.SqliteConfig{DataSourceName: "invalid", DriverName: "invalid"}

	err := db.Open(cfg)
	if err == nil {
		t.Fatalf("Open() did not return an error")
	}

}

func TestClose(t *testing.T) {
	db := NewDatabase()
	cfg := models.SqliteConfig{DataSourceName: ":memory:", DriverName: "sqlite3"}

	err := db.Open(cfg)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	err = db.Close()
	if err != nil {
		t.Fatalf("Failed to close database: %v", err)
	}
}

func TestInit(t *testing.T) {
	cfg := models.SqliteConfig{DataSourceName: ":memory:", DriverName: "sqlite3"}

	db, err := Init(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer func(db *Database) {
		err := db.Close()
		if err != nil {
			t.Fatalf("Failed to close database: %v", err)
		}
	}(db)
}

func TestInitError(t *testing.T) {
	cfg := models.SqliteConfig{DataSourceName: "invalid", DriverName: "invalid"}

	_, err := Init(cfg)
	if err == nil {
		t.Fatalf("Init() did not return an error")
	}
}
