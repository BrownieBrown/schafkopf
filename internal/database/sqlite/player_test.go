package sqlite

import (
	"github.com/BrownieBrown/schafkopf/internal/models"
	"testing"
)

func TestCreatePlayer(t *testing.T) {
	db := setupDatabase(t)
	defer func(db *Database) {
		err := db.Close()
		if err != nil {
			t.Fatalf("Failed to close database: %v", err)
		}
	}(db)
	repo := NewPlayerRepository(db.database)

	id := createPlayer(t, repo)

	if id == "" {
		t.Fatalf("CreatePlayer() returned an empty id")
	}

}

func TestGetPlayer(t *testing.T) {
	db := setupDatabase(t)
	defer func(db *Database) {
		err := db.Close()
		if err != nil {
			t.Fatalf("Failed to close database: %v", err)
		}
	}(db)
	repo := NewPlayerRepository(db.database)
	id := createPlayer(t, repo)

	getPlayer(t, repo, id)
}

func TestUpdatePlayer(t *testing.T) {
	db := setupDatabase(t)
	defer func(db *Database) {
		err := db.Close()
		if err != nil {
			t.Fatalf("Failed to close database: %v", err)
		}
	}(db)
	repo := NewPlayerRepository(db.database)
	id := createPlayer(t, repo)

	player := getPlayer(t, repo, id)
	player.Score = 100
	err := repo.UpdatePlayer(player)
	if err != nil {
		t.Fatalf("Failed to update player: %v", err)
	}

	updatedPlayer := getPlayer(t, repo, id)
	if updatedPlayer.Score != 100 {
		t.Fatalf("UpdatePlayer() did not update the player")
	}
}

func TestDeletePlayer(t *testing.T) {
	db := setupDatabase(t)
	defer func(db *Database) {
		err := db.Close()
		if err != nil {
			t.Fatalf("Failed to close database: %v", err)
		}
	}(db)
	repo := NewPlayerRepository(db.database)
	id := createPlayer(t, repo)

	err := repo.DeletePlayer(id)
	if err != nil {
		t.Fatalf("Failed to delete player: %v", err)
	}
}

func createPlayer(t *testing.T, repo models.PlayerRepository) string {
	player := models.Player{}
	id, err := repo.CreatePlayer(&player)
	if err != nil {
		t.Fatalf("Failed to create player: %v", err)
	}

	return id
}

func getPlayer(t *testing.T, repo models.PlayerRepository, id string) *models.Player {
	player, err := repo.GetPlayer(id)
	if err != nil {
		t.Fatalf("Failed to get player: %v", err)
	}

	return player
}

func setupDatabase(t *testing.T) *Database {
	cfg := models.SqliteConfig{DataSourceName: ":memory:", DriverName: "sqlite3"}
	db, err := Init(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	return db
}
