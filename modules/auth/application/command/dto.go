package command

import (
	"time"
)

type AuthRegisterRequestDTO struct {
	NIP             string `json:"nip" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type AuthLoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

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

type AuthRegisterResponseDTO struct {
	Token  string          `json:"token"`
	Expire time.Time       `json:"expire"`
	User   UserResponseDTO `json:"user"`
}

type AuthLoginResponseDTO struct {
	Token  string          `json:"token"`
	Expire time.Time       `json:"expire"`
	User   UserResponseDTO `json:"user"`
}
