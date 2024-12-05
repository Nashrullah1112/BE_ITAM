package command

import (
	"time"
)

type CreatePositionRequestDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePositionRequestDTO struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdatePartialPositionRequestDTO struct {
	ID   int     `json:"id" validate:"required"`
	Name *string `json:"name" validate:"omitempty"`
}

type DeletePositionRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type CreatePositionResponseDTO struct {
	ID        int       `json:"id"`
	Name      *string   `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePositionResponseDTO struct {
	ID        int       `json:"id"`
	Name      *string   `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePartialPositionResponseDTO struct {
	ID        int       `json:"id"`
	Name      *string   `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletePositionResponseDTO struct {
	ID int `json:"id"`
}
