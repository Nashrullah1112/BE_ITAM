package command

import (
	"time"
)

type CreateApplicationRequestDTO struct {
	ApplicationName          string    `json:"application_name" validate:"required"`
	CreationDate             time.Time `json:"creation_date" validate:"required"`
	AcceptanceDate           time.Time `json:"acceptance_date" validate:"required"`
	StorageServerLocation    string    `json:"storage_server_location" validate:"required"`
	ApplicationType          string    `json:"application_type" validate:"required"`
	ApplicationLink          string    `json:"application_link" validate:"required"`
	ApplicationCertification string    `json:"application_certification" validate:"required"`
	ActivationDate           time.Time `json:"activation_date" validate:"required"`
	ExpirationDate           time.Time `json:"expiration_date" validate:"required"`
	AssetID                  *int      `json:"asset_id" validate:"omitempty"`
}

type UpdateApplicationRequestDTO struct {
	ID                       int       `json:"id" validate:"required"`
	ApplicationName          string    `json:"application_name" validate:"required"`
	CreationDate             time.Time `json:"creation_date" validate:"required"`
	AcceptanceDate           time.Time `json:"acceptance_date" validate:"required"`
	StorageServerLocation    string    `json:"storage_server_location" validate:"required"`
	ApplicationType          string    `json:"application_type" validate:"required"`
	ApplicationLink          string    `json:"application_link" validate:"required"`
	ApplicationCertification string    `json:"application_certification" validate:"required"`
	ActivationDate           time.Time `json:"activation_date" validate:"required"`
	ExpirationDate           time.Time `json:"expiration_date" validate:"required"`
	AssetID                  *int      `json:"asset_id" validate:"omitempty"`
}

type UpdatePartialApplicationRequestDTO struct {
	ID                       int        `json:"id" validate:"required"`
	ApplicationName          *string    `json:"application_name" validate:"omitempty"`
	CreationDate             *time.Time `json:"creation_date" validate:"omitempty"`
	AcceptanceDate           *time.Time `json:"acceptance_date" validate:"omitempty"`
	StorageServerLocation    *string    `json:"storage_server_location" validate:"omitempty"`
	ApplicationType          *string    `json:"application_type" validate:"omitempty"`
	ApplicationLink          *string    `json:"application_link" validate:"omitempty"`
	ApplicationCertification *string    `json:"application_certification" validate:"omitempty"`
	ActivationDate           *time.Time `json:"activation_date" validate:"omitempty"`
	ExpirationDate           *time.Time `json:"expiration_date" validate:"omitempty"`
	AssetID                  *int       `json:"asset_id" validate:"omitempty"`
}

type DeleteApplicationRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type CreateApplicationResponseDTO struct {
	ID                       int        `json:"id"`
	ApplicationName          *string    `json:"application_name"`
	CreationDate             *time.Time `json:"creation_date"`
	AcceptanceDate           *time.Time `json:"acceptance_date"`
	StorageServerLocation    *string    `json:"storage_server_location"`
	ApplicationType          *string    `json:"application_type"`
	ApplicationLink          *string    `json:"application_link"`
	ApplicationCertification *string    `json:"application_certification"`
	ActivationDate           *time.Time `json:"activation_date"`
	ExpirationDate           *time.Time `json:"expiration_date"`
	AssetID                  *int       `json:"asset_id"`
	CreatedAt                time.Time  `json:"created_at"`
	UpdatedAt                time.Time  `json:"updated_at"`
}

type UpdateApplicationResponseDTO struct {
	ID                       int        `json:"id"`
	ApplicationName          *string    `json:"application_name"`
	CreationDate             *time.Time `json:"creation_date"`
	AcceptanceDate           *time.Time `json:"acceptance_date"`
	StorageServerLocation    *string    `json:"storage_server_location"`
	ApplicationType          *string    `json:"application_type"`
	ApplicationLink          *string    `json:"application_link"`
	ApplicationCertification *string    `json:"application_certification"`
	ActivationDate           *time.Time `json:"activation_date"`
	ExpirationDate           *time.Time `json:"expiration_date"`
	AssetID                  *int       `json:"asset_id"`
	CreatedAt                time.Time  `json:"created_at"`
	UpdatedAt                time.Time  `json:"updated_at"`
}

type UpdatePartialApplicationResponseDTO struct {
	ID                       int        `json:"id"`
	ApplicationName          *string    `json:"application_name"`
	CreationDate             *time.Time `json:"creation_date"`
	AcceptanceDate           *time.Time `json:"acceptance_date"`
	StorageServerLocation    *string    `json:"storage_server_location"`
	ApplicationType          *string    `json:"application_type"`
	ApplicationLink          *string    `json:"application_link"`
	ApplicationCertification *string    `json:"application_certification"`
	ActivationDate           *time.Time `json:"activation_date"`
	ExpirationDate           *time.Time `json:"expiration_date"`
	AssetID                  *int       `json:"asset_id"`
	CreatedAt                time.Time  `json:"created_at"`
	UpdatedAt                time.Time  `json:"updated_at"`
}

type DeleteApplicationResponseDTO struct {
	ID int `json:"id"`
}
