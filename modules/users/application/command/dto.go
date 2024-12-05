package command

import (
	"time"
)

type CreateUserRequestDTO struct {
	NIP        string    `json:"nip" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	Role       string    `json:"role" validate:"required"`
	JoinDate   time.Time `json:"join_date" validate:"required"`
	DivisionID *int      `json:"division_id" validate:"omitempty"`
	PositionID *int      `json:"position_id" validate:"omitempty"`
}

type UpdateUserRequestDTO struct {
	ID         int       `json:"id" validate:"required"`
	NIP        string    `json:"nip" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	Role       string    `json:"role" validate:"required"`
	JoinDate   time.Time `json:"join_date" validate:"required"`
	DivisionID *int      `json:"division_id" validate:"omitempty"`
	PositionID *int      `json:"position_id" validate:"omitempty"`
}

type UpdatePartialUserRequestDTO struct {
	ID         int        `json:"id" validate:"required"`
	NIP        *string    `json:"nip" validate:"omitempty"`
	Name       *string    `json:"name" validate:"omitempty"`
	Email      *string    `json:"email" validate:"omitempty,email"`
	Password   *string    `json:"password" validate:"omitempty"`
	Role       *string    `json:"role" validate:"omitempty"`
	JoinDate   *time.Time `json:"join_date" validate:"omitempty"`
	DivisionID *int       `json:"division_id" validate:"omitempty"`
	PositionID *int       `json:"position_id" validate:"omitempty"`
}

type DeleteUserRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type CreateUserResponseDTO struct {
	ID         int        `json:"id"`
	NIP        *string    `json:"nip"`
	Name       *string    `json:"name"`
	Email      *string    `json:"email"`
	Role       *string    `json:"role"`
	JoinDate   *time.Time `json:"join_date"`
	DivisionID *int       `json:"division_id"`
	PositionID *int       `json:"position_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type UpdateUserResponseDTO struct {
	ID         int        `json:"id"`
	NIP        *string    `json:"nip"`
	Name       *string    `json:"name"`
	Email      *string    `json:"email"`
	Role       *string    `json:"role"`
	JoinDate   *time.Time `json:"join_date"`
	DivisionID *int       `json:"division_id"`
	PositionID *int       `json:"position_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type UpdatePartialUserResponseDTO struct {
	ID         int        `json:"id"`
	NIP        *string    `json:"nip"`
	Name       *string    `json:"name"`
	Email      *string    `json:"email"`
	Role       *string    `json:"role"`
	JoinDate   *time.Time `json:"join_date"`
	DivisionID *int       `json:"division_id"`
	PositionID *int       `json:"position_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type DeleteUserResponseDTO struct {
	ID int `json:"id"`
}
