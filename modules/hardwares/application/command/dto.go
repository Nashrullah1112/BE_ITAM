package command

import (
	"time"
)

type CreateHardwareRequestDTO struct {
	ReceiptDate          time.Time `json:"receipt_date" validate:"required"`
	ReceiptProof         string    `json:"receipt_proof" validate:"required"`
	AssetType            string    `json:"asset_type" validate:"required"`
	DeviceActivationDate time.Time `json:"device_activation_date" validate:"required"`
	InspectionResult     string    `json:"inspection_result" validate:"required"`
	SerialNumber         string    `json:"serial_number" validate:"required"`
	Model                string    `json:"model" validate:"required"`
	WarrantyStartDate    time.Time `json:"warranty_start_date" validate:"required"`
	WarrantyEndDate      time.Time `json:"warranty_end_date" validate:"required"`
	WarrantyCardNumber   string    `json:"warranty_card_number" validate:"required"`
	DeviceSpecifications string    `json:"device_specifications" validate:"required"`
	AssetStatus          string    `json:"asset_status" validate:"required"`
	AssetResponsible     string    `json:"asset_responsible" validate:"required"`
	StorageLocation      int       `json:"storage_location" validate:"required"`
	UsagePeriod          int       `json:"usage_period" validate:"required"`
	AssetOutTime         time.Time `json:"asset_out_time" validate:"required"`
	AssetCondition       string    `json:"asset_condition" validate:"required"`
	PurchaseReceipt      string    `json:"purchase_receipt" validate:"required"`
	AssetID              *int      `json:"asset_id" validate:"omitempty"`
	DivisionID           *int      `json:"division_id" validate:"omitempty"`
}

type UpdateHardwareRequestDTO struct {
	ID                   int       `json:"id" validate:"required"`
	ReceiptDate          time.Time `json:"receipt_date" validate:"required"`
	ReceiptProof         string    `json:"receipt_proof" validate:"required"`
	AssetType            string    `json:"asset_type" validate:"required"`
	DeviceActivationDate time.Time `json:"device_activation_date" validate:"required"`
	InspectionResult     string    `json:"inspection_result" validate:"required"`
	SerialNumber         string    `json:"serial_number" validate:"required"`
	Model                string    `json:"model" validate:"required"`
	WarrantyStartDate    time.Time `json:"warranty_start_date" validate:"required"`
	WarrantyEndDate      time.Time `json:"warranty_end_date" validate:"required"`
	WarrantyCardNumber   string    `json:"warranty_card_number" validate:"required"`
	DeviceSpecifications string    `json:"device_specifications" validate:"required"`
	AssetStatus          string    `json:"asset_status" validate:"required"`
	AssetResponsible     string    `json:"asset_responsible" validate:"required"`
	StorageLocation      int       `json:"storage_location" validate:"required"`
	UsagePeriod          int       `json:"usage_period" validate:"required"`
	AssetOutTime         time.Time `json:"asset_out_time" validate:"required"`
	AssetCondition       string    `json:"asset_condition" validate:"required"`
	PurchaseReceipt      string    `json:"purchase_receipt" validate:"required"`
	AssetID              *int      `json:"asset_id" validate:"omitempty"`
	DivisionID           *int      `json:"division_id" validate:"omitempty"`
}

type UpdatePartialHardwareRequestDTO struct {
	ID                   int        `json:"id" validate:"required"`
	ReceiptDate          *time.Time `json:"receipt_date" validate:"omitempty"`
	ReceiptProof         *string    `json:"receipt_proof" validate:"omitempty"`
	AssetType            *string    `json:"asset_type" validate:"omitempty"`
	DeviceActivationDate *time.Time `json:"device_activation_date" validate:"omitempty"`
	InspectionResult     *string    `json:"inspection_result" validate:"omitempty"`
	SerialNumber         *string    `json:"serial_number" validate:"omitempty"`
	Model                *string    `json:"model" validate:"omitempty"`
	WarrantyStartDate    *time.Time `json:"warranty_start_date" validate:"omitempty"`
	WarrantyEndDate      *time.Time `json:"warranty_end_date" validate:"omitempty"`
	WarrantyCardNumber   *string    `json:"warranty_card_number" validate:"omitempty"`
	DeviceSpecifications *string    `json:"device_specifications" validate:"omitempty"`
	AssetStatus          *string    `json:"asset_status" validate:"omitempty"`
	AssetResponsible     *string    `json:"asset_responsible" validate:"omitempty"`
	StorageLocation      *int       `json:"storage_location" validate:"omitempty"`
	UsagePeriod          *int       `json:"usage_period" validate:"omitempty"`
	AssetOutTime         *time.Time `json:"asset_out_time" validate:"omitempty"`
	AssetCondition       *string    `json:"asset_condition" validate:"omitempty"`
	PurchaseReceipt      *string    `json:"purchase_receipt" validate:"omitempty"`
	AssetID              *int       `json:"asset_id" validate:"omitempty"`
	DivisionID           *int       `json:"division_id" validate:"omitempty"`
}

type DeleteHardwareRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type CreateHardwareResponseDTO struct {
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

type UpdateHardwareResponseDTO struct {
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

type UpdatePartialHardwareResponseDTO struct {
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

type DeleteHardwareResponseDTO struct {
	ID int `json:"id"`
}
