package domain

import (
	"time"
)

type License struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
