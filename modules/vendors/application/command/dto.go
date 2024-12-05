package command

import (
	"time"
)

type CreateVendorRequestDTO struct {
	ContactPerson string `json:"contact_person" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	ContactNumber string `json:"contact_number" validate:"required"`
	Location      string `json:"location" validate:"required"`
	SIUPNumber    string `json:"siup_number" validate:"required"`
	NIBNumber     string `json:"nib_number" validate:"required"`
	NPWPNumber    string `json:"npwp_number" validate:"required"`
}

type UpdateVendorRequestDTO struct {
	ID            int    `json:"id" validate:"required"`
	ContactPerson string `json:"contact_person" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	ContactNumber string `json:"contact_number" validate:"required"`
	Location      string `json:"location" validate:"required"`
	SIUPNumber    string `json:"siup_number" validate:"required"`
	NIBNumber     string `json:"nib_number" validate:"required"`
	NPWPNumber    string `json:"npwp_number" validate:"required"`
}

type UpdatePartialVendorRequestDTO struct {
	ID            int     `json:"id" validate:"required"`
	ContactPerson *string `json:"contact_person" validate:"omitempty"`
	Email         *string `json:"email" validate:"omitempty,email"`
	ContactNumber *string `json:"contact_number" validate:"omitempty"`
	Location      *string `json:"location" validate:"omitempty"`
	SIUPNumber    *string `json:"siup_number" validate:"omitempty"`
	NIBNumber     *string `json:"nib_number" validate:"omitempty"`
	NPWPNumber    *string `json:"npwp_number" validate:"omitempty"`
}

type DeleteVendorRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type CreateVendorResponseDTO struct {
	ID            int       `json:"id"`
	ContactPerson *string   `json:"contact_person"`
	Email         *string   `json:"email"`
	ContactNumber *string   `json:"contact_number"`
	Location      *string   `json:"location"`
	SIUPNumber    *string   `json:"siup_number"`
	NIBNumber     *string   `json:"nib_number"`
	NPWPNumber    *string   `json:"npwp_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UpdateVendorResponseDTO struct {
	ID            int       `json:"id"`
	ContactPerson *string   `json:"contact_person"`
	Email         *string   `json:"email"`
	ContactNumber *string   `json:"contact_number"`
	Location      *string   `json:"location"`
	SIUPNumber    *string   `json:"siup_number"`
	NIBNumber     *string   `json:"nib_number"`
	NPWPNumber    *string   `json:"npwp_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UpdatePartialVendorResponseDTO struct {
	ID            int       `json:"id"`
	ContactPerson *string   `json:"contact_person"`
	Email         *string   `json:"email"`
	ContactNumber *string   `json:"contact_number"`
	Location      *string   `json:"location"`
	SIUPNumber    *string   `json:"siup_number"`
	NIBNumber     *string   `json:"nib_number"`
	NPWPNumber    *string   `json:"npwp_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type DeleteVendorResponseDTO struct {
	ID int `json:"id"`
}
