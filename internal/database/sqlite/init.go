package sqlite

import (
	"database/sql"
	"github.com/BrownieBrown/schafkopf/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	database *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Close() error {
	return d.database.Close()
}

func (d *Database) Open(cfg models.SqliteConfig) error {
	dataSourceName := cfg.DataSourceName
	driverName := cfg.DriverName
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	d.database = db
	return nil
}

func Init(cfg models.SqliteConfig) (*Database, error) {
	sqliteDB := NewDatabase()
	err := sqliteDB.Open(cfg)
	if err != nil {
		return nil, err
	}
	defer sqliteDB.Close()

	err = sqliteDB.createPlayerTable()
	if err != nil {
		return nil, err
	}

	err = sqliteDB.createGameTable()
	if err != nil {
		return nil, err
	}

	err = sqliteDB.createSessionTable()
	if err != nil {
		return nil, err
	}

	return sqliteDB, nil
}

func (d *Database) createPlayerTable() error {
	_, err := d.database.Exec(`
		CREATE TABLE IF NOT EXISTS player (
			id TEXT PRIMARY KEY,
			score FLOAT
		);
	`)

	return err
}

func (d *Database) createGameTable() error {
	_, err := d.database.Exec(`
		CREATE TABLE IF NOT EXISTS game (
			id TEXT PRIMARY KEY,
			created_at TEXT
		);
	`)

	return err
}

func (d *Database) createSessionTable() error {
	_, err := d.database.Exec(`
		CREATE TABLE IF NOT EXISTS session (
			id TEXT PRIMARY KEY,
			game_id TEXT,
			FOREIGN KEY (game_id) REFERENCES game(id)
		);
	`)

	return err
}
