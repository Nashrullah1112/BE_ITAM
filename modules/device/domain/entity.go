package domain

import (
	"time"
)

type Device struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
