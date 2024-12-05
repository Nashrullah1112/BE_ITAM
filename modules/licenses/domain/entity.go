package domain

import (
	"time"
)

type License struct {
	ID                  int
	PurchaseDate        *time.Time
	InstalledDeviceSN   *string
	ActivationDate      *time.Time
	ExpirationDate      *time.Time
	AssetOwnershipType  *string
	LicenseCategory     *string
	LicenseVersion      *string
	MaxApplicationUsers *int
	MaxDeviceLicenses   *int
	LicenseType         *string
	AssetID             *int
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
