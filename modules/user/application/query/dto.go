package query

import (
	"time"
)

type UserResponseDTO struct {
	ID         int        `json:"id"`
	NIP        string     `json:"nip"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Role       string     `json:"role"`
	JoinDate   *time.Time `json:"join_date"`
	DivisionID *int       `json:"division_id"`
	PositionID *int       `json:"position_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
