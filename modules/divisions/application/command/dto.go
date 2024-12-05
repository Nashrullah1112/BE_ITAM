package command

import (
	"time"
)

type CreateDivisionRequestDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateDivisionRequestDTO struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdatePartialDivisionRequestDTO struct {
	ID   int     `json:"id" validate:"required"`
	Name *string `json:"name" validate:"omitempty"`
}

type DeleteDivisionRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type CreateDivisionResponseDTO struct {
	ID        int       `json:"id"`
	Name      *string   `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateDivisionResponseDTO struct {
	ID        int       `json:"id"`
	Name      *string   `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePartialDivisionResponseDTO struct {
	ID        int       `json:"id"`
	Name      *string   `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteDivisionResponseDTO struct {
	ID int `json:"id"`
}
