package models

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"title"`
	Description string    `json:"description"`
	Level       string    `json:"level"`
	Duration    int       `json:"duration"` // Duração em horas
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
