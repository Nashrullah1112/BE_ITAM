package query

import (
	"time"
)

type HardwareResponseDTO struct {
	ID                   int        `json:"id"`
	ReceiptDate          *time.Time `json:"receipt_date"`
	ReceiptProof         *string    `json:"receipt_proof"`
	AssetType            *string    `json:"asset_type"`
	DeviceActivationDate *time.Time `json:"device_activation_date"`
	InspectionResult     *string    `json:"inspection_result"`
	SerialNumber         *string    `json:"serial_number"`
	Model                *string    `json:"model"`
	WarrantyStartDate    *time.Time `json:"warranty_start_date"`
	WarrantyEndDate      *time.Time `json:"warranty_end_date"`
	WarrantyCardNumber   *string    `json:"warranty_card_number"`
	DeviceSpecifications *string    `json:"device_specifications"`
	AssetStatus          *string    `json:"asset_status"`
	AssetResponsible     *string    `json:"asset_responsible"`
	StorageLocation      *int       `json:"storage_location"`
	UsagePeriod          *int       `json:"usage_period"`
	AssetOutTime         *time.Time `json:"asset_out_time"`
	AssetCondition       *string    `json:"asset_condition"`
	PurchaseReceipt      *string    `json:"purchase_receipt"`
	AssetID              *int       `json:"asset_id"`
	DivisionID           *int       `json:"division_id"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}
