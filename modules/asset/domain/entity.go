package domain

import (
	"time"
)

type Asset struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
