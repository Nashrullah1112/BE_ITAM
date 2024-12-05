package domain

import (
	"time"
)

type Hardware struct {
	ID                   int
	ReceiptDate          *time.Time
	ReceiptProof         *string
	AssetType            *string
	DeviceActivationDate *time.Time
	InspectionResult     *string
	SerialNumber         *string
	Model                *string
	WarrantyStartDate    *time.Time
	WarrantyEndDate      *time.Time
	WarrantyCardNumber   *string
	DeviceSpecifications *string
	AssetStatus          *string
	AssetResponsible     *string
	StorageLocation      *int
	UsagePeriod          *int
	AssetOutTime         *time.Time
	AssetCondition       *string
	PurchaseReceipt      *string
	AssetID              *int
	DivisionID           *int
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
