package models

import (
	"github.com/google/uuid"
	"time"
)

type Game struct {
	ID        uuid.UUID
	CreatedAt time.Time
}
