package domain

import (
	"time"
)

type Application struct {
	ID                       int
	ApplicationName          *string
	CreationDate             *time.Time
	AcceptanceDate           *time.Time
	StorageServerLocation    *string
	ApplicationType          *string
	ApplicationLink          *string
	ApplicationCertification *string
	ActivationDate           *time.Time
	ExpirationDate           *time.Time
	AssetID                  *int
	CreatedAt                time.Time
	UpdatedAt                time.Time
}
