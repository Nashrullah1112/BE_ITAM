package domain

import (
	"time"
)

type Device struct {
	ID                    int
	RecipientLocation     *string
	ReceiptTime           *time.Time
	ReceiptProof          *string
	AssetType             *string
	AssetActivationTime   *time.Time
	AssetInspectionResult *string
	SerialNumber          *string
	Model                 *string
	WarrantyStartTime     *time.Time
	WarrantyCardNumber    *string
	Processor             *string
	RAMCapacity           *string
	ROMCapacity           *string
	RAMType               *string
	StorageType           *string
	AssetStatus           *string
	AssetValue            *string
	DepreciationValue     *string
	UsagePeriod           *int
	AssetOutTime          *time.Time
	AssetConditionOnExit  *string
	PurchaseReceipt       *string
	AssetID               *int
	DivisionID            *int
	UserID                *int
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
