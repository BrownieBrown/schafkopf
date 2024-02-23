package sqlite

import (
	"database/sql"
	"github.com/BrownieBrown/schafkopf/internal/models"
	"github.com/google/uuid"
	"log"
)

type PlayerRepository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) models.PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (r *PlayerRepository) CreatePlayer(player *models.Player) (string, error) {
	player.ID = uuid.New()
	player.Score = 0
	query := `INSERT INTO player (id, score) VALUES (?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "", err

	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	_, err = stmt.Exec(player.ID.String(), player.Score)
	if err != nil {
		return "", err
	}

	return player.ID.String(), nil
}

func (r *PlayerRepository) GetPlayer(id string) (*models.Player, error) {
	query := `SELECT id, score FROM player WHERE id = ?`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	var player models.Player
	err = stmt.QueryRow(id).Scan(&player.ID, &player.Score)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (r *PlayerRepository) UpdatePlayer(player *models.Player) error {
	query := `UPDATE player SET score = ? WHERE id = ?`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	_, err = stmt.Exec(player.Score, player.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PlayerRepository) DeletePlayer(id string) error {
	query := `DELETE FROM player WHERE id = ?`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
