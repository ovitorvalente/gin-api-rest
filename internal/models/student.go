package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Major     string    `json:"major"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StudentFilter struct {
	ID   string `form:"id"`
	Name string `form:"name"`
}
