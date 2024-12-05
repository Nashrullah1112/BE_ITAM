package query

import (
	"time"
)

type DeviceResponseDTO struct {
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
