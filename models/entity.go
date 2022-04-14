package models

import (
	"time"
)

// Entity contains the basic fields for database entities
type Entity struct {
	ID        int          `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
