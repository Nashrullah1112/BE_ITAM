package command

import (
	"time"
)

type UserCreateRequestDTO struct {
	NIP        string    `json:"nip" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	Role       string    `json:"role" validate:"required"`
	JoinDate   time.Time `json:"join_date" validate:"required"`
	DivisionID *int      `json:"division_id" validate:"omitempty"`
	PositionID *int      `json:"position_id" validate:"omitempty"`
}

type UserUpdateRequestDTO struct {
	ID         int        `json:"id" validate:"required"`
	NIP        string     `json:"nip" validate:"required"`
	Name       string     `json:"name" validate:"required"`
	Email      string     `json:"email" validate:"required,email"`
	Password   string     `json:"password" validate:"required"`
	Role       string     `json:"role" validate:"required"`
	JoinDate   *time.Time `json:"join_date" validate:"omitempty"`
	DivisionID *int       `json:"division_id" validate:"omitempty"`
	PositionID *int       `json:"position_id" validate:"omitempty"`
}

type UserUpdatePartialRequestDTO struct {
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

type UserDeleteRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type UserCreateResponseDTO struct {
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

type UserUpdateResponseDTO struct {
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

type UserUpdatePartialResponseDTO struct {
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

type UserDeleteResponseDTO struct {
	ID int `json:"id"`
}
