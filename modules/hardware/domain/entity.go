package domain

import (
	"time"
)

type Hardware struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
