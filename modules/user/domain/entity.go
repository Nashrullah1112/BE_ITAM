package domain

import (
	"time"
)

type User struct {
	ID         int
	NIP        *string
	Name       *string
	Email      *string
	Password   *string
	Role       *string
	JoinDate   *time.Time
	DivisionID *int
	PositionID *int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
