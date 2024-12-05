package query

import (
	"time"
)

type LicenseResponseDTO struct {
	ID                  int        `json:"id"`
	PurchaseDate        *time.Time `json:"purchase_date"`
	InstalledDeviceSN   *string    `json:"installed_device_sn"`
	ActivationDate      *time.Time `json:"activation_date"`
	ExpirationDate      *time.Time `json:"expiration_date"`
	AssetOwnershipType  *string    `json:"asset_ownership_type"`
	LicenseCategory     *string    `json:"license_category"`
	LicenseVersion      *string    `json:"license_version"`
	MaxApplicationUsers *int       `json:"max_application_users"`
	MaxDeviceLicenses   *int       `json:"max_device_licenses"`
	LicenseType         *string    `json:"license_type"`
	AssetID             *int       `json:"asset_id"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
