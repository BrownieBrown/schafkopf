package models

import "github.com/google/uuid"

type Player struct {
	ID    uuid.UUID
	Score float64
}

type PlayerRepository interface {
	CreatePlayer(player *Player) (string, error)
	GetPlayer(id string) (*Player, error)
	UpdatePlayer(player *Player) error
	DeletePlayer(id string) error
}
