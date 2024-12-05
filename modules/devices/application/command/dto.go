package command

import (
	"time"
)

type CreateDeviceRequestDTO struct {
	RecipientLocation     string    `json:"recipient_location" validate:"required"`
	ReceiptTime           time.Time `json:"receipt_time" validate:"required"`
	ReceiptProof          string    `json:"receipt_proof" validate:"required"`
	AssetType             string    `json:"asset_type" validate:"required"`
	AssetActivationTime   time.Time `json:"asset_activation_time" validate:"required"`
	AssetInspectionResult string    `json:"asset_inspection_result" validate:"required"`
	SerialNumber          string    `json:"serial_number" validate:"required"`
	Model                 string    `json:"model" validate:"required"`
	WarrantyStartTime     time.Time `json:"warranty_start_time" validate:"required"`
	WarrantyCardNumber    string    `json:"warranty_card_number" validate:"required"`
	Processor             string    `json:"processor" validate:"required"`
	RAMCapacity           string    `json:"ram_capacity" validate:"required"`
	ROMCapacity           string    `json:"rom_capacity" validate:"required"`
	RAMType               string    `json:"ram_type" validate:"required"`
	StorageType           string    `json:"storage_type" validate:"required"`
	AssetStatus           string    `json:"asset_status" validate:"required"`
	AssetValue            string    `json:"asset_value" validate:"required"`
	DepreciationValue     string    `json:"depreciation_value" validate:"required"`
	UsagePeriod           int       `json:"usage_period" validate:"required,min=1"`
	AssetOutTime          time.Time `json:"asset_out_time" validate:"required"`
	AssetConditionOnExit  string    `json:"asset_condition_on_exit" validate:"required"`
	PurchaseReceipt       string    `json:"purchase_receipt" validate:"required"`
	AssetID               *int      `json:"asset_id" validate:"omitempty"`
	DivisionID            *int      `json:"division_id" validate:"omitempty"`
	UserID                *int      `json:"user_id" validate:"omitempty"`
}

type UpdateDeviceRequestDTO struct {
	ID                    int       `json:"id" validate:"required"`
	RecipientLocation     string    `json:"recipient_location" validate:"required"`
	ReceiptTime           time.Time `json:"receipt_time" validate:"required"`
	ReceiptProof          string    `json:"receipt_proof" validate:"required"`
	AssetType             string    `json:"asset_type" validate:"required"`
	AssetActivationTime   time.Time `json:"asset_activation_time" validate:"required"`
	AssetInspectionResult string    `json:"asset_inspection_result" validate:"required"`
	SerialNumber          string    `json:"serial_number" validate:"required"`
	Model                 string    `json:"model" validate:"required"`
	WarrantyStartTime     time.Time `json:"warranty_start_time" validate:"required"`
	WarrantyCardNumber    string    `json:"warranty_card_number" validate:"required"`
	Processor             string    `json:"processor" validate:"required"`
	RAMCapacity           string    `json:"ram_capacity" validate:"required"`
	ROMCapacity           string    `json:"rom_capacity" validate:"required"`
	RAMType               string    `json:"ram_type" validate:"required"`
	StorageType           string    `json:"storage_type" validate:"required"`
	AssetStatus           string    `json:"asset_status" validate:"required"`
	AssetValue            string    `json:"asset_value" validate:"required"`
	DepreciationValue     string    `json:"depreciation_value" validate:"required"`
	UsagePeriod           int       `json:"usage_period" validate:"required"`
	AssetOutTime          time.Time `json:"asset_out_time" validate:"required"`
	AssetConditionOnExit  string    `json:"asset_condition_on_exit" validate:"required"`
	PurchaseReceipt       string    `json:"purchase_receipt" validate:"required"`
	AssetID               *int      `json:"asset_id" validate:"omitempty"`
	DivisionID            *int      `json:"division_id" validate:"omitempty"`
	UserID                *int      `json:"user_id" validate:"omitempty"`
}

type UpdatePartialDeviceRequestDTO struct {
	ID                    int        `json:"id" validate:"required"`
	RecipientLocation     *string    `json:"recipient_location" validate:"omitempty"`
	ReceiptTime           *time.Time `json:"receipt_time" validate:"omitempty"`
	ReceiptProof          *string    `json:"receipt_proof" validate:"omitempty"`
	AssetType             *string    `json:"asset_type" validate:"omitempty"`
	AssetActivationTime   *time.Time `json:"asset_activation_time" validate:"omitempty"`
	AssetInspectionResult *string    `json:"asset_inspection_result" validate:"omitempty"`
	SerialNumber          *string    `json:"serial_number" validate:"omitempty"`
	Model                 *string    `json:"model" validate:"omitempty"`
	WarrantyStartTime     *time.Time `json:"warranty_start_time" validate:"omitempty"`
	WarrantyCardNumber    *string    `json:"warranty_card_number" validate:"omitempty"`
	Processor             *string    `json:"processor" validate:"omitempty"`
	RAMCapacity           *string    `json:"ram_capacity" validate:"omitempty"`
	ROMCapacity           *string    `json:"rom_capacity" validate:"omitempty"`
	RAMType               *string    `json:"ram_type" validate:"omitempty"`
	StorageType           *string    `json:"storage_type" validate:"omitempty"`
	AssetStatus           *string    `json:"asset_status" validate:"omitempty"`
	AssetValue            *string    `json:"asset_value" validate:"omitempty"`
	DepreciationValue     *string    `json:"depreciation_value" validate:"omitempty"`
	UsagePeriod           *int       `json:"usage_period" validate:"omitempty"`
	AssetOutTime          *time.Time `json:"asset_out_time" validate:"omitempty"`
	AssetConditionOnExit  *string    `json:"asset_condition_on_exit" validate:"omitempty"`
	PurchaseReceipt       *string    `json:"purchase_receipt" validate:"omitempty"`
	AssetID               *int       `json:"asset_id" validate:"omitempty"`
	DivisionID            *int       `json:"division_id" validate:"omitempty"`
	UserID                *int       `json:"user_id" validate:"omitempty"`
}

type DeleteDeviceRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type CreateDeviceResponseDTO struct {
	ID                    int        `json:"id"`
	RecipientLocation     *string    `json:"recipient_location"`
	ReceiptTime           *time.Time `json:"receipt_time"`
	ReceiptProof          *string    `json:"receipt_proof"`
	AssetType             *string    `json:"asset_type"`
	AssetActivationTime   *time.Time `json:"asset_activation_time"`
	AssetInspectionResult *string    `json:"asset_inspection_result"`
	SerialNumber          *string    `json:"serial_number"`
	Model                 *string    `json:"model"`
	WarrantyStartTime     *time.Time `json:"warranty_start_time"`
	WarrantyCardNumber    *string    `json:"warranty_card_number"`
	Processor             *string    `json:"processor"`
	RAMCapacity           *string    `json:"ram_capacity"`
	ROMCapacity           *string    `json:"rom_capacity"`
	RAMType               *string    `json:"ram_type"`
	StorageType           *string    `json:"storage_type"`
	AssetStatus           *string    `json:"asset_status"`
	AssetValue            *string    `json:"asset_value"`
	DepreciationValue     *string    `json:"depreciation_value"`
	UsagePeriod           *int       `json:"usage_period"`
	AssetOutTime          *time.Time `json:"asset_out_time"`
	AssetConditionOnExit  *string    `json:"asset_condition_on_exit"`
	PurchaseReceipt       *string    `json:"purchase_receipt"`
	AssetID               *int       `json:"asset_id"`
	DivisionID            *int       `json:"division_id"`
	UserID                *int       `json:"user_id"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

type UpdateDeviceResponseDTO struct {
	ID                    int        `json:"id"`
	RecipientLocation     *string    `json:"recipient_location"`
	ReceiptTime           *time.Time `json:"receipt_time"`
	ReceiptProof          *string    `json:"receipt_proof"`
	AssetType             *string    `json:"asset_type"`
	AssetActivationTime   *time.Time `json:"asset_activation_time"`
	AssetInspectionResult *string    `json:"asset_inspection_result"`
	SerialNumber          *string    `json:"serial_number"`
	Model                 *string    `json:"model"`
	WarrantyStartTime     *time.Time `json:"warranty_start_time"`
	WarrantyCardNumber    *string    `json:"warranty_card_number"`
	Processor             *string    `json:"processor"`
	RAMCapacity           *string    `json:"ram_capacity"`
	ROMCapacity           *string    `json:"rom_capacity"`
	RAMType               *string    `json:"ram_type"`
	StorageType           *string    `json:"storage_type"`
	AssetStatus           *string    `json:"asset_status"`
	AssetValue            *string    `json:"asset_value"`
	DepreciationValue     *string    `json:"depreciation_value"`
	UsagePeriod           *int       `json:"usage_period"`
	AssetOutTime          *time.Time `json:"asset_out_time"`
	AssetConditionOnExit  *string    `json:"asset_condition_on_exit"`
	PurchaseReceipt       *string    `json:"purchase_receipt"`
	AssetID               *int       `json:"asset_id"`
	DivisionID            *int       `json:"division_id"`
	UserID                *int       `json:"user_id"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

type UpdatePartialDeviceResponseDTO struct {
	ID                    int        `json:"id"`
	RecipientLocation     *string    `json:"recipient_location"`
	ReceiptTime           *time.Time `json:"receipt_time"`
	ReceiptProof          *string    `json:"receipt_proof"`
	AssetType             *string    `json:"asset_type"`
	AssetActivationTime   *time.Time `json:"asset_activation_time"`
	AssetInspectionResult *string    `json:"asset_inspection_result"`
	SerialNumber          *string    `json:"serial_number"`
	Model                 *string    `json:"model"`
	WarrantyStartTime     *time.Time `json:"warranty_start_time"`
	WarrantyCardNumber    *string    `json:"warranty_card_number"`
	Processor             *string    `json:"processor"`
	RAMCapacity           *string    `json:"ram_capacity"`
	ROMCapacity           *string    `json:"rom_capacity"`
	RAMType               *string    `json:"ram_type"`
	StorageType           *string    `json:"storage_type"`
	AssetStatus           *string    `json:"asset_status"`
	AssetValue            *string    `json:"asset_value"`
	DepreciationValue     *string    `json:"depreciation_value"`
	UsagePeriod           *int       `json:"usage_period"`
	AssetOutTime          *time.Time `json:"asset_out_time"`
	AssetConditionOnExit  *string    `json:"asset_condition_on_exit"`
	PurchaseReceipt       *string    `json:"purchase_receipt"`
	AssetID               *int       `json:"asset_id"`
	DivisionID            *int       `json:"division_id"`
	UserID                *int       `json:"user_id"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

type DeleteDeviceResponseDTO struct {
	ID int `json:"id"`
}
