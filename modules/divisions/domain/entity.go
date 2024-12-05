package domain

import (
	"time"
)

type Division struct {
	ID        int
	Name      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
