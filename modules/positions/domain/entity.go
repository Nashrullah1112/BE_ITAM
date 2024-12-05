package domain

import (
	"time"
)

type Position struct {
	ID        int
	Name      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
