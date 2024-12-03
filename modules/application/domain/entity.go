package domain

import (
	"time"
)

type Application struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
