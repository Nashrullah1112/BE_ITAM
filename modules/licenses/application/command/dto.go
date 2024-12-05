package command

import (
	"time"
)

type CreateLicenseRequestDTO struct {
	PurchaseDate        time.Time `json:"purchase_date" validate:"required"`
	InstalledDeviceSN   string    `json:"installed_device_sn" validate:"required"`
	ActivationDate      time.Time `json:"activation_date" validate:"required"`
	ExpirationDate      time.Time `json:"expiration_date" validate:"required"`
	AssetOwnershipType  string    `json:"asset_ownership_type" validate:"required"`
	LicenseCategory     string    `json:"license_category" validate:"required"`
	LicenseVersion      string    `json:"license_version" validate:"required"`
	MaxApplicationUsers int       `json:"max_application_users" validate:"required"`
	MaxDeviceLicenses   int       `json:"max_device_licenses" validate:"required"`
	LicenseType         string    `json:"license_type" validate:"required"`
	AssetID             *int      `json:"asset_id" validate:"omitempty"`
}

type UpdateLicenseRequestDTO struct {
	ID                  int       `json:"id" validate:"required"`
	PurchaseDate        time.Time `json:"purchase_date" validate:"required"`
	InstalledDeviceSN   string    `json:"installed_device_sn" validate:"required"`
	ActivationDate      time.Time `json:"activation_date" validate:"required"`
	ExpirationDate      time.Time `json:"expiration_date" validate:"required"`
	AssetOwnershipType  string    `json:"asset_ownership_type" validate:"required"`
	LicenseCategory     string    `json:"license_category" validate:"required"`
	LicenseVersion      string    `json:"license_version" validate:"required"`
	MaxApplicationUsers int       `json:"max_application_users" validate:"required"`
	MaxDeviceLicenses   int       `json:"max_device_licenses" validate:"required"`
	LicenseType         string    `json:"license_type" validate:"required"`
	AssetID             *int      `json:"asset_id" validate:"omitempty"`
}

type UpdatePartialLicenseRequestDTO struct {
	ID                  int        `json:"id" validate:"required"`
	PurchaseDate        *time.Time `json:"purchase_date" validate:"omitempty"`
	InstalledDeviceSN   *string    `json:"installed_device_sn" validate:"omitempty"`
	ActivationDate      *time.Time `json:"activation_date" validate:"omitempty"`
	ExpirationDate      *time.Time `json:"expiration_date" validate:"omitempty"`
	AssetOwnershipType  *string    `json:"asset_ownership_type" validate:"omitempty"`
	LicenseCategory     *string    `json:"license_category" validate:"omitempty"`
	LicenseVersion      *string    `json:"license_version" validate:"omitempty"`
	MaxApplicationUsers *int       `json:"max_application_users" validate:"omitempty"`
	MaxDeviceLicenses   *int       `json:"max_device_licenses" validate:"omitempty"`
	LicenseType         *string    `json:"license_type" validate:"omitempty"`
	AssetID             *int       `json:"asset_id" validate:"omitempty"`
}

type DeleteLicenseRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type CreateLicenseResponseDTO struct {
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

type UpdateLicenseResponseDTO struct {
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

type UpdatePartialLicenseResponseDTO struct {
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

type DeleteLicenseResponseDTO struct {
	ID int `json:"id"`
}
