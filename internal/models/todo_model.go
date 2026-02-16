package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	UserID      uuid.UUID `json:"users_id"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type TodoResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	UserID      string    `json:"users_id"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
