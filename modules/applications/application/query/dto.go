package query

import (
	"time"
)

type ApplicationResponseDTO struct {
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
